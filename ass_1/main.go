package main

import (
	"ass1/data"
	"fmt"
)

func showMainMenu() {
	fmt.Println("1. Admin Panel")
	fmt.Println("2. Register user")
	fmt.Println("3. Login User")
	fmt.Println("4. Exit")
}

func showLogInMenu() {
	fmt.Println("1. Log Out")
	fmt.Println("2. Product")
	fmt.Println("3. Exit")
}

func ProductMenu() {
	fmt.Println("1. Product List")
	fmt.Println("2. Search Product by title")
	fmt.Println("3. Filter Product")
	fmt.Println("4. Give rating a Product")
}

func AdminPanel() {
	fmt.Println("1. User")
	fmt.Println("2. Product")
}

func AdminUser() {
	fmt.Println("1. Create User")
	fmt.Println("2. Delete User")
	fmt.Println("3. List User")
}

func AdminProduct() {
	fmt.Println("1. Create Product")
	fmt.Println("2. Delete Product")
	fmt.Println("3. List Product")
}

func FilterMenu() {
	fmt.Println("1. Filter by Rating")
	fmt.Println("2. Filter by Price")
}

func main() {
	ok := true
	id := 1
	logIn := false
	var command int
	for ok {
		if !logIn {
			showMainMenu()
			fmt.Scan(&command)
			if command == 1 {
				AdminPanel()
				var adminCommand int
				fmt.Scan(&adminCommand)
				if adminCommand == 1 {
					AdminUser()
					var adminUserCommand int
					fmt.Scan(&adminUserCommand)

					if adminUserCommand == 1 {
						fmt.Println("If u want to create user enter your username/email and password")
						var email, password string
						fmt.Scan(&email, &password)
						u1 := data.User{Email: email, Password: password}
						err := data.Service{}.CreateUser(u1)
						if err != nil {
							fmt.Println("New User authorization Failed!")
						} else {
							fmt.Println("New User authorization Success!")
						}

					} else if adminUserCommand == 2 {
						fmt.Println("If u want to delete account enter your username/email and password")
						var email, password string
						fmt.Scan(&email, &password)
						u1 := data.User{Email: email, Password: password}
						err := data.Service{}.DeleteUser(u1)
						if err != nil {
							fmt.Println("User deletion failed!")
						} else {
							fmt.Println("User successfully deleted")
						}
					} else if adminUserCommand == 3 {
						fmt.Println("List of users: ")
						data.Service{}.ListUsers()
					} else {
						fmt.Println("Wrong command!")
					}

				} else if adminCommand == 2 {
					AdminProduct()
					var adminProductCommand int
					fmt.Scan(&adminProductCommand)

					if adminProductCommand == 1 {
						fmt.Println("If u want to create new product title, price")
						var price int
						var title string
						fmt.Scan(&title, &price)
						p1 := data.Product{ID: id, Title: title, Price: price}
						err := data.Service{}.CreateProduct(p1)
						id++
						if err != nil {
							fmt.Println("Create new product failed!")
						} else {
							fmt.Println("New Product successfully created!")
						}
					} else if adminProductCommand == 2 {
						var idD, price int
						var title string
						fmt.Scan(&idD, &title, &price)
						p1 := data.Product{ID: idD, Title: title, Price: price}
						err := data.Service{}.DeleteProduct(p1)
						if err != nil {
							fmt.Println("Product deletion failed!")
						} else {
							fmt.Println("Delete product failed!")
						}
					} else if adminProductCommand == 3 {
						fmt.Println("Product List: ")
						data.Service{}.ListProducts()
					} else {
						fmt.Println("Wrong command!")

					}
				} else {
					fmt.Println("Wrong command!")
				}

			} else if command == 2 {
				fmt.Println("Enter your email/username and password")
				var email, password string
				fmt.Scan(&email, &password)
				newUser := data.User{Email: email, Password: password}

				err := data.Service{}.CreateUser(newUser)

				if err != nil {
					fmt.Println("New User authorization Failed!")
				} else {
					fmt.Println("New User authorization Success!")
				}
			} else if command == 3 {
				fmt.Println("Enter your email/username and password")
				var email, password string
				fmt.Scan(&email, &password)
				User := data.User{Email: email, Password: password}
				found := data.Service{}.VerifyUser(User)
				if found {
					logIn = true
					fmt.Println("User Log In Success!")
				} else {
					fmt.Println("User Log In Failed!")
				}
			} else if command == 4 {
				ok = false
				break
			} else {
				fmt.Println("Wrong command!")
			}

		} else {
			showLogInMenu()
			fmt.Scan(&command)
			if command == 1 {
				fmt.Println("Thank u for using our system")
				logIn = false
			} else if command == 2 {
				ProductMenu()
				var productCommand int
				fmt.Scan(&productCommand)
				if productCommand == 1 {
					fmt.Println("Product List: ")
					data.Service{}.ListProducts()
				} else if productCommand == 2 {
					fmt.Println("Please enter title if product: ")
					var title string
					fmt.Scan(&title)
					data.SearchProduct(title)
				} else if productCommand == 4 {
					fmt.Println("If u want to give rating to product enter id product and your rating 1-10")
					var rate float64
					var idG int
					fmt.Scan(&idG, &rate)
					data.GiveRating(idG, rate)
				} else if productCommand == 3 {
					FilterMenu()
					var filterCommand int
					fmt.Scan(&filterCommand)

					if filterCommand == 1 {
						fmt.Println("Enter your low rating and high rating")
						var highRating, lowRating float64
						fmt.Scan(&lowRating, &highRating)
						fmt.Println("Products by filter: ")
						data.FilterByRating(lowRating, highRating)
					} else if filterCommand == 2 {
						fmt.Println("Enter your low price and high price")
						var highPrice, lowPrice int
						fmt.Scan(&lowPrice, &highPrice)
						fmt.Println("Products by filter: ")
						data.FilterByPrice(lowPrice, highPrice)
					} else {
						fmt.Println("Wrong command!")
					}

				} else {
					fmt.Println("Wrong command!")
				}
			} else if command == 3 {
				fmt.Println("Thanks for using our system!")
				ok = false
				break
			} else {
				fmt.Println("Wrong command!")
			}
		}
	}
	fmt.Println("End. Thanks for using our system!")
}
