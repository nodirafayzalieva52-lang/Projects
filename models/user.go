package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct{
	Message string `json:"message"`
}