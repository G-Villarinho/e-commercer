package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/g-villarinho/xp-life-api/config"
	"github.com/g-villarinho/xp-life-api/handlers"
	"github.com/g-villarinho/xp-life-api/models"
	"github.com/g-villarinho/xp-life-api/pkgs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware interface {
	Authenticate(next echo.HandlerFunc) echo.HandlerFunc
	AuthenticateWithoutEmailVerification(next echo.HandlerFunc) echo.HandlerFunc
	GetClaims(tokenString string) (*models.TokenClaims, error)
}

type authMiddleware struct {
	di  *pkgs.Di
	ep  pkgs.EcdsaKeyPair
	rdp pkgs.RequestDataCtx
}

func NewAuthMiddleware(di *pkgs.Di) (AuthMiddleware, error) {
	ecdsaKeyPair, err := pkgs.Invoke[pkgs.EcdsaKeyPair](di)
	if err != nil {
		return nil, fmt.Errorf("invoke ecdsa key pair: %w", err)
	}

	ctxData, err := pkgs.Invoke[pkgs.RequestDataCtx](di)
	if err != nil {
		return nil, err
	}

	return authMiddleware{
		di:  di,
		ep:  ecdsaKeyPair,
		rdp: ctxData,
	}, nil
}

func (a authMiddleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		cookie, err := ectx.Cookie(config.Env.Cookie.Name)
		if err != nil {
			return ectx.NoContent(http.StatusUnauthorized)
		}

		claims, err := a.GetClaims(cookie.Value)
		if err != nil {
			handlers.DelCookieSession(ectx)
			return ectx.NoContent(http.StatusUnauthorized)
		}

		ctx := a.rdp.SetToken(ectx.Request().Context(), cookie.Value)
		ctx = a.rdp.SetUserID(ctx, claims.Sub)
		ctx = a.rdp.SetEmail(ctx, claims.Email)
		ectx.SetRequest(ectx.Request().WithContext(ctx))

		return next(ectx)
	}
}

func (a authMiddleware) AuthenticateWithoutEmailVerification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		cookie, err := ectx.Cookie(config.Env.Cookie.Name)
		if err != nil {
			return ectx.NoContent(http.StatusUnauthorized)
		}

		claims, err := a.GetClaims(cookie.Value)
		if err != nil {
			handlers.DelCookieSession(ectx)
			return ectx.NoContent(http.StatusUnauthorized)
		}

		ctx := a.rdp.SetToken(ectx.Request().Context(), cookie.Value)
		ctx = a.rdp.SetEmail(ctx, claims.Email)
		ectx.SetRequest(ectx.Request().WithContext(ctx))

		return next(ectx)
	}
}

func (a authMiddleware) GetClaims(tokenString string) (*models.TokenClaims, error) {
	publicKey, err := a.ep.ParseECDSAPublicKey(config.Env.Key.PublicKey)
	if err != nil || publicKey == nil {
		return nil, fmt.Errorf("failed to convert public key: %w", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.TokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
