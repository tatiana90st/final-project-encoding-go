package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var data []byte
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	var dockerCompose models.DockerCompose
	j.DockerCompose = &dockerCompose
	err = json.Unmarshal(data, j.DockerCompose)
	if err != nil {
		return err
	}
	out, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(j.FileOutput, out, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var data []byte
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	var dockerCompose models.DockerCompose
	y.DockerCompose = &dockerCompose
	err = yaml.Unmarshal(data, y.DockerCompose)
	if err != nil {
		return err
	}
	out, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}
	err = os.WriteFile(y.FileOutput, out, 0644)
	if err != nil {
		return err
	}
	return nil
}
