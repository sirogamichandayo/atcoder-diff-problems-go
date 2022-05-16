package converter

import (
	"diff-problems/domain/vo"
	"diff-problems/util_tool"
	"encoding/json"
	"fmt"
)

func ConvertContestResultList(jsonBytes []byte) (vo.ContestResultList, error) {
	var rawList []map[string]interface{}
	err := json.Unmarshal(jsonBytes, &rawList)
	if err != nil {
		return nil, err
	}

	list := make(vo.ContestResultList, 0, len(rawList))
	for _, raw := range rawList {
		isRated, hasKey := raw["IsRated"].(bool)
		if !hasKey {
			return nil, fmt.Errorf("not exist isRated")
		}
		rawNewRating, hasKey := raw["NewRating"]
		if !hasKey {
			return nil, fmt.Errorf("not exist newRating")
		}
		newRating := int(rawNewRating.(float64))
		rawEndTime, hasKey := raw["EndTime"].(string)
		if !hasKey {
			return nil, fmt.Errorf("not exist endTime")
		}
		endTime, err := util_tool.ConvertIso8601JstToEpochTime(rawEndTime)
		if err != nil {
			return nil, err
		}

		list = append(list, vo.ContestResult{IsRated: isRated, NewRating: newRating, EndTime: endTime})
	}

	return list, nil
}
