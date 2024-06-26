package iwads

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errNotFound()
)

func errNotFound() error { return errors.New("iwad not found") }

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
		Id:   uuid.NewString(),
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

func (ic IWADCollection) FindIndexOf(iwadId string) (int, error) {
	for i, iwad := range ic {
		if iwad.Id == iwadId {
			return i, nil
		}
	}

	return -1, errNotFound()
}

func (ic IWADCollection) FindIWAD(iwadId string) (IWADDefinition, error) {
	for _, iwad := range ic {
		if iwad.Id == iwadId {
			return iwad, nil
		}
	}

	return IWADDefinition{}, errNotFound()
}

func (ic *IWADCollection) RemoveIWAD(iwadId string) error {
	i, err := ic.FindIndexOf(iwadId)
	if err != nil {
		return err
	}

	*ic = append((*ic)[:i], (*ic)[i+1:]...)
	return nil
}

func (ic *IWADCollection) MoveIWAD(iwadId string, destinationIndex int) error {
	i, err := ic.FindIndexOf(iwadId)
	if err != nil {
		return err
	}

	iwad := (*ic)[i]
	*ic = append((*ic)[:i], (*ic)[i+1:]...)
	*ic = append((*ic)[:destinationIndex+1], (*ic)[destinationIndex:]...)
	(*ic)[destinationIndex] = iwad
	return nil
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
