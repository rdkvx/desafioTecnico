package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rdkvx/desafioTecnico/v2/mock"
	"github.com/rdkvx/desafioTecnico/v2/models"
	"github.com/rdkvx/desafioTecnico/v2/routes"
	"github.com/rdkvx/desafioTecnico/v2/services"
	"github.com/rdkvx/desafioTecnico/v2/utils"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
)

func TestVerifyPassword(t *testing.T) {
	assert := assert.New(t)
	gin.SetMode(gin.TestMode)

	t.Run("failed to check password", func(t *testing.T) {
		expectMsgErr := string("{\"Verify\":false,\"NoMatch\":[\"minSize\"]}")

		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(func(c *gin.Context) {})

		routes.AddRoutes(router)

		
		bodyRequest := models.ValidatePassword{
			Password: "abc",
		}
		data, _ := json.Marshal(bodyRequest)
		reader := bytes.NewReader(data)

		failedRules := []string{}
		failedRules = append(failedRules, utils.MinSize)
		
		iValidatePasswordMock := new(mock.IpasswordValidator)
		iValidatePasswordMock.On("ValidatePassword", tmock.Anything, tmock.Anything).Return(failedRules)
		services.PasswordValidator = iValidatePasswordMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/verify", reader)

		router.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusCreated, w.Code)
		assert.Equal(expectMsgErr, w.Body.String())
	})

	t.Run("success!", func(t *testing.T) {
		expectMsgErr := string("{\"Verify\":true,\"NoMatch\":[]}")

		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(func(c *gin.Context) {})

		routes.AddRoutes(router)

		type rule struct {
			Rule  string `json:"rule"`
			Value int    `json:"value"`
		}

		
		bodyRequest := models.ValidatePassword{
			Password: "abc",
		}

		newRule := rule{
			Rule: "minSize",
			Value: 1,
		}

		bodyRequest.Rules = append(bodyRequest.Rules, newRule)
		data, _ := json.Marshal(bodyRequest)
		reader := bytes.NewReader(data)

		failedRules := []string{}
		
		iValidatePasswordMock := new(mock.IpasswordValidator)
		iValidatePasswordMock.On("ValidatePassword", tmock.Anything, tmock.Anything).Return(failedRules)
		services.PasswordValidator = iValidatePasswordMock

		c.Request, _ = http.NewRequest(http.MethodPost, "/verify", reader)

		router.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusCreated, w.Code)
		assert.Equal(expectMsgErr, w.Body.String())
	})
}
