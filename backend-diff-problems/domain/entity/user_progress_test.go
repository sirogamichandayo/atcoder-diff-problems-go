package entity

import (
	"diff-problems/domain/vo"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UserProgressList_Add_正常系(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		assert.Nil(t, err)
	}

	day1 := time.Date(2000, 1, 23, 0, 0, 0, 0, jst)
	day11 := time.Date(2000, 1, 23, 0, 0, 0, 0, jst)
	day11ClipDifficulty := vo.ClipDifficulty{
		Difficulty: 100.0,
		Valid:      true,
	}
	day12 := time.Date(2000, 1, 23, 1, 0, 0, 0, jst)
	day12ClipDifficulty := vo.ClipDifficulty{
		Difficulty: 300.0,
		Valid:      true,
	}

	day2 := time.Date(2000, 1, 24, 0, 0, 0, 0, jst)
	day21 := time.Date(2000, 1, 24, 1, 0, 0, 0, jst)
	day21ClipDifficulty := vo.ClipDifficulty{
		Difficulty: 200.0,
		Valid:      true,
	}

	list := UserProgressList{}
	assert.Nil(t, list.Add(day11, day11ClipDifficulty))
	assert.Nil(t, list.Add(day12, day12ClipDifficulty))
	assert.Nil(t, list.Add(day21, day21ClipDifficulty))

	p := list.Progress()
	assert.Equal(t, 2, len(p))
	fmt.Println("hoge")
	fmt.Println(&day1)
	for a, _ := range p {
		fmt.Println(&a)
	}
	assert.Equal(t, float64(400), p[day1].Difficulty)
	assert.Equal(t, float64(200), p[day2].Difficulty)

}
