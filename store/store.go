package store

import "fmt"

var storage []string

func AddProduct(product string) {
	storage = append(storage, product)
	fmt.Printf("Продукт %s был успешно добавлен в корзину \n", product)
}

func Order() []string {
	result := storage
	storage = []string{}
	fmt.Println("Товары успешно куплены: ")
	return result
}
