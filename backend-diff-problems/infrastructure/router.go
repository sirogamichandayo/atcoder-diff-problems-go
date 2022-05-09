package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Initialize() {
	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	setApiV1Router(router.Group("api/v1"), config)
	Router = router
}

func setApiV1Router(v1 *gin.RouterGroup, config *conf.Config) {
	userController := controllers.NewUserController(NewSqlHandler(config.SinDb))
	v1.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.ApiV1.AllowOrigin},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	v1.POST("/users", func(c *gin.Context) { userController.Create(c) })
	v1.GET("/users", func(c *gin.Context) { userController.Index(c) })
	v1.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
}
