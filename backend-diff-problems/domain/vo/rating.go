package vo

import "fmt"

type Rating struct {
	rating int
	color  Color
}

func NewNoRating() Rating {
	return Rating{0, Black}
}

func NewRating(rating int) (Rating, error) {
	if rating < 0 {
		return Rating{}, fmt.Errorf("rating must be greater 0")
	}

	return Rating{rating, ratingToColor(rating)}, nil
}

func ratingToColor(rating int) Color {
	tmp := rating / 400
	if tmp == 0 {
		return Gray
	} else if tmp == 1 {
		return Brown
	} else if tmp == 2 {
		return Green
	} else if tmp == 3 {
		return Cyan
	} else if tmp == 4 {
		return Blue
	} else if tmp == 5 {
		return Yellow
	} else if tmp == 6 {
		return Orange
	} else {
		return Red
	}
}
