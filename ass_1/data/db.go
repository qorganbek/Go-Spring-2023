package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var UserDb = map[string]AuthUser{}
var ProductDb = map[int]*Product{}

func UserSerializer() {
	byteValue, _ := json.MarshalIndent(UserDb, "", "  ")
	_ = ioutil.WriteFile("users.json", byteValue, 0644)
}

func UserDeserializer() {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &UserDb)
}

func ProductSerializer() {
	byteValue, _ := json.MarshalIndent(ProductDb, "", "  ")
	_ = ioutil.WriteFile("products.json", byteValue, 0644)
}

func ProductDeserializer() {
	jsonFile, err := os.Open("products.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &ProductDb)
}
