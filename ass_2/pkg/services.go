package pkg

import (
	"errors"
	"fmt"
	"net/http"
)

func SearchProduct(r http.Request) {
	title := r.FormValue("search")
	fmt.Println(title)
	//ProductDeserializer()
	//var filteredProducts []Product
	//for _, val := range ProductDb {
	//	if strings.Contains(val.Title, title) {
	//		filteredProducts = append(filteredProducts, *val)
	//	}
	//}
	//return filteredProducts
}

func GiveRating(id int, rate float64) error {
	ProductDeserializer()

	list, err := ProductDb[id]

	if !err {
		return errors.New("product doesn't exists")
	}

	list.ArrRating = append(list.ArrRating, rate)

	ProductSerializer()
	return nil
}

func FilterByRating(lowRating float64, highRating float64) {
	ProductDeserializer()

	for key, val := range ProductDb {
		if lowRating < val.Rating() && highRating > val.Rating() {
			fmt.Println(key, val.Title, val.Rating(), val.Price)
		}
	}

}

func FilterByPrice(lowPrice int, highPrice int) {
	ProductDeserializer()

	for key, val := range ProductDb {
		if lowPrice < val.Price && highPrice > val.Price {
			fmt.Println(key, val.Title, val.Rating(), val.Price)
		}
	}
}
