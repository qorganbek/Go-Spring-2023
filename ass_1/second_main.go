package main

import (
	"ass1/data"
	"fmt"
)

// boolean LogIn, Give Rating
func showMenu2() {
	fmt.Println("1. Admin Panel")
	fmt.Println("2. Register user")
	fmt.Println("3. Login User")
	fmt.Println("4. Product")
	fmt.Println("5. Exit")
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

func ProductMenu() {
	fmt.Println("1. Product List")
	fmt.Println("2. Search Product by title")
	fmt.Println("3. Filter Product")
	fmt.Println("4. Give rating a Product")
}

func main() {

	p1 := data.Product{ID: 1, Title: "Iphone 14", Price: 800, IsActive: true}
	fmt.Println(p1)
	data.GiveRating(1, 10)
	data.GiveRating(1, 7)
	fmt.Println(p1.ArrRating)
	data.GiveRating(1, 7)
	//ok := true
	//command := 0
	//for ok {
	//	showMenu2()
	//	fmt.Scan(&command)
	//	if command == 5 {
	//		ok = false
	//		break
	//	} else if command == 1 {
	//		AdminPanel()
	//		var adminCommand int
	//		fmt.Scan(&adminCommand)
	//		if adminCommand == 1 {
	//			AdminUser()
	//			var adminUserCommand int
	//			fmt.Scan(&adminUserCommand)
	//			if adminUserCommand == 3 {
	//				data.Service{}.ListUsers()
	//			} else if adminUserCommand == 1 {
	//				var email, password string
	//				fmt.Scan(&email, &password)
	//				u1 := data.User{Email: email, Password: password}
	//				data.Service{}.CreateUser(u1)
	//			} else if adminUserCommand == 2 {
	//				var email, password string
	//				fmt.Scan(&email, &password)
	//				u1 := data.User{Email: email, Password: password}
	//				data.Service{}.DeleteUser(u1)
	//			}
	//		} else if adminCommand == 2 {
	//			AdminProduct()
	//			var adminProductCommand int
	//			fmt.Scan(&adminProductCommand)
	//			if adminProductCommand == 1 {
	//				var id, price int
	//				var title string
	//				var isActive bool
	//				fmt.Scan(&id, &price, &title, &isActive)
	//				p1 := data.Product{ID: id, Title: title, Price: price, IsActive: isActive}
	//				data.Service{}.CreateProduct(p1)
	//			} else if adminProductCommand == 2 {
	//				var id, price int
	//				var title string
	//				var isActive bool
	//				fmt.Scan(&id, &price, &title, &isActive)
	//				p1 := data.Product{ID: id, Title: title, Price: price, IsActive: isActive}
	//				data.Service{}.DeleteProduct(p1)
	//
	//			} else if adminProductCommand == 3 {
	//				data.Service{}.ListProducts()
	//			}
	//
	//		} else {
	//			fmt.Println("Wrong command!")
	//		}
	//	} else if command == 2 {
	//		var email, password string
	//		fmt.Scan(&email, &password)
	//		newUser := data.User{Email: email, Password: password}
	//
	//		err := data.Service{}.CreateUser(newUser)
	//
	//		if err != nil {
	//			fmt.Println("New User authorization Failed!")
	//		} else {
	//			fmt.Println("New User authorization Success!")
	//		}
	//
	//	} else if command == 3 {
	//		var email, password string
	//		fmt.Scan(&email, &password)
	//		User := data.User{Email: email, Password: password}
	//		found := data.Service{}.VerifyUser(User)
	//		if found {
	//			fmt.Println("User Log In Success!")
	//		} else {
	//			fmt.Println("User Log In Failed!")
	//		}
	//	} else if command == 4 {
	//		ProductMenu()
	//		var commandProduct int
	//		fmt.Scan(&commandProduct)
	//		if commandProduct == 1 {
	//			data.Service{}.ListProducts()
	//		} else if commandProduct == 2 {
	//			var title string
	//			fmt.Scan(&title)
	//			data.SearchProduct(title)
	//		} else if commandProduct == 4 {
	//			var rate, id int
	//			fmt.Scan(&rate, &id)
	//
	//		} else {
	//			fmt.Println("Wrong command!")
	//		}
	//	} else {
	//		fmt.Println("Wrong command!")
	//	}
	//
	//}
	//fmt.Println("End. Thanks for using our system!")
}
