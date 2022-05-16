package controllers

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_api"
	"diff-problems/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRateController struct {
	Interactor usecase.UserRateInteractor
}

func NewUserRateController(requestHandler api.RequestHandler) *UserRateController {
	return &UserRateController{
		Interactor: usecase.UserRateInteractor{
			ContestResultClient: &atcoder_api.ContestResultClient{
				RequestHandler: requestHandler,
			},
		},
	}
}

func (controller *UserRateController) ShowLatest(c Context) {
	userId := c.Param("user_id")
	entity, err := controller.Interactor.ShowLatest(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (controller *UserRateController) Index(c Context) {
	userId := c.Param("user_id")
	list, err := controller.Interactor.Index(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	c.JSON(http.StatusOK, list)
}
