package services

import "github.com/gofiber/fiber/v2"

func Appie(c *fiber.Ctx) error {

	return c.SendString("Hello, World 👋!")
}
