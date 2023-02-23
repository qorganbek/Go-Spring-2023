package pkg

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
}

type Service Admin

func (Service) CreateUser(newUser User) error {
	UserDeserializer()
	_, ok := UserDb[newUser.Email]

	if ok {
		return errors.New("user already exists")
	}

	hash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return nil
	}
	newAuthUser := AuthUser{Email: newUser.Email, PasswordHash: hash}
	UserDb[newUser.Email] = newAuthUser

	UserSerializer()

	return nil
}

func (Service) VerifyUser(user User) bool {
	UserDeserializer()
	authUser, ok := UserDb[user.Email]

	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(user.Password))
	return err == nil
}

func (Service) ListUsers() {

	UserDeserializer()

	for key, val := range UserDb {
		fmt.Println(key, val.PasswordHash)
	}
}

func (Service) DeleteUser(currentUser User) error {
	UserDeserializer()

	authUser, ok := UserDb[currentUser.Email]

	if !ok {
		return errors.New("account doesn't exists")
	}

	delete(UserDb, authUser.Email)

	UserSerializer()
	return nil
}

func (Service) ListProducts() {

	ProductDeserializer()

	for key, val := range ProductDb {
		fmt.Println(key, val.Title, val.Rating())
	}
}

func (Service) CreateProduct(newProduct Product) error {
	ProductDeserializer()
	_, ok := ProductDb[newProduct.ID]

	if ok {
		return errors.New("product already exists")
	}

	newAuthProduct := Product{
		ID:    newProduct.ID,
		Title: newProduct.Title,
		Price: newProduct.Price,
	}
	ProductDb[newProduct.ID] = &newAuthProduct

	ProductSerializer()

	return nil
}

func (Service) DeleteProduct(currentProduct Product) error {
	ProductDeserializer()

	curProduct, ok := ProductDb[currentProduct.ID]

	if !ok {
		return errors.New("product doesn't exists")
	}

	delete(ProductDb, curProduct.ID)

	ProductSerializer()
	return nil
}

func (Service) GetLastIdProduct() int {
	ProductDeserializer()
	res := 1
	for _, val := range ProductDb {
		if val.ID > res {
			res = val.ID
		}
	}
	return res
}
