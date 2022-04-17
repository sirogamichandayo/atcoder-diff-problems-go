package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

var Router *gin.Engine

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config, err := conf.LoadConfig(path.Join(currentDir, "config"))
	if err != nil {
		panic(err)
	}

	userController := controllers.NewUserController(NewSqlHandler(config.SinDb))

	router := gin.Default()
	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	Router = router
}
