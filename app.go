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

	var tpl bytes.Buffer
	render(&tpl, "home-page.tmpl.html", nil)
	return tpl.String()
}

func (a *App) GetEnginesPage() string {
	slog.Info("Getting Engines")

	var tpl bytes.Buffer
	render(&tpl, "engines-page.tmpl.html", nil)
	return tpl.String()
}

func (a *App) GetAddIWADModal() string {
	slog.Info("Getting Add IWAD Modal")

	var tpl bytes.Buffer
	render(&tpl, "add-iwad-modal.tmpl.html", nil)
	return tpl.String()
}

func (a *App) GetIWADsPage() string {
	slog.Info("Getting IWADs")

	pagedata := struct {
		IWADs []IWADDef
	}{
		IWADs: []IWADDef{},
	}

	var tpl bytes.Buffer
	render(&tpl, "iwads-page.tmpl.html", &pagedata)
	return tpl.String()
}

func (a *App) GetPageTitle(page string) string {
	switch page {
	case "home":
		return "Home"
	case "engines":
		return "Engines"
	case "iwads":
		return "IWADs"
	}

	return ""
}

func (a *App) GetPageDescription(page string) string {
	switch page {
	case "home":
		return "This is your new home."
	case "engines":
		return "This is where you define your source ports."
	case "iwads":
		return "Set the location of base game content."
	}

	return ""
}
