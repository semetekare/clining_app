package app

import (
	"clining_app/app/config"
	"clining_app/app/router"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	httpServer *http.Server
}

func (a *App) Run(port string) {
	// Подключение к базе данных Irgups
	config.ConnectDatabasePostgre()

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "EventService",
		AppName:       "EventService IrGUPS v 1.0",
	})
	app.Use(logger.New(), recover.New(), cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		//AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	
	app.Route(
		"/",
		router.RegisterHTTPEndpoints,
		"abit.",
	)

	log.Fatal(app.Listen(":" + port))
}