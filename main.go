package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	jwtware "github.com/gofiber/contrib/jwt"
)

var users = [][]string{
	{"admin", "admin"},
	{"user", "user"},
	{"noir", "1234"},
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		Prefork:            true,
		CaseSensitive:      true,
		StrictRouting:      true,
		ServerHeader:       "Sola AI",
		AppName:            "Sola AI",
		BodyLimit:          4 * 1024 * 1024,
		Concurrency:        256 * 1024,
		EnableIPValidation: true,
		EnablePrintRoutes:  true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	v2 := app.Group("/v2")

	v2.Post("/login", Login)

	v2.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	v2.Use(func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)

		if !admin {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()

	})

	v2.Post("/app", App)

	app.Listen(":3000")
}

func App(c *fiber.Ctx) error {
	message := new(Request)

	if err := c.BodyParser(message); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if len(message.Message) < 1 {
		return c.SendString("Message is empty")
	}

	return c.SendString(message.Message)

}
