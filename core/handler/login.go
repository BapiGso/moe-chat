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
	switch c.Request().Method {
	case http.MethodGet:
		return c.Render(http.StatusOK, "login.html", nil)
	case http.MethodPost:
		req := &struct {
			Email string `form:"email" validate:"email"`
			Pwd   string `form:"pwd" validate:"required,min=1,max=200"`
		}{}
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		user := new(database.User)
		if err := database.DB.Get(user, `SELECT * FROM user WHERE email = ?`, req.Email); err != nil {
			return err
		}
		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Pwd)); err == nil {
			if user.Level == database.LevelPending {
				return c.JSON(http.StatusUnauthorized, "等待管理员审核")
			}
			token, err := jwt.NewWithClaims(jwt.SigningMethodHS256,
				&jwt.RegisteredClaims{
					Subject:   user.Email,
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 14)), //过期日期设置14天

				},
			).SignedString([]byte(strconv.Itoa(os.Getpid())))
			if err != nil {
				return err
			}
			c.SetCookie(&http.Cookie{
				Name:     "moe-chat_token",
				Value:    token,
				HttpOnly: true,
			})
			return c.NoContent(http.StatusNoContent)
		} else {
			return c.JSON(http.StatusUnauthorized, "用户名或密码不正确")
		}
	}
	return echo.ErrMethodNotAllowed
}
