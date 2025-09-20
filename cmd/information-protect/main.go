package main

import (
	"information-protect/internal/app"
	"information-protect/internal/config"
)

func main() {
	cfg := config.LoadConfig("configs/config-local.yaml")
	app.Run(cfg)
}
