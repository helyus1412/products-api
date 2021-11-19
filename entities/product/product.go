package entities

import "time"

type Product struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Description string    `json:"description"`
	Quantity    string    `json:"quantity"`
}

type ProductInput struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
