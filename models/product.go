package models

import "time"

type Product struct {
	Id              int         `json:"id"`
	Sku             string      `json:"sku"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Category        string      `json:"category"`
	Etalase         string      `json:"etalase"`
	Weight          float32     `json:"weight"`
	Price           float32     `json:"price"`
	AddedTime       time.Time   `json:"added_time"`
	AverageRating   float32     `json:"average_rating"`
}