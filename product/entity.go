package product

import "time"

type Product struct {
	ID          int
	Jenis        string
	Name        string
	Description string
	Rating      int
	PriceBefore int
	PriceToday  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
