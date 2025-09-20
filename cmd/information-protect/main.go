package main

import (
	"information-protect/internal/app"
	"information-protect/internal/config"
)

var (
	ConfigPath string
)

func main() {
	if ConfigPath == "" {
		ConfigPath = "configs/config-local.yaml"
	}
	cfg := config.LoadConfig(ConfigPath)
	app.Run(cfg)
}
