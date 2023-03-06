package pkg

type Product struct {
	ID          int
	Title       string
	Price       int
	Description string
	ImgLink     string
	ArrRating   []float64
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
