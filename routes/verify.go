package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rdkvx/desafioTecnico/v2/controllers"
)

func AddRoutes(router *gin.Engine){
	router.POST("/verify", controllers.VerifyPassword)
}