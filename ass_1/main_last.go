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
	fmt.Println("1. Log Out User")
	fmt.Println("2. Product")
	fmt.Println("3. Exit")
}

func ProductMenu2() {
	fmt.Println("1. Product List")            // +
	fmt.Println("2. Search Product by title") // +
	fmt.Println("3. Filter Product")          // -
	fmt.Println("4. Give rating a Product")   // +-
}

func AdminPanel2() {
	fmt.Println("1. User")
	fmt.Println("2. Product")
}

func AdminUser2() {
	fmt.Println("1. Create User")
	fmt.Println("2. Delete User")
	fmt.Println("3. List User")
}

func AdminProduct2() {
	fmt.Println("1. Create Product")
	fmt.Println("2. Delete Product")
	fmt.Println("3. List Product")
}

func main() {
	ok := true
	logIn := false
	var command int
	for ok {
		if !logIn {
			showMainMenu()
			fmt.Scan(&command)
			if command == 1 {
				AdminPanel2()
				var adminCommand int
				fmt.Scan(&adminCommand)
				if adminCommand == 1 {
					AdminUser2()
					var adminUserCommand int
					fmt.Scan(&adminUserCommand)

					if adminUserCommand == 1 {
						var email, password string
						fmt.Scan(&email, &password)
						u1 := data.User{Email: email, Password: password}
						data.Service{}.CreateUser(u1)
					} else if adminUserCommand == 2 {
						var email, password string
						fmt.Scan(&email, &password)
						u1 := data.User{Email: email, Password: password}
						data.Service{}.DeleteUser(u1)
					} else if adminUserCommand == 3 {
						data.Service{}.ListUsers()
					} else {
						fmt.Println("Wrong command!")
					}

				} else if adminCommand == 2 {
					AdminProduct2()
					var adminProductCommand int
					fmt.Scan(&adminProductCommand)

					if adminProductCommand == 1 {
						var id, price int
						var title string
						var isActive bool
						fmt.Scan(&id, &price, &title, &isActive)
						p1 := data.Product{ID: id, Title: title, Price: price, IsActive: isActive}
						data.Service{}.CreateProduct(p1)

					} else if adminProductCommand == 2 {
						var id, price int
						var title string
						var isActive bool
						fmt.Scan(&id, &price, &title, &isActive)
						p1 := data.Product{ID: id, Title: title, Price: price, IsActive: isActive}
						data.Service{}.DeleteProduct(p1)

					} else if adminProductCommand == 3 {
						data.Service{}.ListProducts()
					} else {
						fmt.Println("Wrong command!")

					}
				}

			} else if command == 2 {
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
				logIn = false
			} else if command == 2 {
				ProductMenu2()
				var productCommand int
				fmt.Scan(&productCommand)
				if productCommand == 1 {
					data.Service{}.ListProducts()
				} else if productCommand == 2 {
					var title string
					fmt.Scan(&title)
					data.SearchProduct(title)
				} else if productCommand == 4 {
					var rate, id int
					fmt.Scan(&id, &rate)
					data.GiveRating(id, rate)
				} else {
					fmt.Println("Wrong command!")
				}
			} else if command == 3 {
				ok = false
				break
			} else {
				fmt.Println("Wrong command!")
			}
		}
	}

}
