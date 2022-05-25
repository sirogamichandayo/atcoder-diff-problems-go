package controllers

import (
	"diff-problems/interfaces/api"
	"diff-problems/interfaces/api/atcoder_api"
	"diff-problems/interfaces/database"
	"diff-problems/interfaces/query_service"
	"diff-problems/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler, requestHandler api.RequestHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserService: queryService.UserService{

				ContestResultClient: atcoder_api.ContestResultClient{
					RequestHandler: requestHandler,
				},
			},
		},
	}
}

func (controller *UserController) ShowDiff(c Context) {
	//	userId := c.Param("user_id")
	panic("todo implement")
	entity, err := controller.Interactor.DiffRankById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}
