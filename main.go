package main

import (
	. "fibergo-kickstarter/controllers"

	"github.com/gofiber/fiber/v2"
	// . "fibergo-example/database"
)

func main() {
	app := fiber.New()

	Controller(app)

	app.Listen(":3000")

}
