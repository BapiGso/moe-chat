package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"moechat/core/database"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Login(c echo.Context) error {
	qpu := &database.QPU{}
	req := &struct {
		Action string `query:"action"`
		Mail   string `form:"mail" validate:"email"`
		Name   string `form:"Name"`
		Pwd    string `form:"pwd"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodGet:
		return c.Render(http.StatusOK, "login.html", nil)
	case http.MethodPost:
		if err := database.DB.Get(&qpu.User, `SELECT * FROM  user WHERE name = ?`, req.Mail); err != nil {
			return err
		}
		if err := bcrypt.CompareHashAndPassword([]byte(qpu.User.ID), []byte(req.Pwd)); err == nil {
			token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), //过期日期设置7天
			}).SignedString([]byte(strconv.Itoa(os.Getpid())))
			if err != nil {
				return err
			}
			c.SetCookie(&http.Cookie{
				Name:     "panel_token",
				Value:    token,
				HttpOnly: true,
			})
			return c.Redirect(302, "/c/")
		}
		return echo.ErrUnauthorized
	}
	return echo.ErrMethodNotAllowed
}
