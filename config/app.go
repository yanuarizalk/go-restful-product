package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type app struct {
	Env        string
	ListenAddr string
}

var App app

var Fiber = fiber.Config{}

func init() {
	Fiber.AppName = os.Getenv("APP_NAME")
	App.Env = os.Getenv("ENV")
	App.ListenAddr = os.Getenv("LISTEN_ADDR")

	if App.ListenAddr == "" {
		App.ListenAddr = ":3000"
	}
}

func (data app) IsDevelopment() bool {
	switch strings.ToLower(data.Env) {
	case "dev", "development", "test":
		return true
	}

	return false
}
