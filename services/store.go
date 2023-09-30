package services

import "fmt"

var cart []string

func AddProduct(product string) {
	cart = append(cart, product)
	fmt.Printf("Продукт %s успешно добавлен в корзину \n", product)
}

func Order() {
	fmt.Println(cart)
	cart = []string{}
}
