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
		user := &database.User{}
		req := &struct {
			Action string `form:"action"`
			Email  string `form:"email" validate:"email"`
			Pwd    string `form:"pwd"`
		}{}
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		if req.Action == "register" {
			hash, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user := &database.User{
				Email:           req.Email,
				Password:        hash,
				Level:           database.LevelPending,
				ProfileImageURL: "",
				CreatedAt:       int(time.Now().Unix()),
				UpdatedAt:       int(time.Now().Unix()),
				Settings:        "",
			}
			_, err = database.DB.NamedExec(`
    INSERT INTO user (  email, password, level, profile_image_url, created_at, updated_at, settings)
    VALUES (  :email,:password, :level,:profile_image_url, :created_at, :updated_at, :settings)
`, user)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, "success")
		}
		if req.Action == "login" {
			if err := database.DB.Get(user, `SELECT * FROM user WHERE email = ?`, req.Email); err != nil {
				return err
			}
			if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Pwd)); err == nil {
				token, err := jwt.NewWithClaims(jwt.SigningMethodHS256,
					&jwt.RegisteredClaims{
						Subject:   user.Level,
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), //过期日期设置7天

					},
				).SignedString([]byte(strconv.Itoa(os.Getpid())))
				if err != nil {
					return err
				}
				c.SetCookie(&http.Cookie{
					Name:     "moechat_token",
					Value:    token,
					HttpOnly: true,
				})
				return c.Redirect(http.StatusFound, "/c/")
			} else {
				return echo.ErrUnauthorized
			}
		}
	}
	return echo.ErrMethodNotAllowed
}
