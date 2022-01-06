package models

type Sales struct {
	ID       int      `json:"id"`
	Cuantity float64  `json:"cuantity"`
	Invoice  Invoice  `json:"invoice"`
	Produc   Products `json:"product"`
}
