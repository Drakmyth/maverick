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

	tmpl.ExecuteTemplate(w, "home-page", nil)
}

func getEnginesPage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting Engines")

	tmpl.ExecuteTemplate(w, "engines-page", nil)
}

func getIWADsPage(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting IWADs")

	tmpl.ExecuteTemplate(w, "iwads-page", nil)
}
