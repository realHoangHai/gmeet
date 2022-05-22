package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func SetupApiV1(app *fiber.App) {
	v1 := app.Group("/api/v1")
	{
		v1.Get("/hello", func(c *fiber.Ctx) error {
			fmt.Println("Hello World")
			return c.SendString("Hello World")
		})
	}
}
