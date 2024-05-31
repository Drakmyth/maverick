package main

import (
	"bytes"
	"context"
	"embed"
	"html/template"
	"io"
	"log"
	"log/slog"
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
	ctx context.Context
}

//go:embed templates/*.tmpl.html
var tmplFS embed.FS

var baseTmpl *template.Template = nil

// NewApp creates a new App application struct
func NewApp() *App {
	baseTmpl = template.Must(template.ParseFS(tmplFS, "templates/*.tmpl.html"))

	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func render(w io.Writer, name string, data any) {
	tmpl := template.Must(baseTmpl.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "templates/"+name))
	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Fatal("error rendering template", err)
	}
}

func (a *App) GetHomePage() string {
	slog.Info("Getting Home")

	pagedata := struct {
		Header PageHeader
	}{
		Header: PageHeader{
			Title:       "Home",
			Description: "This is your new home.",
		},
	}

	var tpl bytes.Buffer
	render(&tpl, "home-page.tmpl.html", &pagedata)
	return tpl.String()
}

func (a *App) GetEnginesPage() string {
	slog.Info("Getting Engines")

	pagedata := struct {
		Header PageHeader
	}{
		Header: PageHeader{
			Title:       "Engines",
			Description: "This is where you define your source ports.",
		},
	}

	var tpl bytes.Buffer
	render(&tpl, "engines-page.tmpl.html", &pagedata)
	return tpl.String()
}

func (a *App) GetIWADsPage() string {
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

	var tpl bytes.Buffer
	render(&tpl, "iwads-page.tmpl.html", &pagedata)
	return tpl.String()
}
