package delivery

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	// Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}
