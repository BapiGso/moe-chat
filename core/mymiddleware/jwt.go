package mymiddleware

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var debug bool

var JWT, _ = echojwt.Config{
	SuccessHandler: func(c echo.Context) {
		user, _ := c.Get("user").(*jwt.Token)
		issuer, _ := user.Claims.GetSubject()
		c.Set("email", issuer)
	},
	ErrorHandler: func(c echo.Context, err error) error {
		return c.Redirect(http.StatusFound, "/login")
	},
	SigningKey:  []byte(strconv.Itoa(os.Getpid())),
	TokenLookup: "cookie:moechat_token ,query:moechat_token",
	Skipper: func(c echo.Context) bool {
		assetsPath := strings.HasPrefix(c.Path(), "/assets/")
		loginPath := c.Path() == "/login"
		if debug {
			c.Set("email", "debug")
		}
		return assetsPath || loginPath || debug
	},
}.ToMiddleware()
