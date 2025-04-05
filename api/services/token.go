package services

import (
	"context"
	"fmt"
	"time"

	"github.com/g-villarinho/flash-buy-api/config"
	"github.com/g-villarinho/flash-buy-api/pkgs"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	CreateToken(ctx context.Context, userID, sessionID, email string, iat, exp time.Time) (string, error)
}

type tokenService struct {
	di *pkgs.Di
	kp pkgs.EcdsaKeyPair
}

func NewTokenService(di *pkgs.Di) (TokenService, error) {
	ecdsaKeyPair, err := pkgs.Invoke[pkgs.EcdsaKeyPair](di)
	if err != nil {
		return nil, fmt.Errorf("invoke ecdsa key pair: %w", err)
	}

	return &tokenService{
		di: di,
		kp: ecdsaKeyPair,
	}, nil
}

func (t *tokenService) CreateToken(ctx context.Context, userID string, sessionID string, email string, iat time.Time, exp time.Time) (string, error) {
	privateKey, err := t.kp.ParseECDSAPrivateKey(config.Env.Key.PrivateKey)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"iss":   "xp-life-app",
		"email": email,
		"iat":   iat.Unix(),
		"exp":   exp.Unix(),
	}

	if userID != "" {
		claims["sub"] = userID
	}

	if sessionID != "" {
		claims["sid"] = sessionID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signedToken, nil
}
