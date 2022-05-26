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

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(scrapeHandler web.ScrapeHandler, requestHandler api.RequestHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserService: queryService.UserService{
				ScrapeHandler:       scrapeHandler,
				ContestResultClient: atcoder_api.ContestResultClient{RequestHandler: requestHandler},
			},
		},
	}
}

func (controller *UserController) Show(c Context) {
	userId := c.Param("user_id")
	entity, err := controller.Interactor.FindByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
