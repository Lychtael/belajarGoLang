package models

type Provider struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Status  string `json:"status"`
	Address string `json:"address"`
}
