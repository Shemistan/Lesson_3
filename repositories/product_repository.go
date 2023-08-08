package repositories

import (
	"github.com/Shemistan/Lesson_3/errors"
	"github.com/Shemistan/Lesson_3/models"
)

type ProductRepository struct {
	dbHandler *models.Database
}

func NewProductRepository(dbHandler *models.Database) *ProductRepository {
	return &ProductRepository{dbHandler: dbHandler}
}

func (pr ProductRepository) GetProducts() ([]models.Product, error) {
	if len(pr.dbHandler.Products) == 0 {
		return []models.Product{}, errors.CollectionEmpty
	}
	return pr.dbHandler.Products, nil
}

func (pr ProductRepository) AddProduct(product models.Product) error {
	pr.dbHandler.Products = append(pr.dbHandler.Products, product)
	return nil
}

func (pr ProductRepository) UpdateProduct(id int, product models.Product) error {
	if id < len(pr.dbHandler.Products)-1 && id >= 0 {
		pr.dbHandler.Products[id] = product
		return nil
	}
	return errors.NotExists
}
