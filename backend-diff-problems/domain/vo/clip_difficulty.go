package vo

import "database/sql/driver"

type ClipDifficulty struct {
	Difficulty float64
	Valid      bool
}

func (cd ClipDifficulty) Value() (driver.Value, error) {
	if !cd.Valid {
		return nil, nil
	}
	return cd.Difficulty, nil
}
