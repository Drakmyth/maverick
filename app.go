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

type IWADOptionsModalData struct {
	IWADId          string
	MoveUpEnabled   bool
	MoveDownEnabled bool
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
	return a.GetContentWithData(name, nil)
}

func (a *App) GetContentWithData(name string, data any) string {
	var tpl bytes.Buffer
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/"+name+".tmpl.html"))

	tmplData := data
	if tmplData == nil {
		tmplData = a
	}

	err := tmpl.Execute(&tpl, tmplData)
	if err != nil {
		panic(err)
	}

	return tpl.String()
}

func (a *App) GetIWADOptionsModal(iwadId string) string {
	i, _ := a.IWADs.FindIndexOf(iwadId)
	iwadCnt := len(a.IWADs)

	data := IWADOptionsModalData{
		IWADId:          iwadId,
		MoveUpEnabled:   iwadCnt > 1 && i > 0,
		MoveDownEnabled: iwadCnt > 1 && i < iwadCnt-1 && i > -1,
	}

	return a.GetContentWithData("iwad-options-modal", data)
}

func (a *App) GetRemoveIWADModal(iwadId string) string {
	iwad, err := a.IWADs.FindIWAD(iwadId)
	if err != nil {
		panic(err)
	}

	return a.GetContentWithData("remove-iwad-modal", iwad)
}

func (a *App) MoveIWADUp(iwadId string) bool {
	i, err := a.IWADs.FindIndexOf(iwadId)
	if err != nil {
		return false
	}

	err = a.IWADs.MoveIWAD(iwadId, i-1)
	if err != nil {
		return false
	}

	err = a.saveIWADConfigFile()
	return err == nil
}

func (a *App) MoveIWADDown(iwadId string) bool {
	i, err := a.IWADs.FindIndexOf(iwadId)
	if err != nil {
		return false
	}

	err = a.IWADs.MoveIWAD(iwadId, i+1)
	if err != nil {
		return false
	}

	err = a.saveIWADConfigFile()
	return err == nil
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

func (a *App) SaveIWAD(name string, path string) {
	a.IWADs = append(a.IWADs, iwads.NewIWAD(name, path))
	a.saveIWADConfigFile()
}

func (a *App) RemoveIWAD(iwadId string) bool {
	err := a.IWADs.RemoveIWAD(iwadId)
	if err != nil {
		panic(err)
	}

	err = a.saveIWADConfigFile()
	return err == nil
}

func (a *App) saveIWADConfigFile() error {
	iwadConfigPath := filepath.Join(a.configDirectoryPath, IWAD_CONFIG_FILENAME)
	return a.IWADs.SaveToFile(iwadConfigPath)
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
