package controllers

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_api"
	"diff-problems/interfaces/query_service"
	"diff-problems/interfaces/web"
	"diff-problems/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserProgressController struct {
	Interactor usecase.UserProgressInteractor
}

func NewUserProgressController(scrapeHandler web.ScrapeHandler, requestHandler api.RequestHandler) *UserController {
	return &UserProgressController{
		Interactor: usecase.UserProgressInteractor{

			UserService: queryService.UserService{
				ScrapeHandler:       scrapeHandler,
				ContestResultClient: atcoder_api.ContestResultClient{RequestHandler: requestHandler},
			},
		},
	}
}

func (controller *UserProgressController) Show(c Context) {
	userId := c.Param("user_id")
	entity, err := controller.Interactor.exec(userId)
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
}
