package infrastructure

import (
	conf "diff-problems/config"
	"diff-problems/interfaces/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func RouterInitialize() {
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

	v1.GET("/users/:user_id/diff-sum", func(c *gin.Context) { userController.ShowDiff(c) })
}
