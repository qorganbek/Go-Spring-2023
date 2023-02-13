package data

import (
	"errors"
	"fmt"
)

type Product struct {
	ID        int
	Title     string
	Price     int
	ArrRating []int
}

func (p Product) Rating() float32 {
	if len(p.ArrRating) == 0 {
		return 0.0
	}
	sum := 0
	for i := range p.ArrRating {
		sum += p.ArrRating[i]
	}
	return float32(sum / len(p.ArrRating))
}

func SearchProduct(title string) {
	ProductDeserializer()
	for key, val := range ProductDb {
		if val.Title == title {
			fmt.Println(key, val.Price, val.Title, val.Rating())
		}
	}
}

func GiveRating(id int, rate int) error {
	ProductDeserializer()

	list, err := ProductDb[id]

	if !err {
		return errors.New("product doesn't exists")
	}

	list.ArrRating = append(list.ArrRating, rate)

	ProductSerializer()
	return nil
}
