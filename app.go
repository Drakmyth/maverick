package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type PageHeader struct {
	Title       string
	Description string
}

type IWADDef struct {
	Name string
	Path string
}

// App struct
type App struct {
	ctx    context.Context
	router *chi.Mux
}

//go:embed templates/*.tmpl.html
var tmplFS embed.FS

var baseTmpl *template.Template = nil

func (a *App) middleware(next http.Handler) http.Handler {
	a.router.NotFound(next.ServeHTTP)
	return a.router
}

// NewApp creates a new App application struct
func NewApp() *App {
	r := chi.NewRouter()
	r.Get("/home", getHomePage)
	r.Get("/engines", getEnginesPage)
	r.Get("/iwads", getIWADsPage)

	baseTmpl = template.Must(template.ParseFS(tmplFS, "templates/*.tmpl.html"))

	return &App{
		router: r,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func render(w http.ResponseWriter, name string, data any) {
	tmpl := template.Must(baseTmpl.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "templates/"+name))
	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal("error rendering template", err)
	}
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting Home")

	pagedata := struct {
		Header PageHeader
	}{
		Header: PageHeader{
			Title:       "Home",
			Description: "This is your new home.",
		},
	}

	render(w, "home-page.tmpl.html", &pagedata)
}

func getEnginesPage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting Engines")

	pagedata := struct {
		Header PageHeader
	}{
		Header: PageHeader{
			Title:       "Engines",
			Description: "This is where you define your source ports.",
		},
	}

	render(w, "engines-page.tmpl.html", &pagedata)
}

func getIWADsPage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting IWADs")

	pagedata := struct {
		Header PageHeader
		IWADs  []IWADDef
	}{
		Header: PageHeader{
			Title:       "IWADs",
			Description: "Set the location of base game content.",
		},
		IWADs: []IWADDef{},
	}

	render(w, "iwads-page.tmpl.html", &pagedata)
}
