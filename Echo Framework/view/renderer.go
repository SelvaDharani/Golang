package view

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4" // Import the correct version of Echo
)

// TemplateRenderer is a custom template renderer for Echo to use HTML templates
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer initializes the TemplateRenderer with the templates.
func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

// Render renders the specified template with the given data to the io.Writer.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}
