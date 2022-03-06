package product

import (
	"fmt"
)

type Service interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	FindByJenis(Jenis string) ([]Product, error)
	FindByName(Name string) (Product, error)
	FindByRating(Rating int) ([]Product, error)
	FindByPriceBefore(PriceBefore int) ([]Product, error)
	FindByPriceToday(PriceToday int) ([]Product, error)
	Create(product ProductRequest) (Product, error)
	Update(ID int, product ProductRequest) (Product, error)
	Delete(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Product, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Product, error) {
	return s.repository.FindByID(ID)
}

func (s *service) FindByJenis(Jenis string) ([]Product, error) {
	return s.repository.FindByJenis(Jenis)
}

func (s *service) FindByName(Name string) (Product, error) {
	return s.repository.FindByName(Name)
}

func (s *service) FindByRating(Rating int) ([]Product, error) {
	return s.repository.FindByRating(Rating)
}

func (s *service) FindByPriceBefore(PriceBefore int) ([]Product, error) {
	return s.repository.FindByPriceBefore(PriceBefore)
}

func (s *service) FindByPriceToday(PriceToday int) ([]Product, error) {
	return s.repository.FindByPriceToday(PriceToday)
}

func (s *service) Create(productRequest ProductRequest) (Product, error) {

	rating, _ := productRequest.Rating.Int64()
	priceBefore, _ := productRequest.PriceBefore.Int64()
	priceToday, _ := productRequest.PriceToday.Int64()

	product := Product{
		Jenis:       productRequest.Jenis,
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Rating:      int(rating),
		PriceBefore: int(priceBefore),
		PriceToday:  int(priceToday),
	}

	return s.repository.Create(product)
}

func (s *service) Update(ID int, productRequest ProductRequest) (Product, error) {

	product, err := s.repository.FindByID(ID)
	if err != nil {
		fmt.Println("cannot update data")
	}

	rating, _ := productRequest.Rating.Int64()
	priceBefore, _ := productRequest.PriceBefore.Int64()
	priceToday, _ := productRequest.PriceToday.Int64()

	product.Jenis = productRequest.Jenis
	product.Name = productRequest.Name
	product.Description = productRequest.Description
	product.Rating = int(rating)
	product.PriceBefore = int(priceBefore)
	product.PriceToday = int(priceToday)

	return s.repository.Update(product)
}


func (s *service) Delete(ID int) (Product, error) {

	product, err := s.repository.FindByID(ID)
	if err != nil {
		fmt.Println("cannot update data")
	}

	return s.repository.Delete(product)
}