package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/vrischmann/envconfig"
)

// Config type
type Config struct {
	Env            string `envconfig:"default=development"`
	Port           int    `envconfig:"default=8080"`
	Host           string `envconfig:"default=127.0.0.1"`
	ExecutablePath string `envconfig:"default=cjpeg"`
	Quality        int    `envconfig:"default=80"`
}

func mustConfigure() {
	if err := envconfig.Init(&config); err != nil {
		panic(err)
	}
}
