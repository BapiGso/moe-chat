package handler

import (
	"database/sql"
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"moechat/core/database"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	switch c.Request().Method {
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

		//如果没有任何admin账户则这个账户为admin账户
		err = database.DB.Get(user, `SELECT null FROM user WHERE level='admin'`)
		if errors.Is(err, sql.ErrNoRows) {
			user.Level = database.LevelAdmin
		}

		config := new(database.Config)
		err = database.DB.Get(config, `SELECT * FROM config WHERE key = 'enableRegister'`)
		if !errors.Is(err, sql.ErrNoRows) || config.Val == "0" {
			return c.JSON(http.StatusUnauthorized, "管理员未开启注册")
		}

		_, err = database.DB.NamedExec(`
    INSERT INTO user ( email, password, level, profile_image_url, created_at, updated_at, settings)
    VALUES (  :email,:password, :level,:profile_image_url, :created_at, :updated_at, :settings)
`, user)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "邮箱已被注册")
		}
		return c.JSON(http.StatusOK, "注册成功")
	}
	return echo.ErrMethodNotAllowed
}
