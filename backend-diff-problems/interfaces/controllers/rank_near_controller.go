package controllers

import (
	"diff-problems/usecase"
)

type RankNearController struct {
	Interactor usecase.UserInteractor
}

func NewRankNearController() *RankNearController {
	return &RankNearController{}
}

func (controller *RankNearController) Show(c Context) {
	// 	userId := c.Param("user_id")
	// count, err := strconv.Atoi(c.DefaultQuery("count", "5"))
	/*
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
			return
		}

		entity, err := controller.Interactor.FindByUserId(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"userId":   entity.UserId(),
			"imageUrl": entity.ImageUrl(),
			"ranking":  entity.Ranking(),
			"rating":   entity.Rating().Rating(),
			"color":    entity.Rating().Color(),
		})
	*/
}
