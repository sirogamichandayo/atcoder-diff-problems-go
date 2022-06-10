package _interface

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/vo"
	"time"
)

type UserProgressMaker struct{}

func (_ UserProgressMaker) Make(dList entity.ProblemDifficultyList, sList entity.AcUserSubmissionList) (entity.UserProgressList, error) {
	clipDifficultyMap := make(map[string]vo.ClipDifficulty, len(dList))
	for _, d := range dList {
		clipDifficultyMap[d.ProblemId] = d.ClipDifficulty
	}

	resList := entity.UserProgressList{}
	for _, s := range sList {
		if err := resList.Add(
			time.Unix(s.EpochTime, 0),
			clipDifficultyMap[s.ProblemId],
		); err != nil {
			return resList, err
		}
	}

	return resList, nil
}
