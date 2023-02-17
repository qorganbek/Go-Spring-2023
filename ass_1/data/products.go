package data

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	ID        int
	Title     string
	Price     int
	ArrRating []float64
}

func (p Product) Rating() float64 {
	if len(p.ArrRating) == 0 {
		return 0.0
	}
	sum := 0.0
	for i := range p.ArrRating {
		sum += p.ArrRating[i]
	}
	return sum / float64(len(p.ArrRating))
}

func SearchProduct(title string) {
	ProductDeserializer()
	for key, val := range ProductDb {
		if strings.Contains(val.Title, title) {
			fmt.Println(key, val.Price, val.Title, val.Rating())
		}
	}
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
