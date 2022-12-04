package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rdkvx/desafioTecnico/v2/models"
	"github.com/rdkvx/desafioTecnico/v2/routes"
	"github.com/stretchr/testify/assert"
)

func TestVerifyPassword(t *testing.T) {
	assert := assert.New(t)
	gin.SetMode(gin.TestMode)

	t.Run("failed to check password", func(t *testing.T) {
		expectMsgErr := "failed to check password"

		w := httptest.NewRecorder()
		c, router := gin.CreateTestContext(w)
		router.Use(func(c *gin.Context){})

		routes.AddRoutes(router)

		bodyRequest := models.ValidatePassword{
			Password: "",
		}
		data, _ := json.Marshal(bodyRequest)
		reader := bytes.NewReader(data)

		c.Request, _ = http.NewRequest(http.MethodPost, "/verify", reader)

		router.ServeHTTP(w, c.Request)

		assert.Equal(http.StatusBadRequest, w.Code)
		assert.Equal(expectMsgErr, w.Body.String())
	})
}
