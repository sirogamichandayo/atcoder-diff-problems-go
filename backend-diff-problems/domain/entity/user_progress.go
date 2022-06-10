package entity

import (
	"diff-problems/domain/vo"
	"fmt"
	"time"
)

type progressType = map[time.Time]vo.ClipDifficultySum

type UserProgressList struct {
	progress progressType
}

func (list *UserProgressList) Add(t time.Time, difficulty vo.ClipDifficulty) error {
	if list == nil {
		return fmt.Errorf("UserProgressList is nil")
	}
	if list.progress == nil {
		list.progress = make(progressType, 0)
	}
	newT := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(newT)

	p := list.progress[newT]
	if err := p.Add(difficulty); err != nil {
		return err
	}
	list.progress[newT] = p

	return nil
}

func (list UserProgressList) Progress() progressType {
	return list.progress
}
