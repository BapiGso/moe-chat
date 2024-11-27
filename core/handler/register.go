package handler

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"moechat/core/database"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	req := &struct {
		Email string `form:"email" validate:"email"`
		Pwd   string `form:"pwd"`
	}{}
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	var config database.Config
	err := database.DB.Get(&config, `SELECT * FROM config WHERE key = 'enableRegister'`)
	if err != nil || config.Val != "1" {
		return c.JSON(http.StatusUnauthorized, "管理员未开启注册")
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
	//如果没有任何账户则这个账户为管理员账户
	var count int
	err = database.DB.Get(&count, `SELECT COUNT(*) FROM user`)
	if err != nil {
		return err
	}

	if count == 0 {
		user.Level = database.LevelAdmin
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
