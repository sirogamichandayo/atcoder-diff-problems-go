package vo

import "fmt"

type Rating struct {
	rating *int
	color  RatingColor
}

func NewNoRating() Rating {
	return Rating{color: BlackRating}
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

func (r Rating) Color() RatingColor {
	return r.color
}

func ratingToColor(rating int) (RatingColor, error) {
	if rating < 0 {
		return BlackRating, fmt.Errorf("rating must be greater than or equal to zero")
	}
	tmp := rating / 400
	if tmp == 0 {
		return GrayRating, nil
	} else if tmp == 1 {
		return BrownRating, nil
	} else if tmp == 2 {
		return GreenRating, nil
	} else if tmp == 3 {
		return CyanRating, nil
	} else if tmp == 4 {
		return BlueRating, nil
	} else if tmp == 5 {
		return YellowRating, nil
	} else if tmp == 6 {
		return OrangeRating, nil
	} else if tmp == 7 {
		return RedRating, nil
	} else if tmp == 8 {
		return SilverRating, nil
	} else {
		return GoldRating, nil
	}
}
