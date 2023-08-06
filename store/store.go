package main

import "fmt"

const (
	reg  = 1
	auth = 2
	exit = 3
)

func order(products []string) {
	fmt.Println("Now you have ordered eveything. Congrats!\nReturning to main menu...")
	return
}
func shop() {
	fmt.Println("Now you are in the shop.\nPlease type the name of the item you want to buy,\nand it will automatically be added to your cart.")
	fmt.Println("After you type all the items, please enter 0, \nso that this will indicate that you are done.")
	var item string
	productList := []string{}
	for item != "0" && len(productList) < 10 {
		fmt.Scanln(&item)
		if item == "0" {
			break
		}
		productList = append(productList, item)
	}
	fmt.Println("items in your cart are: ")
	for _, i := range productList {
		fmt.Println(i)
	}
	fmt.Println("If you want to buy all these items please type 'order'")
	fmt.Scanln(&item)
	for {
		if item == "order" {
			order(productList)
			break
		} else {
			fmt.Println("Please enter 'order'")
		}
	}
}

func register(usersPasswords []string) {
	fmt.Println("Please enter your name and password like this: name_password")
	var name string
	fmt.Scan(&name)

	for i := range usersPasswords {
		if name == usersPasswords[i] {
			fmt.Println("You are already registered")
			return
		} else {
			usersPasswords = append(usersPasswords, name)
			shop()
		}
	}
}

func authenticate(usersPasswords []string) {
	fmt.Println("Please enter your name and password like this: name_password")
	var name string
	fmt.Scanln(&name)
	found := false
	for _, v := range usersPasswords {
		if v == name {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Welcome!")
		shop()
	} else {
		fmt.Println("You are not registered")
	}

}

func main() {
	fmt.Println("Welcome to the store!")
	var choice int
	usersPasswords := []string{"name_password"}
	for choice != exit {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Register")
		fmt.Println("2. Authenticate")
		fmt.Println("3. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case reg:
			fmt.Println("Registering...")
			register(usersPasswords)
		case auth:
			fmt.Println("Authenticating...")
			authenticate(usersPasswords)
		case exit:
			fmt.Println("Exiting...")
			break
		}
	}
}

// func SaveInt(req string) bool {
// 	// code

// 	return true
// }

// func daveText(req string) bool {
// 	// code

// 	return true
// }
