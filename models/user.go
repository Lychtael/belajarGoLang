package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	Provider   string `json:"provider"`
	Birthdate  string `json:"birthdate"`
	Birthplace string `json:"birthplace"`
	Contact    string `json:"contact"`
	Address    string `json:"address"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}
