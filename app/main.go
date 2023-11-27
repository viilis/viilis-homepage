package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/viilis/app/utils"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type TemplateMain struct {
	Content TemplateContent
	Meta    TemplateMeta
}

type TemplateLink struct {
	Title string
	Url   string
}
type TemplateContent struct {
	Head    string
	Content string
	Links   []TemplateLink
}

type TemplateMeta struct {
	Title       string
	Description string
}

func main() {

	// load config
	utils.InitConfig()

	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "assets")

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("assets/templates/*.html")),
	}

	e.GET("/", func(c echo.Context) error {

		pageContent := TemplateMain{
			Content: TemplateContent{
				"Ville Lindberg",
				`Software Engineer | Co-Founder of Kvanttori | CS Student`,
				[]TemplateLink{
					{
						"LinkedIn",
						"https://www.linkedin.com/in/ville-l-4a6398206/",
					},
					{
						"Kvanttori Oy",
						"https://kvanttori.fi",
					},
					{
						"GitHub",
						"https://github.com/viilis",
					}
				},
			},
			Meta: TemplateMeta{
				"Homepage",
				"Nice and good SEO desc here",
			},
		}

		return c.Render(http.StatusOK, "main", pageContent)
	})

	e.Logger.Fatal(e.Start(":" + utils.Config.Port))
}
