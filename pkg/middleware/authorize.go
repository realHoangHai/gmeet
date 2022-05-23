package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/realHoangHai/gmeet-biz/pkg/config"
	"time"
)

func IsAuthenticated(config *config.Config) func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: config.Jwt.Secret,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "unauthorized",
			})
		},
	})
}

func GetUserIdFromContext(ctx *fiber.Ctx) (string, error) {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idStr := claims["id"].(string)
	return idStr, nil
}

func GenerateToken(id uuid.UUID) (string, error) {
	cfg := config.New()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // one week
	s, err := token.SignedString(cfg.Jwt.Secret)
	return s, err
}
