package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_正常系_NewNoRating(t *testing.T) {
	rating := NewNoRating()
	assert.Nil(t, rating.Rating())
	assert.Equal(t, rating.Color(), Black)
}

func Test_正常系_NewRating(t *testing.T) {
	dataList := []struct {
		name          string
		argRating     int
		expectedColor Color
	}{
		{name: "灰色下限", argRating: 0, expectedColor: Gray},
		{name: "灰色上限", argRating: 399, expectedColor: Gray},
		{name: "茶色下限", argRating: 400, expectedColor: Brown},
		{name: "茶色上限", argRating: 799, expectedColor: Brown},
		{name: "緑色下限", argRating: 800, expectedColor: Green},
		{name: "緑色上限", argRating: 1199, expectedColor: Green},
		{name: "水色下限", argRating: 1200, expectedColor: Cyan},
		{name: "水色上限", argRating: 1599, expectedColor: Cyan},
		{name: "青色下限", argRating: 1600, expectedColor: Blue},
		{name: "青色上限", argRating: 1999, expectedColor: Blue},
		{name: "黄色下限", argRating: 2000, expectedColor: Yellow},
		{name: "黄色上限", argRating: 2399, expectedColor: Yellow},
		{name: "橙色下限", argRating: 2400, expectedColor: Orange},
		{name: "橙色上限", argRating: 2799, expectedColor: Orange},
		{name: "赤色下限", argRating: 2800, expectedColor: Red},
		{name: "赤色上限", argRating: 3199, expectedColor: Red},
		{name: "銀色下限", argRating: 3200, expectedColor: Silver},
		{name: "銀色上限", argRating: 3599, expectedColor: Silver},
		{name: "金色下限", argRating: 3600, expectedColor: Gold},
		{name: "金色2", argRating: 3799, expectedColor: Gold},
		{name: "金色(tourist最高値)", argRating: 4229, expectedColor: Gold},
		{name: "金色(システム最高値)", argRating: 9999, expectedColor: Gold},
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
