package vo

import (
	"database/sql/driver"
	"fmt"
	"math"
)

type RawDifficulty struct {
	Difficulty float64
	Valid      bool
}

const DifficultyThreshold = float64(400)

func NewRawDifficulty(value any, valid bool) (RawDifficulty, error) {
	if !valid {
		return RawDifficulty{Valid: false}, nil
	}

	difficulty, ok := value.(float64)
	if !ok {
		return RawDifficulty{}, fmt.Errorf("%f can not convert to float64", value)
	}

	return RawDifficulty{difficulty, true}, nil
}

func (rd RawDifficulty) MakeClipDifficulty() ClipDifficulty {
	if !rd.Valid {
		return ClipDifficulty{}
	}

	difficulty := rd.Difficulty
	if difficulty < DifficultyThreshold {
		difficulty = 400 / math.Exp(1.0-rd.Difficulty/400)
	}

	return ClipDifficulty{difficulty, true}
}

func (rd RawDifficulty) Value() (driver.Value, error) {
	if !rd.Valid {
		return nil, nil
	}
	return rd.Difficulty, nil
}
