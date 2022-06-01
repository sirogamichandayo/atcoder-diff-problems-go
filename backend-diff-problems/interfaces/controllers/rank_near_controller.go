package controllers

import (
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RankNearController struct {
	Interactor usecase.RankInteractor
}

func NewRankNearController(sqlHandler database.SqlHandler) *RankNearController {
	return &RankNearController{
		Interactor: usecase.RankInteractor{
			UserSolveProblemDifficultySumRepository: database.UserSolveProblemDifficultySumRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RankNearController) Show(c Context) {
	userId := c.Param("user_id")
	count, err := strconv.Atoi(c.DefaultQuery("count", "5"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	list, err := controller.Interactor.Near(userId, count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	var res []gin.H
	for _, entity := range list.List() {
		res = append(res, gin.H{
			"userId":            entity.UserId,
			"clipDifficultySum": entity.ClipDifficultySum,
			"rank":              entity.Rank,
		})
	}
	c.JSON(http.StatusOK, res)
}
