package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realHoangHai/gmeet-biz/pkg/handlers"
)

func SetupApiV1(app *fiber.App, handler *handlers.Handlers) {
	v1 := app.Group("/api/v1")
	SetupUserRoutes(v1, handler)
}
