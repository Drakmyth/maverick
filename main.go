package main

import (
	"embed"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

type PageHeader struct {
	Title       string
	Description string
}

type IWADDef struct {
	Name string
	Path string
}

//go:embed all:frontend/dist
var assets embed.FS

var tmpl *template.Template = nil

func main() {
	// Create an instance of the app structure
	app := NewApp()
	r := chi.NewRouter()
	r.Get("/home", getHomePage)
	r.Get("/engines", getEnginesPage)
	r.Get("/iwads", getIWADsPage)

	var err error = nil
	tmpl, err = template.ParseGlob("./templates/*.tmpl.html")
	if err != nil {
		log.Fatal("Error loading templates:" + err.Error())
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "Maverick",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: func(next http.Handler) http.Handler {
				r.NotFound(next.ServeHTTP)
				return r
			},
		},
		Windows: &windows.Options{
			BackdropType:         windows.Tabbed,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
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

	tmpl.ExecuteTemplate(w, "home-page", &pagedata)
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

	tmpl.ExecuteTemplate(w, "engines-page", &pagedata)
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

	tmpl.ExecuteTemplate(w, "iwads-page", &pagedata)
}
