package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rdkvx/desafioTecnico/v2/models"
	"github.com/rdkvx/desafioTecnico/v2/services"
	"github.com/rdkvx/desafioTecnico/v2/utils"
)

func VerifyPassword(c *gin.Context) {
	bodyRequest := models.ValidatePassword{}
	if err := c.ShouldBindBodyWith(&bodyRequest, binding.JSON); err != nil {
		utils.SendJSONError(c, http.StatusBadRequest, fmt.Errorf("%v: %v", utils.FailedtoCheckPassword, err))
	}

	verify := true
	failedRules := services.ValidatePassword(bodyRequest)
	if len(failedRules) > 0 {
		verify = false
	}

	response := models.JsonResponse{
		Verify: verify,
		NoMatch: failedRules,
	}

	utils.SendJSONResponse(c, http.StatusCreated, response)
}
