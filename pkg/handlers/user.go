package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/realHoangHai/gmeet-biz/ent/user"
	"github.com/realHoangHai/gmeet-biz/pkg/log"
	"github.com/realHoangHai/gmeet-biz/pkg/middleware"
	"github.com/realHoangHai/gmeet-biz/pkg/utils"
)

func (h *Handlers) UserRegister(ctx *fiber.Ctx) error {
	var request registerRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := request.Validate(); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	exist, _ := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if exist != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Email already exists",
		})
	}

	hassPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Errorf("Failed hashing password: %v", err)
		return nil
	}

	_, err = h.Client.User.Create().
		SetEmail(request.Email).
		SetPassword(hassPassword).
		SetFirstName(request.FirstName).
		SetLastName(request.LastName).
		SetAvatar(request.Avatar).
		Save(ctx.Context())
	if err != nil {
		log.Errorf("Failed to create user: %v", err)
		return ctx.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	_ = ctx.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "register successfully",
	})
	return nil
}

func (h *Handlers) UserLogin(ctx *fiber.Ctx) error {
	var request loginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	u, err := h.Client.User.Query().Where(user.Email(request.Email)).Only(ctx.Context())
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := utils.ComparePassword(u.Password, request.Password); err != nil {
		log.Errorf("Failed to compare password: %v", err)
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid password",
		})
	}

	token, err := middleware.GenerateToken(u.ID)
	if err != nil {
		log.Errorf("Failed to generate token: %v", err)
		return nil
	}

	return ctx.Status(200).JSON(fiber.Map{
		"error": false,
		"data": userResponse{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Avatar:    u.Avatar,
		},
		"token": token,
	})
}

func (h *Handlers) Profile(ctx *fiber.Ctx) error {
	userId, err := middleware.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.Status(401)
		return nil
	}

	uid, _ := uuid.Parse(userId)

	u, err := h.Client.User.Query().Where(user.ID(uid)).Only(ctx.Context())
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		return ctx.Status(400).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"error": false,
		"data": userResponse{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Avatar:    u.Avatar,
		},
	})
}
