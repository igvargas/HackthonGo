package models

type Products struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
