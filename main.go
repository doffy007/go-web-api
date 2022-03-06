package main

import (
	"fmt"
	"pustaka-api/handler"
	"pustaka-api/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("cannot connect to database")
	}

	db.AutoMigrate(&product.Product{})

	productRepository := product.NewRepository(db)
	productService := product.NewService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/products", productHandler.GetProductsHandler)
	v1.GET("/products/:id", productHandler.GetProductByID)
	v1.GET("/jenisProducts/:jenis", productHandler.GetProductByJenis)
	v1.GET("/nameProduct/:name", productHandler.GetProductByName)
	v1.GET("/ratingProducts/:rating", productHandler.GetProductByRating)
	v1.GET("/PriceBefore/:price_before", productHandler.GetProductByPriceBefore)
	v1.GET("/PriceToday/:price_today", productHandler.GetProductByPriceToday)
	v1.POST("/products", productHandler.CreateProduct)
	v1.PUT("/products/:id", productHandler.UpdateProduct)
	v1.DELETE("/products/:id", productHandler.DeleteProduct)

	router.Run(":8888") //on localhost:8080

	// main
	// repository 
	// service 
	//request
	//response
	//database

}
