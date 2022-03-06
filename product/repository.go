package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	FindByJenis(Jenis string) ([]Product, error)
	FindByName(Name string) (Product, error)
	FindByRating(Rating int) ([]Product, error)
	FindByPriceBefore(PriceBefore int) ([]Product, error)
	FindByPriceToday(PriceToday int) ([]Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product

	err := r.db.Find(&product, ID).Error
	return product, err
}

func (r *repository) FindByJenis(Jenis string) ([]Product, error) {
	var products []Product

	err := r.db.Where("jenis LIKE ?", "%ako%").Find(&products).Error
	return products, err

}

func (r *repository) FindByName(Name string) (Product, error) {
	var product Product

	err := r.db.Where("name LIKE ?", Name).Find(&product).Error
	return product, err
}

func (r *repository) FindByRating(Rating int) ([]Product, error) {
	var products []Product

	err := r.db.Where("rating = ?", Rating).Find(&products).Error
	return products, err
}

func (r *repository) FindByPriceBefore(PriceBefore int) ([]Product, error) {
	var products []Product

	err := r.db.Where("price_before = ?", PriceBefore).Find(&products).Error
	return products, err
}

func (r *repository) FindByPriceToday(PriceToday int) ([]Product, error) {
	var products []Product

	err := r.db.Where("price_today = ?", PriceToday).Find(&products).Error
	return products, err
}

func (r *repository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) Delete(product Product) (Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}