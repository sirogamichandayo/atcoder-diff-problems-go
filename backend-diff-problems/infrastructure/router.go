package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	userController := controllers.NewUserController(NewSqlHandler(config.SinDb))

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
