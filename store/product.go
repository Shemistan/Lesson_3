package store

import "fmt"

func AddProduct(product string, productList *[]string) bool {
	check := checkProduct(product, *productList)
	if check {
		fmt.Println("Такой продукт уже есть")
		return false
	}
	*productList = append(*productList, product)
	message := fmt.Sprintf("Продукт %s успешно добавлен", product)
	fmt.Println(message)
	fmt.Println(productList)
	return true
}

func Order(productList *[]string) bool {
	if len(*productList) == 0 {
		fmt.Println("Корзина пуста")
		return false
	}
	fmt.Printf("Вы купили %v\n", productList)
	*productList = []string{}
	return true
}

func checkProduct(product string, productList []string) bool {
	for _, v := range productList {
		if v == product {
			return true
		}
	}
	return false
}
