package model

import "time"

type Promotion struct {
	Id              string    `json:"id"`
	Price           float64   `json:"price"`
	Expiration_date time.Time `json:"expiration_date"`
}
