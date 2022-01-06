package models

type Customer struct {
	ID        int64  `json:"id"`
	Nombre    string `json:"first_name"`
	Apellido  string `json:"last_name"`
	Condicion string `json:"condition"`
}
