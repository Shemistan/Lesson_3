package models

type Cart struct {
	Products []Product
	TotalSum float64
}

func (c *Cart) AddProduct(product Product, quantity float64) {
	if quantity > 0 {
		product.Quantity = quantity
		c.Products = append(c.Products, product)
		c.TotalSum += product.Price * quantity
	}
}

func (c *Cart) RemoveProduct(id int) {
	for i, product := range c.Products {
		if product.Id == id {
			c.Products = append(c.Products[:i], c.Products[i+1:]...)
			c.TotalSum -= product.Price * product.Quantity
		}
	}
}
