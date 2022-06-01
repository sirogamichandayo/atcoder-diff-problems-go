package controllers

import (
	"diff-problems/interfaces/database"
	"diff-problems/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RankController struct {
	Interactor usecase.RankInteractor
}

func NewRankController(sqlHandler database.SqlHandler) *RankController {
	return &RankController{
		Interactor: usecase.RankInteractor{
			UserSolveProblemDifficultySumRepository: database.UserSolveProblemDifficultySumRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

const DefaultLimit = "0"
const DefaultOffset = "20"

func (controller *RankController) Show(c Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", DefaultOffset))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", DefaultLimit))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	list, err := controller.Interactor.Paging(offset, limit)
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
