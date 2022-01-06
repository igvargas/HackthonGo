package models

type Customer struct {
	ID        int    `json:"id"`
	Nombre    string `json:"first_name"`
	Apellido  string `json:"last_name"`
	Condicion string `json:"condition"`
}
