package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpl *template.Template
}

func newTemplate() *Template {
    return &Template{
        tmpl: template.Must(template.ParseGlob("views/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Page struct {
    Title string
}

func newPage(title string) Page {
    return Page{
        Title: title,
    }
}

func main() {
    e := echo.New()

    e.Renderer = newTemplate()
    e.Use(middleware.Logger())

    e.Static("/images", "images")
    e.Static("/css", "css")

    page := newPage("Welcome")

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", page)
    })

    e.Logger.Fatal(e.Start(":42069"))
}
