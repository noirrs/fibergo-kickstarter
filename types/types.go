package types

import "github.com/gofiber/fiber/v2"



type User struct {
	ID       string   `json:"id"`
	Admin    bool     `json:"admin"`
	Username string   `json:"username" validate:"required"`
	Password string   `json:"password" validate:"required"`
	Words    []string `json:"words"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Preferences struct {
	Version string `json:"version"`
	Status  string `json:"status"`
	Dbname  string `json:"dbname"`
	Port    string `json:"port"`
	ConnURL string `json:"connURL"`
}

type Response struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	Data    *fiber.Map `json:"data"`
}

type UserCreatedResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    *User  `json:"data"`
}

/*
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
*/
