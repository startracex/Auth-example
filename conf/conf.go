package conf

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
	"os"
)

var Global globalConfig

type globalConfig struct {
	Name     string       `json:"name" yaml:"name"`
	Host     string       `json:"host" yaml:"host"`
	Port     string       `json:"port" yaml:"port"`
	Webkey   string       `json:"webkey" yaml:"webkey"`
	Env      string       `json:"env" yaml:"env"`
	Allow    []string     `json:"allow" yaml:"allow"`
	Database databaseType `json:"database" yaml:"database"`
}

type databaseType struct {
	Mongodb mongodbConfigType `json:"mongodb" yaml:"mongodb"`
	Redis   redisConfigType   `json:"redis" yaml:"redis"`
}

type mongodbConfigType struct {
	URI       string              `json:"uri" yaml:"uri"`
	Instances map[string][]string `json:"instances" yaml:"instances"`
}

type redisConfigType struct {
	URI       string         `json:"uri" yaml:"uri"`
	Instances map[string]int `json:"instances" yaml:"instances"`
}

func ReadJson(path string) error {
	fileread, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileread, &Global)
}

func ReadYaml(path string) error {
	fileread, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(fileread, &Global)
}

func LoadConfig(path string) {
	err := ReadJson(path)
	if err != nil {
		err = ReadYaml(path)
		if err != nil {
			panic(err)
		}
	}
}
