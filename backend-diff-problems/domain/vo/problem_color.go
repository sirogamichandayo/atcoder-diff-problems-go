package vo

type ProblemColor string

const (
	BlackProblem  ProblemColor = "Black"
	GrayProblem   ProblemColor = "Gray"
	BrownProblem  ProblemColor = "Brown"
	GreenProblem  ProblemColor = "Green"
	CyanProblem   ProblemColor = "Cyan"
	BlueProblem   ProblemColor = "Blue"
	YellowProblem ProblemColor = "Yellow"
	OrangeProblem ProblemColor = "Orange"
	RedProblem    ProblemColor = "Red"
	BronzeProblem ProblemColor = "Bronze"
	SilverProblem ProblemColor = "Silver"
	GoldProblem   ProblemColor = "Gold"
)

func (pc ProblemColor) Equal(color ProblemColor) bool {
	return pc == color
}
