package main

import (
	"fmt"
)

func main() {
	//var a map[int]int
	//
	//a[1] = 1
	increment()

}

func increment() {
	defer fmt.Println("вызвался defer 1")
	defer fmt.Println("вызвался defer 2")
	defer fmt.Println("вызвался defer 3")
	fmt.Println("1")
	fmt.Println("2")
	var a map[int]int

	a[1] = 1

	fmt.Println("3")
}
