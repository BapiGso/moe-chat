package core

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"moechat/core/handler"
	"moechat/core/mymiddleware"
	"net"
	"net/http"
)

func (c *Core) Route() {
	c.e.Validator = mymiddleware.DefaultValidator
	c.e.Renderer = mymiddleware.DefaultTemplateRender

	c.e.Use(mymiddleware.Slog)
	c.e.Use(mymiddleware.JWT) //登录逻辑也在这里面
	c.e.Use(middleware.Recover())
	c.e.Use(middleware.Gzip())

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

	//c.e.Any("/", handler.Login, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(3)))
	c.e.Any("/chat", handler.Chat)
	admin := c.e.Group("/admin")
	admin.Any("", handler.Admin)
	admin.Any("/user", handler.User)
	admin.Any("/model", handler.User)
	showPanelAddr()
	//c.e.StartTLS(viper.GetString("panel.port"), []byte(certPEM), []byte(keyPEM))
	c.e.Start(":8080")
}

func showPanelAddr() {
	// Attempt to show IP address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		} else if ipNet.IP.To4() != nil && ipNet.IP.IsGlobalUnicast() {
			//ip := ipNet.IP.To4()
			//if ip[0] == 10 || (ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) || (ip[0] == 192 && ip[1] == 168) {
			//	continue
			//}
			fmt.Printf("moe-chat started on http://%v%v/\n", ipNet.IP, ":8080")
		} else if ipNet.IP.To16() != nil && ipNet.IP.IsGlobalUnicast() {
			// Check for IPv6 unicast addresses
			fmt.Printf("moe-chat started on http://[%v]%v/\n", ipNet.IP, ":8080")
		}
	}
}
