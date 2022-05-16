package converter

import (
	"diff-problems/domain/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConvertContestResultList_正常系(t *testing.T) {
	bytes := []byte(`[{"IsRated":true,"Place":59,"OldRating":0,"NewRating":1255,"Performance":2455,"InnerPerformance":2455,"ContestScreenName":"arc061.contest.atcoder.jp","ContestName":"AtCoder Regular Contest 061","ContestNameEn":"","EndTime":"2016-09-11T22:40:00+09:00"},{"IsRated":true,"Place":9,"OldRating":1255,"NewRating":2155,"Performance":3192,"InnerPerformance":3192,"ContestScreenName":"arc064.contest.atcoder.jp","ContestName":"AtCoder Regular Contest 064","ContestNameEn":"","EndTime":"2016-12-04T22:40:00+09:00"},{"IsRated":true,"Place":8,"OldRating":2155,"NewRating":2676,"Performance":3623,"InnerPerformance":3623,"ContestScreenName":"agc010.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 010","ContestNameEn":"","EndTime":"2017-02-04T22:50:00+09:00"},{"IsRated":true,"Place":41,"OldRating":2676,"NewRating":2728,"Performance":2974,"InnerPerformance":2974,"ContestScreenName":"agc014.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 014","ContestNameEn":"","EndTime":"2017-05-06T23:10:00+09:00"},{"IsRated":true,"Place":45,"OldRating":2728,"NewRating":2733,"Performance":2813,"InnerPerformance":2813,"ContestScreenName":"agc015.contest.atcoder.jp","ContestName":"AtCoder Grand Contest 015","ContestNameEn":"","EndTime":"2017-05-27T22:50:00+09:00"}]`)
	actual, err := ConvertContestResultList(bytes)
	assert.Nil(t, err)

	expected := vo.ContestResultList{
		{true, 1255, 1473601200},
		{true, 2155, 1480858800},
		{true, 2676, 1486216200},
		{true, 2728, 1494079800},
		{true, 2733, 1495893000},
	}

	assert.Equal(t, expected, actual)
}
