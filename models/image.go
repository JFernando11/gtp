package models

type Image struct {
	Id         	int     `json:"id"`
	ProductId  	int     `json:"product_id"`
	ImageUrl   	string  `json:"image_url"`
	Description string  `json:"description"`
}