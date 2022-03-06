package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/product"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type productHandler struct {
	productService product.Service
}

func NewProductHandler(productService product.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetProductsHandler(c *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var productsResponse []product.ProductResponse

	for _, p := range products {
		productResponse := convertToProductResponse(p)

		productsResponse = append(productsResponse, productResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productsResponse,
	})
}

func (h *productHandler) GetProductByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	p, err := h.productService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	idResponse := convertToProductResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data": idResponse,
	})
}

func (h *productHandler) GetProductByJenis(c *gin.Context) {
	Jenis := c.Param("jenis")
	p, err := h.productService.FindByJenis(Jenis)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	var jenisResponse []product.ProductResponse

	for _, p := range p {
		jenisResponses := convertToProductResponse(p)

		jenisResponse = append(jenisResponse, jenisResponses)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jenisResponse,
	})
}

func (h *productHandler) GetProductByName(c *gin.Context)  {
	name :=  c.Param("name")

	p,err := h.productService.FindByName(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
	}

	nameResponse := convertToProductResponse(p)

	c.JSON(http.StatusOK, gin.H{
		"data" : nameResponse,
	})
}

func (h *productHandler) GetProductByRating(c *gin.Context)  {
	ratingString := c.Param("rating")
	id, _ := strconv.Atoi(ratingString)

	r, err := h.productService.FindByRating(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors" : err,
		})
	}

	var ratingResponses []product.ProductResponse

	for _, r := range r {
		ratingResponse := convertToProductResponse(r)

		ratingResponses = append(ratingResponses, ratingResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ratingResponses,
	})

}

func (h *productHandler) GetProductByPriceBefore(c *gin.Context) {
	priceBeforeString := c.Param("price_before")
	id, _ := strconv.Atoi(priceBeforeString)

	pb, err := h.productService.FindByPriceBefore(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var beforeResponses []product.ProductResponse

	for _, pb := range pb {
		beforeResponse := convertToProductResponse(pb)

		beforeResponses = append(beforeResponses, beforeResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": beforeResponses,
	})
}

func (h *productHandler) GetProductByPriceToday(c *gin.Context) {
	priceTodayString := c.Param("price_today")
	id, _ := strconv.Atoi(priceTodayString)

	pb, err := h.productService.FindByPriceToday(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var todayResponses []product.ProductResponse

	for _, pb := range pb {
		todayResponse := convertToProductResponse(pb)

		todayResponses = append(todayResponses, todayResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todayResponses,
	})
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	product, err := h.productService.Create(productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var productRequest product.ProductRequest

	err := c.ShouldBindJSON(&productRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	product, err := h.productService.Update(id, productRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	product, err := h.productService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToProductResponse(product),
	})
}


func convertToProductResponse(p product.Product) product.ProductResponse {
	return product.ProductResponse{
		ID:          p.ID,
		Jenis:       p.Jenis,
		Name:        p.Name,
		Description: p.Description,
		Rating:      p.Rating,
		PriceBefore: p.PriceBefore,
		PriceToday:  p.PriceToday,
	}
}
