package services

import (
	. "fibergo-kickstarter/database"
	. "fibergo-kickstarter/types"

	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var validate = validator.New()

	user_collection := GetCollection("users")

	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Status: fiber.StatusBadRequest, Message: "Bad Request", Data: &fiber.Map{"error": err.Error()}})
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Status: fiber.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	if len(user.Username) < 3 || len(user.Password) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(Response{Status: fiber.StatusBadRequest, Message: "Bad Request", Data: &fiber.Map{"error": "Username or password is empty or too short"}})
	}

	err := user_collection.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}, {Key: "password", Value: user.Password}}).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(Response{Status: fiber.StatusNotFound, Message: "User Not Found", Data: &fiber.Map{"data": err}})
	}

	if user.Password != user.Password || user.Username != user.Username {
		return c.Status(fiber.StatusUnauthorized).JSON(Response{Status: fiber.StatusUnauthorized, Message: "Unauthorized", Data: &fiber.Map{"data": err}})
	}

	claims := jwt.MapClaims{
		"admin": user.Admin,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Status: fiber.StatusInternalServerError, Message: "Internal Server Error", Data: &fiber.Map{"data": err}})
	}

	return c.JSON(fiber.Map{"token": t})

}
