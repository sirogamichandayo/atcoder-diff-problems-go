package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var gRouter *gin.Engine

func RouterInitialize() *gin.Engine {
	if gRouter != nil {
		return gRouter
	}

	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	gRouter = gin.Default()
	setApiV1Router(gRouter.Group("api/v1"), config)
	return gRouter
}

func setApiV1Router(v1 *gin.RouterGroup, config *conf.Config) {
	userController := controllers.NewUserController(NewScrapeHandler(), NewRequestHandler())
	userProblemController := controllers.NewUserProblemController(NewSqlHandler(config.SinDb))
	rankNearController := controllers.NewRankNearController()
	v1.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.ApiV1.AllowOrigin},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type"},
	}))

	v1.GET("/users/:user_id", func(c *gin.Context) { userController.Show(c) })
	v1.GET("/users/:user_id/problems", func(c *gin.Context) { userProblemController.Show(c) })
	v1.GET("/ranks/near/:user_id", func(c *gin.Context) { rankNearController.Show(c) })
}
