package main

import (
	"ass1/data"

	"fmt"
)

func showMenu() {
	fmt.Println("1. View users")
	fmt.Println("2. Register user")
	fmt.Println("3. Login User")
	fmt.Println("4. Product List")
	fmt.Println("5. Exit")
}

func _AdminPanel() {
	fmt.Println("1. User")
	fmt.Println("2. Product")
}

func _AdminUser() {
	fmt.Println("1. Create User")
	fmt.Println("2. Update User")
	fmt.Println("3. Delete User")
	fmt.Println("4. List User")
}

func _AdminProduct() {
	fmt.Println("1. Create Product")
	fmt.Println("2. Update Product")
	fmt.Println("3. Delete Product")
	fmt.Println("4. List Product")
}

//Main Menu
//1. Admin Panel -> User, Product[Create, Delete, Update, List]
//2. Register User
//3. Log In User
//4. Exit

//After Registration
//1.Product List
//2.Log out
//3.Exit

//Product List
//...

func main() {
	ok := true
	var req int
	for ok {
		showMenu()
		fmt.Scan(&req)

		switch req {
		case 1:
			data.Service{}.ListUsers()
		case 2:
			var email, password string
			fmt.Scan(&email, &password)

			newUser := data.User{Email: email, Password: password}

			err := data.Service{}.CreateUser(newUser)
			if err != nil {
				fmt.Println("New User authorization Failed!")
			} else {
				fmt.Println("New User authorization Success!")
			}
		case 3:
			var email, password string
			fmt.Scan(&email, &password)

			user := data.User{Email: email, Password: password}

			found := data.Service{}.VerifyUser(user)
			if found {
				fmt.Println("User Log In Success!")
			} else {
				fmt.Println("User Log In Failed!")
			}
		case 4:
			ok = false
			break
		}
	}

	fmt.Println("End. Thanks for using our system!")
}
