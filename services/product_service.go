package services

import (
	"github.com/Shemistan/Lesson_3/models"
	"github.com/Shemistan/Lesson_3/repositories"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{ProductRepository: productRepository}
}

func (ps ProductService) GetProducts() ([]models.Product, error) {
	return ps.ProductRepository.GetProducts()
}

func (ps ProductService) AddProduct(product models.Product) error {
	return ps.ProductRepository.AddProduct(product)
}

func (ps ProductService) UpdateProduct(id int, product models.Product) error {
	return ps.ProductRepository.UpdateProduct(id, product)
}
