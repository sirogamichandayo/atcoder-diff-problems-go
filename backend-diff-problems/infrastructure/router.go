package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

var Router *gin.Engine

func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config, err := conf.LoadConfig(currentDir)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	makeApiV1(router, config)
	Router = router
}

func makeApiV1(router *gin.Engine, config *conf.Config) {
	v1 := router.Group("api/v1")
	userController := controllers.NewUserController(NewSqlHandler(config.SinDb))
	v1.Use(cors.New(cors.Config{
		AllowOrigins: config.ApiV1.AllowOrigins,
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type"},
	}))
	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })
	v1.GET("/users", func(c *gin.Context) { userController.Index(c) })
	v1.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
}
