package _interface

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/vo"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UserProgressMaker_Make_正常系(t *testing.T) {
	problem1Id := "problem1"
	problem2Id := "problem2"
	problem3Id := "problem3"
	problem1Diff, er := vo.NewRawDifficulty(400.0, true)
	assert.Nil(t, er)
	problem2Diff, er := vo.NewRawDifficulty(800.0, true)
	assert.Nil(t, er)
	problem3Diff, er := vo.NewRawDifficulty(1600.0, true)
	assert.Nil(t, er)

	dList := entity.ProblemDifficultyList{
		entity.ProblemDifficulty{
			ProblemId:      problem1Id,
			Difficulty:     problem1Diff,
			ClipDifficulty: problem1Diff.MakeClipDifficulty(),
		},
		entity.ProblemDifficulty{
			ProblemId:      problem2Id,
			Difficulty:     problem2Diff,
			ClipDifficulty: problem2Diff.MakeClipDifficulty(),
		},
		entity.ProblemDifficulty{
			ProblemId:      problem3Id,
			Difficulty:     problem3Diff,
			ClipDifficulty: problem3Diff.MakeClipDifficulty(),
		},
	}

	userId := "userid"
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		assert.Nil(t, err)
	}
	day1 := time.Date(2000, 1, 23, 0, 0, 0, 0, jst)
	day2 := time.Date(2000, 1, 24, 0, 0, 0, 0, jst)
	sList := entity.AcUserSubmissionList{
		entity.AcUserSubmission{
			UserId:    userId,
			ProblemId: problem1Id,
			EpochTime: day1.Unix(),
		},
		entity.AcUserSubmission{
			UserId:    userId,
			ProblemId: problem2Id,
			EpochTime: day2.Unix(),
		},
		entity.AcUserSubmission{
			UserId:    userId,
			ProblemId: problem3Id,
			EpochTime: day2.Unix(),
		},
	}

	maker := UserProgressMaker{}
	pList, err := maker.Make(dList, sList)
	assert.Nil(t, err)

	fmt.Println(day1, day2)
	p := pList.Progress()
	assert.Equal(t, 2, len(p))
	fmt.Println(p)
	fmt.Println(day1.Unix())
	for a, _ := range p {
		if day1.Unix() == a.Unix() {
			fmt.Println("huga")
		}
		if day1 == a {
			fmt.Println("hoge")
		}
	}
}
