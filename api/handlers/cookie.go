package handlers

import (
	"net/http"
	"time"

	"github.com/g-villarinho/xp-life-api/config"
	"github.com/g-villarinho/xp-life-api/models"
	"github.com/labstack/echo/v4"
)

func SetCookieSession(ectx echo.Context, session models.Session) {
	cookie := new(http.Cookie)
	cookie.Name = config.Env.Cookie.Name
	cookie.Path = "/"
	cookie.Value = session.Token
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Expires = session.ExpiresAt
	cookie.MaxAge = int(time.Until(session.ExpiresAt).Seconds())

	ectx.SetCookie(cookie)
}

func DelCookieSession(ectx echo.Context) {
	ectx.SetCookie(&http.Cookie{
		Name:     config.Env.Cookie.Name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}
