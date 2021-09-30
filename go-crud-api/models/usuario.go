package models

type Usuario struct {
	Id uint `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}