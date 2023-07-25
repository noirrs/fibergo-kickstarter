package services

import (
	. "fibergo-kickstarter/database"
	. "fibergo-kickstarter/types"

	"context"
	"fmt"
	"math/rand" // for id_generator func
	"strconv"   // for int to string (in id_generator func)
	"time"      // for rand.Seed(time.Now().Unix()) (in id_generator func) and context expiration (in unique_id_generator and main func)

	"github.com/go-playground/validator/v10" // for validate request's body
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ! add max 10 sec per thread https://github.com/Mr-Malomz/fiber-mongo-api/blob/main/controllers/user_controller.go
func Register(c *fiber.Ctx) error {
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

	exist, error := user_collection.FindOne(ctx, bson.D{{Key: "username", Value: user.Username}}).DecodeBytes()

	if exist != nil && error != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusConflict).JSON(Response{Status: fiber.StatusConflict, Message: "Username Already Exist", Data: &fiber.Map{"data": error}})
	}

	user.ID = unique_id_generator()
	user.Words = []string{}
	user.Admin = false

	_, err := user_collection.InsertOne(ctx, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{Status: fiber.StatusInternalServerError, Message: "Internal Server Error", Data: &fiber.Map{"data": err.Error()}})
	}
	user.Password = ""
	return c.Status(fiber.StatusCreated).JSON(UserCreatedResponse{Status: fiber.StatusCreated, Message: "Created", Data: &user})
}

// creates a loop while the id is not unique
func unique_id_generator() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var id string = id_generator()

	var user_collection = GetCollection("users")

	var exist = user_collection.FindOne(ctx, bson.D{{Key: "id", Value: id}})

	for exist.Err() == nil {
		fmt.Println("id already exist")
		id = id_generator()
		exist = user_collection.FindOne(ctx, bson.D{{Key: "id", Value: id}})
	}

	return id
}

// creates a random number with 8 digits
func id_generator() string {

	rand.Seed(time.Now().Unix())

	var id string

	for i := 0; i < 8; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}
	return id
}
