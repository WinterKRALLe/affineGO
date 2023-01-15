package main

import (
	"io"
	"net/http"
	"strconv"
	"text/template"

	"affine/affine"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, ok := data.(map[string]interface{}); ok {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

type Data struct {
	Input string `json:"input"`
	A     string `json:"a"`
	B     string `json:"b"`
}

func encryptDecryptHandler(c echo.Context) error {
	data := new(Data)
	if err := c.Bind(data); err != nil {
		return err
	}

	a, _ := strconv.Atoi(data.A)
	b, _ := strconv.Atoi(data.B)

	var result string
	if c.Request().URL.Path == "/encrypt" {
		result = affine.Encrypt(data.Input, a, b)
	} else if c.Request().URL.Path == "/decrypt" {
		result = affine.Decrypt(data.Input, a, b)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}

func main() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.POST("/encrypt", encryptDecryptHandler)
	e.POST("/decrypt", encryptDecryptHandler)

	e.Start(":8080")
}
