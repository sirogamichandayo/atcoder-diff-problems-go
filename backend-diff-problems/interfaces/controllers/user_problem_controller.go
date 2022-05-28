package controllers

import (
	"diff-problems/interfaces/database"
	queryService "diff-problems/interfaces/query_service"
	"diff-problems/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"userId":                       entity.UserId(),
		"imageUrl":                     entity.ImageUrl(),
		"userSolveClipDifficultyTotal": entity.SolvedProblemList().ClipDifficultyTotal(),
		"clipDifficultyTotal":          entity.ClipDifficultyTotal(),
		"updatedEpochTime":             entity.UpdatedEpochTime(),
	})
}
