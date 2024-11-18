package mymiddleware

import (
	"github.com/labstack/echo/v4"
	"io"
	"moechat/assets"
	"text/template"
)

var DefaultTemplateRender = &TemplateRender{
	Template: template.Must(template.ParseFS(assets.Assets, "*.html")),
}

type TemplateRender struct {
	*template.Template //渲染模板
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}
