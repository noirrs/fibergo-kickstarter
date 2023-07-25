package controller

import (
	. "fibergo-kickstarter/services"

	"github.com/gofiber/fiber/v2"
)

func Controller(app *fiber.App) {

	/* JWT Unauthenticated Routes */
	app.Get("/", Appie)
	app.Post("/register", Register)
	app.Post("/login", Login)
	/* JWT Unauthenticated Routes */

	/* Middleware */
	Middleware(app) // if you want to use middleware with a specific group, you need to change type of app to fiber.Group in middleware.go
	/* Middleware */

	/* JWT Authenticated Routes */
	app.Post("/app", Appie)
	/* JWT Authenticated Routes */
}
