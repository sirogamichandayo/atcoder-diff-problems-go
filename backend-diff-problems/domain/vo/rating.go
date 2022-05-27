package vo

import "fmt"

type Rating struct {
	rating *int
	color  Color
}

func NewNoRating() Rating {
	return Rating{color: Black}
}

func NewRating(rating int) (Rating, error) {
	if rating < 0 {
		return Rating{}, fmt.Errorf("rating must be greater than or equal to zero")
	}

	color, err := ratingToColor(rating)
	if err != nil {
		return Rating{}, err
	}

	return Rating{&rating, color}, nil
}

func (r Rating) Rating() *int {
	return r.rating
}

func (r Rating) Color() Color {
	return r.color
}

func ratingToColor(rating int) (Color, error) {
	if rating < 0 {
		return Black, fmt.Errorf("rating must be greater than or equal to zero")
	}
	tmp := rating / 400
	if tmp == 0 {
		return Gray, nil
	} else if tmp == 1 {
		return Brown, nil
	} else if tmp == 2 {
		return Green, nil
	} else if tmp == 3 {
		return Cyan, nil
	} else if tmp == 4 {
		return Blue, nil
	} else if tmp == 5 {
		return Yellow, nil
	} else if tmp == 6 {
		return Orange, nil
	} else if tmp == 7 {
		return Red, nil
	} else if tmp == 8 {
		return Silver, nil
	} else {
		return Gold, nil
	}
}
