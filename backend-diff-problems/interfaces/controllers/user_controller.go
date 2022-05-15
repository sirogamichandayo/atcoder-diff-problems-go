package controllers

import (
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserSolveProblemDifficultySumRepository: &database.UserSolveProblemDifficultySumRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) ShowDiff(c Context) {
	userId := c.Param("user_id")
	user, err := controller.Interactor.DiffRankById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
