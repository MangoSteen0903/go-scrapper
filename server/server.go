package server

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MangoSteen0903/go-scrapper/scrap"
	"github.com/MangoSteen0903/go-scrapper/utils"
	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

const FILE_NAME = "output.csv"

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
func handleHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", "")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILE_NAME)
	company := c.FormValue("company")
	company = strings.ToLower(utils.CleanString(company))
	scrap.GitScrapper(company)
	userFileName := fmt.Sprintf("repository_%s.csv", company)
	return c.Attachment(FILE_NAME, userFileName)
}
func Start() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseFiles("server/home.html")),
	}
	e.Renderer = renderer
	e.GET("/", handleHome)
	e.POST("/scrap", handleScrape)
	e.Logger.Fatal(e.Start(":5000"))
}
