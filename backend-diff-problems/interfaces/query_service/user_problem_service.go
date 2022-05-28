package queryService

import (
	"diff-problems/domain/entity"
	"diff-problems/domain/vo"
	"diff-problems/interfaces/database"
	cqrsDto "diff-problems/usecase/cqrs_dto"
	cqrsService "diff-problems/usecase/cqrs_service"
)

type UserProblemService struct {
	SqlHandler database.SqlHandler
}

func NewUserProblemService(handler database.SqlHandler) cqrsService.UserProblemService {
	return UserProblemService{handler}
}

func (s UserProblemService) FindByUserId(userId string) (cqrsDto.UserProblem, error) {
	solvedProblemIdList, err := s.getSolvedProblemIdList(userId)
	if err != nil {
		return cqrsDto.UserProblem{}, err
	}

	clipDifficultyTotal, err := s.getClipDifficultyTotal()
	if err != nil {
		return cqrsDto.UserProblem{}, err
	}

	updatedEpochTime, err := s.getUpdatedEpochTime()
	if err != nil {
		return cqrsDto.UserProblem{}, err
	}

	return cqrsDto.NewUserProblem(userId, solvedProblemIdList, clipDifficultyTotal, updatedEpochTime), nil
}

func (s UserProblemService) getUpdatedEpochTime() (int64, error) {
	var updatedEpochTime int64
	row, err := s.SqlHandler.Query("SELECT updated_epoch_time FROM user_first_ac_submission_updated_at LIMIT 1")
	if err != nil {
		return updatedEpochTime, err
	}
	defer row.Close()
	row.Next()
	err = row.Scan(&updatedEpochTime)

	return updatedEpochTime, err
}

func (s UserProblemService) getSolvedProblemIdList(userId string) (entity.ProblemDifficultyList, error) {
	rows, err := s.SqlHandler.Query(`
select u.problem_id, difficulty, clip_difficulty
from user_first_ac_submissions as u
join problem_difficulties as p
on u.problem_id = p.problem_id
where user_id = ?;
`, userId)
	if err != nil {
		return nil, err
	}
	rows.Close()

	list := make(entity.ProblemDifficultyList, 0)
	for rows.Next() {
		var pId string
		var d, cd *float64
		if err := rows.Scan(&pId, &d, &cd); err != nil {
			return nil, err
		}

		list = append(list,
			entity.ProblemDifficulty{
				ProblemId:      pId,
				Difficulty:     vo.RawDifficulty{Difficulty: *d, Valid: d != nil},
				ClipDifficulty: vo.ClipDifficulty{Difficulty: *cd, Valid: cd != nil},
			},
		)
	}

	return list, nil
}

func (s UserProblemService) getClipDifficultyTotal() (float64, error) {
	var total float64
	row, err := s.SqlHandler.Query("SELECT SUM(IFNULL(clip_difficulty, 0)) FROM problem_difficulties;")
	if err != nil {
		return total, err
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, err
}
