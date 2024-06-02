package iwads

import (
	"encoding/json"
	"os"
)

type IWADCollection []IWADDefinition
type IWADDefinition struct {
	Id   string
	Name string
	Path string
}

type IWADConfigFile struct {
	IWADs IWADCollection
}

func NewIWAD(name string, path string) IWADDefinition {
	return IWADDefinition{
		Id:   "0",
		Name: name,
		Path: path,
	}
}

func ReadIWADConfigFile(path string) (IWADCollection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return IWADCollection{}, err
	}

	var config IWADConfigFile
	err = json.Unmarshal(data, &config)
	if err != nil {
		return IWADCollection{}, err
	}

	return config.IWADs, nil
}

func (ic IWADCollection) SaveToFile(path string) error {
	config := IWADConfigFile{
		IWADs: ic,
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return f.Sync()
}
