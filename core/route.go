package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"moechat/core/handler"
	"moechat/core/mymiddleware"
	"net/http"
)

func (c *Core) Route() {
	c.e.Validator = mymiddleware.DefaultValidator
	c.e.Renderer = mymiddleware.DefaultTemplateRender

	c.e.Use(mymiddleware.Slog)
	c.e.Use(middleware.Recover())
	c.e.Use(middleware.Gzip())

	c.e.HTTPErrorHandler = func(err error, c echo.Context) {
		c.JSON(400, err.Error())
	}
	// 静态资源
	c.e.Group("/assets", middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: func(c echo.Context) bool {
			c.Response().Header().Set("Cache-Control", "public, max-age=86400")
			return false
		},
		Filesystem: http.FS(c.assetsFS),
	}))

	//用于PWA的路径重写
	c.e.Pre(middleware.Rewrite(map[string]string{
		"/manifest.webmanifest": "/assets/manifest.webmanifest",
		"/sw.js":                "/assets/js/sw.js",
	}))

	c.e.Any("/", handler.Login, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)))
	c.e.Any("/c/:UUID/", handler.Chat)

	//c.e.StartTLS(viper.GetString("panel.port"), []byte(certPEM), []byte(keyPEM))
	c.e.Start(":8080")
}
