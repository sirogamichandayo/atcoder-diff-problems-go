package vo

import "fmt"

type ContestResult struct {
	IsRated   bool
	NewRating int
	EndTime   int64
}

type ContestResultList []ContestResult

func (list ContestResultList) Last() (ContestResult, error) {
	if len(list) == 0 {
		return ContestResult{}, fmt.Errorf("ContestResultList is empty")
	}

	var res ContestResult
	LastEndTime := int64(0)
	for _, entity := range list {
		if LastEndTime < entity.EndTime {
			LastEndTime = entity.EndTime
			res = entity
		}
	}
	return res, nil
}
