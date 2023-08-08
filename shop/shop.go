package shop

import (
	"fmt"
	"github.com/Shemistan/Lesson_3/controllers"
	"github.com/Shemistan/Lesson_3/errors"
	"github.com/Shemistan/Lesson_3/models"
	"github.com/Shemistan/Lesson_3/repositories"
	"github.com/Shemistan/Lesson_3/services"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/*func MultipleInput(arr []float64) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	for _, v := range strings.Split(str, " ") {
		i, _ := strconv.ParseFloat(v, 64)
		arr = append(arr, i)
	}
}*/

func cin[T any](msg string, arg T) {
	fmt.Print(msg)
	fmt.Scan(arg)
}

type Shop struct {
	user models.User
	cart models.Cart

	database models.Database

	usersRepository *repositories.UserRepository
	usersServices   *services.UserService
	usersController *controllers.UserController

	productRepository *repositories.ProductRepository
	productServices   *services.ProductService
	productController *controllers.ProductController
}

func (this *Shop) Init() {
	this.usersRepository = repositories.NewUserRepository(&this.database)
	this.usersServices = services.NewUserService(this.usersRepository)
	this.usersController = controllers.NewUserController(this.usersServices)

	this.productRepository = repositories.NewProductRepository(&this.database)
	this.productServices = services.NewProductService(this.productRepository)
	this.productController = controllers.NewProductController(this.productServices)
}

func (this *Shop) Sell() error {
	ClearScreen()

	var product models.Product

	fmt.Println("Окно добавления товара\n")

	cin("Наименование продукта: ", &product.FullName)
	cin("Описание продукта: ", &product.Description)
	cin("Категория продукта: ", &product.Category)
	cin("Количество продукта: ", &product.Quantity)
	cin("Цена продукта: ", &product.Price)

	return this.productController.AddProduct(product)
}

func (this *Shop) Buy() error {
	ClearScreen()

	products, err := this.productController.GetProducts()
	var productID, quantity float64

	if err != nil {
		return err
	}

	fmt.Println("Окно покупки товара\n")

	for _, product := range products {
		fmt.Printf("[%v]  Наименование продукта: %v\n", product.Id, product.FullName)
		fmt.Printf("     Описание продукта: %v\n", product.Description)
		fmt.Printf("     Категория продукта: %v\n", product.Category)
		fmt.Printf("     Количество продукта: %v\n", product.Quantity)
		fmt.Printf("     Цена продукта: %v\n", product.Price)
	}

	fmt.Println("\nВведите номер продуктa и количество(0 для выхода): ")
	fmt.Scan(&productID, &quantity)
	/*
		if input[0] != 0 && len(input)%2 == 0 {
			for i := 0; i < len(input); i += 2 {
				this.cart.AddProduct(products[int(input[i])], input[i+1])
			}
		}
	*/
	return nil
}

func (this *Shop) Cart() error {
	ClearScreen()

	var command int

	fmt.Println("Корзина\n")

	for _, item := range this.cart.Products {
		fmt.Printf("\n[%v] Наименование продукта: %v\n", item.Id, item.FullName)
		fmt.Printf("     Описание продукта: %v\n", item.Description)
		fmt.Printf("     Категория продукта: %v\n", item.Category)
		fmt.Printf("     Количество продукта: %v\n", item.Quantity)
		fmt.Printf("     Цена продукта: %v\n", item.Price)
	}

	fmt.Println("\nОбщая сумма: ", this.cart.TotalSum)
	fmt.Println("[1]. Оплатить\n[2]. Убрать товар из карзины\n[3]. Выход")

	fmt.Scanf("Выбор: %v\n", &command)

	switch command {
	case 1:
		{
			withdrawErr := this.usersController.Withdraw(this.user, this.cart.TotalSum)
			for _, v := range this.cart.Products {
				updateErr := this.productController.UpdateProduct(v.Id, v)
				if updateErr != nil {
					break
				}
			}

			if withdrawErr != nil {
				return withdrawErr
			}
		}

	case 2:
		fmt.Println("\nУдаление товара из корзины...")
		fmt.Println("Введите номер продукта(0 для выхода): ")

		var productID int
		/*
			if input[0] != 0 {
				for _, id := range input {
					this.cart.RemoveProduct(int(id))
				}
			}
		*/
		this.cart.RemoveProduct(productID)

	case 3:
		return nil

	default:
		return errors.NotExists
	}

	return nil
}

func (this *Shop) Run() error {
	isAuth := false
	var command int

	for {
		//ClearScreen()
		if !isAuth {
			cin("Aвторизация\n\n[1]. Зарегистрироваться\n[2]. Войти\n[3]. Выход\n\nВыбор: ", &command)

			switch command {
			case 1:
				{
					fmt.Println("\nРегистрация")
					newUser := models.User{}
					cin("Имя: ", &newUser.FullName)
					cin("Номер телефона: ", &newUser.PhoneNumber)
					cin("Начальный баланс: ", &newUser.Balance)
					cin("Логин: ", &newUser.Login)
					cin("Пароль: ", &newUser.Password)

					err := this.usersController.AddUser(newUser)
					if err != nil {
						return err
					}

					this.user = newUser

					isAuth = true
				}
			case 2:
				{
					fmt.Println("\nВход")
					var login, password string

					cin("Логин: ", &login)
					cin("Пароль: ", &password)

					enterUser, err := this.usersController.GetUser(login, password)
					if err != nil {
						return err
					}

					this.user = enterUser
				}
			case 3:
				continue
			}
		}

		if isAuth {
			cin("Магазин\n\n[1]. Продать\n[2]. Купить\n[3]. Корзина\n[4]. Выход\n\nВыбор: ", &command)
			switch command {
			case 1:
				{
					err := this.Sell()
					if err != nil {
						fmt.Println(err)
					}
				}
			case 2:
				{
					err := this.Buy()
					if err != nil {
						fmt.Println(err)
					}
				}
			case 3:
				{
					err := this.Cart()
					if err != nil {
						fmt.Println(err)
					}
				}
			case 4:
				continue
			}
		}
	}
}
