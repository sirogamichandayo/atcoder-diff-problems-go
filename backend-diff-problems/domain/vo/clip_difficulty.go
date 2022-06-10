package vo

import (
	"database/sql/driver"
	"fmt"
)

type ClipDifficulty struct {
	Difficulty float64
	Valid      bool
}

type ClipDifficultySum struct {
	Difficulty float64
}

func (s *ClipDifficultySum) Add(difficulty ClipDifficulty) error {
	if s == nil {
		return fmt.Errorf("nil")
	}
	if difficulty.Valid {
		s.Difficulty += difficulty.Difficulty
	}
	return nil
}

func ReconstructClipDifficulty(d *float64) ClipDifficulty {
	if d == nil {
		return ClipDifficulty{0, false}
	}
	return ClipDifficulty{*d, true}
}

func (cd ClipDifficulty) Value() (driver.Value, error) {
	if !cd.Valid {
		return nil, nil
	}
	return cd.Difficulty, nil
}

func (cd ClipDifficulty) EqualColor(color ProblemColor) bool {
	return cd.Color().Equal(color)
}

func (cd ClipDifficulty) Color() ProblemColor {
	if !cd.Valid {
		return BlackProblem
	}
	tmp := int64(cd.Difficulty) / 400
	if tmp == 0 {
		return GrayProblem
	} else if tmp == 1 {
		return BrownProblem
	} else if tmp == 2 {
		return GreenProblem
	} else if tmp == 3 {
		return CyanProblem
	} else if tmp == 4 {
		return BlueProblem
	} else if tmp == 5 {
		return YellowProblem
	} else if tmp == 6 {
		return OrangeProblem
	} else if tmp == 7 {
		return RedProblem
	} else if tmp == 8 {
		return BronzeProblem
	} else if tmp == 9 {
		return SilverProblem
	} else {
		return GoldProblem
	}
}
