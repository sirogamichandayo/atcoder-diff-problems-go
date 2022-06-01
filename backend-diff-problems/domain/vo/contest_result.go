package vo

import (
	"fmt"
	"math"
)

type ContestResult struct {
	IsRated bool
	Rating  Rating
	EndTime int64
}

type ContestResultList []ContestResult

func (list ContestResultList) Last() (ContestResult, error) {
	if len(list) == 0 {
		return ContestResult{}, fmt.Errorf("ContestResultList is empty")
	}

	var res ContestResult
	LastEndTime := int64(math.MinInt64)
	for _, entity := range list {
		if LastEndTime < entity.EndTime {
			LastEndTime = entity.EndTime
			res = entity
		}
	}
	return res, nil
}

func (list ContestResultList) ExactByRated() ContestResultList {
	res := make(ContestResultList, 0)

	for _, entity := range list {
		if entity.IsRated {
			res = append(res, entity)
		}
	}
	return res
}

func (list ContestResultList) Empty() bool {
	return len(list) == 0
}
