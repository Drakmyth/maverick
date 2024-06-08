package engines

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errNotFound()
)

func errNotFound() error { return errors.New("engine not found") }

type EngineCollection []EngineDefinition
type EngineDefinition struct {
	Id   string
	Name string
	Path string
}

type EngineConfigFile struct {
	Engines EngineCollection
}

func NewEngine(name string, path string) EngineDefinition {
	return EngineDefinition{
		Id:   uuid.NewString(),
		Name: name,
		Path: path,
	}
}

func ReadEngineConfigFile(path string) (EngineCollection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return EngineCollection{}, err
	}

	var config EngineConfigFile
	err = json.Unmarshal(data, &config)
	if err != nil {
		return EngineCollection{}, err
	}

	return config.Engines, nil
}

func (ec EngineCollection) FindIndexOf(engineId string) (int, error) {
	for i, engine := range ec {
		if engine.Id == engineId {
			return i, nil
		}
	}

	return -1, errNotFound()
}

func (ec EngineCollection) FindEngine(engineId string) (EngineDefinition, error) {
	for _, engine := range ec {
		if engine.Id == engineId {
			return engine, nil
		}
	}

	return EngineDefinition{}, errNotFound()
}

func (ec *EngineCollection) RemoveEngine(engineId string) error {
	i, err := ec.FindIndexOf(engineId)
	if err != nil {
		return err
	}

	*ec = append((*ec)[:i], (*ec)[i+1:]...)
	return nil
}

func (ec *EngineCollection) MoveEngine(engineId string, destinationIndex int) error {
	i, err := ec.FindIndexOf(engineId)
	if err != nil {
		return err
	}

	engine := (*ec)[i]
	*ec = append((*ec)[:i], (*ec)[i+1:]...)
	*ec = append((*ec)[:destinationIndex+1], (*ec)[destinationIndex:]...)
	(*ec)[destinationIndex] = engine
	return nil
}

func (ec EngineCollection) SaveToFile(path string) error {
	config := EngineConfigFile{
		Engines: ec,
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
