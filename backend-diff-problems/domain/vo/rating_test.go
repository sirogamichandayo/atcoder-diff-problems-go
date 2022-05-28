package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_正常系_NewNoRating(t *testing.T) {
	rating := NewNoRating()
	assert.Nil(t, rating.Rating())
	assert.Equal(t, rating.Color(), BlackRating)
}

func Test_正常系_NewRating(t *testing.T) {
	dataList := []struct {
		name          string
		argRating     int
		expectedColor RatingColor
	}{
		{name: "灰色下限", argRating: 0, expectedColor: GrayRating},
		{name: "灰色上限", argRating: 399, expectedColor: GrayRating},
		{name: "茶色下限", argRating: 400, expectedColor: BrownRating},
		{name: "茶色上限", argRating: 799, expectedColor: BrownRating},
		{name: "緑色下限", argRating: 800, expectedColor: GreenRating},
		{name: "緑色上限", argRating: 1199, expectedColor: GreenRating},
		{name: "水色下限", argRating: 1200, expectedColor: CyanRating},
		{name: "水色上限", argRating: 1599, expectedColor: CyanRating},
		{name: "青色下限", argRating: 1600, expectedColor: BlueRating},
		{name: "青色上限", argRating: 1999, expectedColor: BlueRating},
		{name: "黄色下限", argRating: 2000, expectedColor: YellowRating},
		{name: "黄色上限", argRating: 2399, expectedColor: YellowRating},
		{name: "橙色下限", argRating: 2400, expectedColor: OrangeRating},
		{name: "橙色上限", argRating: 2799, expectedColor: OrangeRating},
		{name: "赤色下限", argRating: 2800, expectedColor: RedRating},
		{name: "赤色上限", argRating: 3199, expectedColor: RedRating},
		{name: "銀色下限", argRating: 3200, expectedColor: SilverRating},
		{name: "銀色上限", argRating: 3599, expectedColor: SilverRating},
		{name: "金色下限", argRating: 3600, expectedColor: GoldRating},
		{name: "金色2", argRating: 3799, expectedColor: GoldRating},
		{name: "金色(tourist最高値)", argRating: 4229, expectedColor: GoldRating},
		{name: "金色(システム最高値)", argRating: 9999, expectedColor: GoldRating},
	}

	for _, d := range dataList {
		t.Run(d.name, func(t *testing.T) {
			rating, err := NewRating(d.argRating)
			assert.Nil(t, err)
			assert.Equal(t, d.argRating, *rating.Rating())
			assert.Equal(t, d.expectedColor, rating.Color())
		})
	}
}
