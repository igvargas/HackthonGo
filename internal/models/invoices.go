package models

type Invoice struct {
	ID          int64   `json:"id"`
	Date        string  `json:"date_time"`
	ID_customer int64   `json:"id_customer"`
	Total       float64 `json:"total"`
}
