package config

import (
	"encoding/json"
	"github.com/goccy/go-yaml"
	"os"
)

var Global config

type config struct {
	Port     string   `json:"port" yaml:"port"`
	Env      string   `json:"env" yaml:"env"`
	Database database `json:"database" yaml:"database"`
}

type database struct {
	Mongodb mongodbConfig `json:"mongodb" yaml:"mongodb"`
	Redis   redisConfig   `json:"redis" yaml:"redis"`
}

type mongodbConfig struct {
	URI       string          `json:"uri" yaml:"uri"`
	Instances mongodbInstance `json:"instances" yaml:"instances"`
}

type mongodbInstance map[string][]string

type redisConfig struct {
	URI       string        `json:"uri" yaml:"uri"`
	Instances redisInstance `json:"instances" yaml:"instances"`
}

type redisInstance map[string]int

func ReadJson(path string) error {
	fileread, _ := os.ReadFile(path)
	return json.Unmarshal(fileread, &Global)
}

func ReadYaml(path string) error {
	fileread, _ := os.ReadFile(path)
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
