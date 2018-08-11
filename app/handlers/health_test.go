package handlers

import (
	"testing"
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/urbn/ordernumbergenerator/app/fixtures"
	"github.com/urbn/ordernumbergenerator/app"
	"github.com/urbn/ordernumbergenerator/app/config"
)

func Test_GetHealth(t *testing.T) {
	app.Configuration,_ = config.LoadConfig()

	testRouter := gin.Default()

	hh := HealthHandler{
		Specification: app.Configuration,
	}

	testRouter.GET("/health", hh.GetHealth)

	response := fixtures.PerformRequest(testRouter, "GET", "/health")

	if code := response.Code; code != http.StatusOK {
		t.Errorf("handler returned wrong status code: received %v expected %v",
			code, http.StatusOK)
	}

	expected := app.Configuration.AppName
	if !strings.Contains(response.Body.String(), expected) {
		t.Errorf("handler did not return app name as expected: %v", expected)
	}

}
