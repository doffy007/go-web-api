package product

import "encoding/json"

type ProductRequest struct {
	Jenis       string      `json:"jenis" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required,numeric,min=1,max=5"`
	PriceBefore json.Number `json:"price_before" binding:"required,numeric"`
	PriceToday  json.Number `json:"price_today" binding:"required,numeric"`
}