package main

type Request struct {
	Message string `json:"message"`
	Token   string `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
