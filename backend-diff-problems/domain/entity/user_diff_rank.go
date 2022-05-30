package entity

type UserDiffRank struct {
	userId            string
	clipDifficultySum float64
	rank              int64
}

func (u UserDiffRank) ClipDifficultySum() float64 {
	return u.clipDifficultySum
}

func (u UserDiffRank) Rank() int64 {
	return u.rank
}

func (u UserDiffRank) UserId() string {
	return u.userId
}

func NewUserDiffRank(userId string, clipDifficultySum float64, rank int64) UserDiffRank {
	return UserDiffRank{userId, clipDifficultySum, rank}
}
