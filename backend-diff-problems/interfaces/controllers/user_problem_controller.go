package controllers

import (
	"diff-problems/domain/vo"
	"diff-problems/interfaces/database"
	queryService "diff-problems/interfaces/query_service"
	"diff-problems/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserProblemController struct {
	Interactor usecase.UserProblemInteractor
}

func NewUserProblemController(sqlHandler database.SqlHandler) *UserProblemController {
	return &UserProblemController{
		Interactor: usecase.UserProblemInteractor{
			UserProblemService: queryService.UserProblemService{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller UserProblemController) Show(c Context) {
	userId := c.Param("user_id")
	entity, err := controller.Interactor.FindByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	solvedPList := entity.SolvedProblemList()
	allPList := entity.AllProblemList()
	c.JSON(http.StatusOK, gin.H{
		"userId": entity.UserId(),
		"problem_count": map[vo.ProblemColor]interface{}{
			vo.BlackProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.BlackProblem),
				"solved": solvedPList.CountByColor(vo.BlackProblem),
			},
			vo.GrayProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.GrayProblem),
				"solved": solvedPList.CountByColor(vo.GrayProblem),
			},
			vo.BrownProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.BrownProblem),
				"solved": solvedPList.CountByColor(vo.BrownProblem),
			},
			vo.GreenProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.GreenProblem),
				"solved": solvedPList.CountByColor(vo.GreenProblem),
			},
			vo.CyanProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.CyanProblem),
				"solved": solvedPList.CountByColor(vo.CyanProblem),
			},
			vo.BlueProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.BlueProblem),
				"solved": solvedPList.CountByColor(vo.BlueProblem),
			},
			vo.YellowProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.YellowProblem),
				"solved": solvedPList.CountByColor(vo.YellowProblem),
			},
			vo.OrangeProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.OrangeProblem),
				"solved": solvedPList.CountByColor(vo.OrangeProblem),
			},
			vo.RedProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.RedProblem),
				"solved": solvedPList.CountByColor(vo.RedProblem),
			},
			vo.BronzeProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.BronzeProblem),
				"solved": solvedPList.CountByColor(vo.BronzeProblem),
			},
			vo.SilverProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.SilverProblem),
				"solved": solvedPList.CountByColor(vo.SilverProblem),
			},
			vo.GoldProblem: map[string]interface{}{
				"all":    allPList.CountByColor(vo.GoldProblem),
				"solved": solvedPList.CountByColor(vo.GoldProblem),
			},
		},
		"userSolveClipDifficultyTotal": entity.SolvedProblemList().ClipDifficultyTotal(),
		"updatedEpochTime":             strconv.FormatInt(entity.UpdatedEpochTime(), 10),
	})
}
