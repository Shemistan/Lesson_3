package controllers

import (
	"github.com/Shemistan/Lesson_3/errors"
	"github.com/Shemistan/Lesson_3/models"
	"github.com/Shemistan/Lesson_3/services"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc ProductController) GetProducts() ([]models.Product, error) {
	return pc.productService.GetProducts()
}

func (pc ProductController) AddProduct(product models.Product) error {
	if product.FullName == "" || product.Category == "" {
		return errors.MissingParams
	}

	if product.Quantity == 0 {
		product.Quantity = 1
	}

	return pc.productService.AddProduct(product)
}

func (pc ProductController) UpdateProduct(id int, product models.Product) error {
	if id < 0 {
		return errors.InvalidArg
	}

	return pc.productService.UpdateProduct(id, product)
}
