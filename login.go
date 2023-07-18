package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if len(user.Username) < 3 || len(user.Password) < 3 {
		return c.SendString("Username or password is empty or too short")
	}

	var foundedUsers []string

	for _, u := range users {
		if u[0] == user.Username && u[1] == user.Password {
			foundedUsers = u
		}
	}

	if len(foundedUsers) == 0 {
		return c.SendString("Username or password is incorrect")
	}

	var ronaldo bool = foundedUsers[0] == "admin"

	claims := jwt.MapClaims{
		"name":  foundedUsers[0],
		"admin": ronaldo,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})

}
