package router

import (
	controllers "clining_app/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router) {
	// OK запрос
	router.Get("/Ok", controllers.Ok)
}
