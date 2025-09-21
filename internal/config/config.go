package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	NumFile1 string `yaml:"num_file1"`
	NumFile2 string `yaml:"num_file2"`
	NumSize  int    `yaml:"num_size"`
	Pow      int64  `yaml:"pow"`
	OutDir   string `yaml:"out_dir"`
}

func LoadConfig(path string) AppConfig {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Ошибка чтения конфигурации %s: %v", path, err)
	}
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Ошибка парсинга YAML %s: %v", path, err)
	}
	return cfg
}
