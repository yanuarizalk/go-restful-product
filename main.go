package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.yanuarizal.net/go-restful-product/config"
	"github.yanuarizal.net/go-restful-product/database"
	"github.yanuarizal.net/go-restful-product/router"

	_ "github.com/joho/godotenv/autoload"
	_ "github.yanuarizal.net/go-restful-product/docs"
)

// @title           Go Restful Test
// @version         1.0
// @description     This is a sample restful api with fiber, gorm & testing.

// @contact.name   Yanuarizal K
// @contact.url    https://www.yanuarizal.net
// @contact.email  me@yanuarizal.net

// @host      localhost:3000
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := database.Connect(config.Database)
	if err != nil {
		panic(err)
	}

	app := fiber.New(config.Fiber)
	app.Use(logger.New(), recover.New())

	router.RegisterRoutes(app)

	start(app)
}

func start(app *fiber.App) {
	conn := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Println(err)
		}

		close(conn)
	}()

	if err := app.Listen(config.App.ListenAddr); err != nil {
		log.Println(err)
	}

	<-conn
}
