package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type PageHeader struct {
	Title       string
	Description string
}

type IWADDef struct {
	Name string
	Path string
}

type IWADsConfig struct {
	IWADs []IWADDef
}

// App struct
type App struct {
	ctx   context.Context
	iwads []IWADDef
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
	a.iwads = loadIWADs()
}

func getOrCreateConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	configDir = filepath.Join(configDir, "Maverick Launcher")
	err = os.MkdirAll(configDir, 0644)
	if err != nil {
		return configDir, err
	}

	return configDir, nil
}

func loadIWADs() []IWADDef {
	configDir, err := getOrCreateConfigDir()
	check(err)

	iwadConfigFilePath := filepath.Join(configDir, "iwads.json")
	data, err := os.ReadFile(iwadConfigFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			createIWADConfigFile()
			return []IWADDef{}
		} else {
			panic(err)
		}
	}

	var config IWADsConfig
	err = json.Unmarshal(data, &config)
	check(err)

	return config.IWADs
}

func createIWADConfigFile() {
	config := IWADsConfig{
		IWADs: []IWADDef{},
	}

	data, err := json.Marshal(config)
	check(err)

	configDir, err := getOrCreateConfigDir()
	check(err)

	iwadConfigFilePath := filepath.Join(configDir, "iwads.json")
	f, err := os.Create(iwadConfigFilePath)
	check(err)
	defer f.Close()

	_, err = f.Write(data)
	check(err)
}

func render(w io.Writer, name string, data any) {
	tmpl := template.Must(baseTmpl.Clone())
	tmpl = template.Must(tmpl.ParseFS(tmplFS, "templates/"+name))
	err := tmpl.ExecuteTemplate(w, name, data)
	check(err)
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

func (a *App) SelectIWADFile() string {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select IWAD",
	})
	check(err)

	return path
}

func (a *App) SaveIWAD(iwad IWADDef) {
	print("TODO: Go - Implement IWAD Save", iwad.Name, iwad.Path)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
