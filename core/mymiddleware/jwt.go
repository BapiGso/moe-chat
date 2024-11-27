package mymiddleware

import (
	"fmt"
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
		fmt.Println(err)
		return c.Redirect(http.StatusFound, "/login")
	},
	SigningKey:  []byte(strconv.Itoa(os.Getpid())),
	TokenLookup: "cookie:moe-chat_token",
	Skipper: func(c echo.Context) bool {
		skipPath := c.Path() == "/login" || c.Path() == "/register" || strings.HasPrefix(c.Path(), "/assets/")
		if debug {
			c.Set("email", "debug")
		}
		return skipPath || debug
	},
}.ToMiddleware()
