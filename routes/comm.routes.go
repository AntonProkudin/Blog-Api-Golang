package routes

import (
	"github.com/gin-gonic/gin"
	"test/controllers"
	"test/middleware"
)

type CommRouteController struct {
	commController controllers.CommController
}

func NewRouteCommController(commController controllers.CommController) CommRouteController {
	return CommRouteController{commController}
}

func (pc *CommRouteController) CommRoute(rg *gin.RouterGroup) {

	router := rg.Group("comms")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.commController.CreateComm)
	router.DELETE("/", pc.commController.DeleteComm)
	router.GET("/:postId", pc.commController.FindComms)
}
