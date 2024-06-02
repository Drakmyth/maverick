package main

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"html/template"
	"io/fs"
	"maverick/iwads"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const IWAD_CONFIG_FILENAME = "iwads.json"

//go:embed templates/*.tmpl.html
var tmplFS embed.FS

// App struct
type App struct {
	context             context.Context
	configDirectoryPath string
	IWADs               iwads.IWADCollection
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(context context.Context) {
	a.context = context
	configDir, err := getOrCreateConfigDirectory()
	if err != nil {
		panic(err)
	}
	a.configDirectoryPath = configDir

	iwadConfigPath := filepath.Join(configDir, IWAD_CONFIG_FILENAME)
	iwads, err := iwads.ReadIWADConfigFile(iwadConfigPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			iwads.SaveToFile(iwadConfigPath)
		}
	}
	a.IWADs = iwads
}

func (a *App) GetContent(name string) string {
	var tpl bytes.Buffer
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/"+name+".tmpl.html"))
	err := tmpl.Execute(&tpl, a)
	if err != nil {
		panic(err)
	}

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
	path, err := runtime.OpenFileDialog(a.context, runtime.OpenDialogOptions{
		Title: "Select IWAD",
	})
	if err != nil {
		panic(err)
	}

	return path
}

func (a *App) SaveIWAD(iwad iwads.IWADDefinition) {
	a.IWADs = append(a.IWADs, iwad)
	iwadConfigPath := filepath.Join(a.configDirectoryPath, IWAD_CONFIG_FILENAME)
	a.IWADs.SaveToFile(iwadConfigPath)
}

func getOrCreateConfigDirectory() (string, error) {
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
