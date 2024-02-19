package cmd

import (
	"todo-list-app/internal/http/router"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginRouter *gin.Engine, params Params) *gin.Engine {
	group := ginRouter.Group("api/v1")

	// initiate routers, if there is new
	// router's group, just add in this line
	router.NewHealthcheckRoutes(group.Group("/healthz"))

	return ginRouter
}
