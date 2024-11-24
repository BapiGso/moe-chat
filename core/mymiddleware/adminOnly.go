package mymiddleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"moechat/core/database"
	"net/http"
	"os"
)

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 从 JWT 中获取用户信息
		if os.Getenv("MOECHAT_DEBUG") == "1" {
			return next(c)
		}
		userToken, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		email, err := userToken.Claims.GetSubject()
		if err != nil {
			return err
		}

		// 从数据库中获取用户角色
		var level string
		err = database.DB.Get(&level, "SELECT level FROM user WHERE email = ?", email)
		if err != nil {
			return err
		}

		// 检查角色是否为 "admin"
		if level != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "访问被拒绝")
		}

		return next(c)
	}
}
