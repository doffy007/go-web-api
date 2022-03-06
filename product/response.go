package product

type ProductResponse struct {
	ID          int    `json:"id"`
	Jenis       string `json:"jenis"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	PriceBefore int    `json:"price_before"`
	PriceToday  int    `json:"price_today"`
}
