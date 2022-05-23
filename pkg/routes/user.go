package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realHoangHai/gmeet-biz/pkg/handlers"
	"github.com/realHoangHai/gmeet-biz/pkg/middleware"
)

func SetupUserRoutes(r fiber.Router, handler *handlers.Handlers) {
	user := r.Group("/user")
	{
		user.Post("/register", handler.UserRegister)
		user.Post("/login", handler.UserLogin)

		user.Use(middleware.IsAuthenticated(nil))
		user.Get("/profile", handler.Profile)
	}
}
