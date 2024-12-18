package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"moechat/core/api"
	"moechat/core/api/part"
	"net/http"
)

// Completion 用于将会话发送给ai todo 可能需要池化
// 前端要发的数据，什么公司的什么模型model，
func Completion(c echo.Context) error {
	req := new(part.Completion)
	if err := c.Bind(req); err != nil {
		return err
	}
	switch c.Request().Method {
	case http.MethodPost:
		apiAdapter, err := api.New(req.Provider)
		if err != nil {
			return err
		}
		if err := apiAdapter.CreateResStream(c, req); err != nil {
			return err
		}
		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		buf := make([]byte, 512)
		for {
			select {
			case <-c.Request().Context().Done():

			default:
				n, err := apiAdapter.Read(buf)
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				}
				if n > 0 {
					fmt.Fprintf(w, "data: %q\n\n", buf[:n])
					w.Flush()
				}
			}
		}
	case http.MethodPut:
		//case http.MethodGet:
	}
	return echo.ErrMethodNotAllowed
}
