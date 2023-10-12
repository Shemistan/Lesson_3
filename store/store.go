package store

import "fmt"

func AddProduct(productList []string, product string) []string {
	productList = append(productList, product)
	fmt.Printf("Продукт %s успешно добавлен \n", product)
	return productList
}

func Order(productList []string) []string {
	fmt.Println("Вы купили", productList)
	productList = make([]string, 0, 10)
	return productList
}
