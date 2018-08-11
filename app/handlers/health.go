package handlers

import (
	"github.com/urbn/ordernumbergenerator/app"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	Specification *app.Specification
}

func (hh HealthHandler) GetHealth(c *gin.Context) {

	s := app.HealthResponse{
		TimeStamp:   time.Now().UTC().Format(time.RFC822Z),
		Status:      "ok",
		AppName:     hh.Specification.AppName,
		Branch:      hh.Specification.Branch,
		BuildNumber: hh.Specification.BuildNumber,
		GitHash:     hh.Specification.GitHash,
		Environment: hh.Specification.Environment,
	}

	c.JSON(http.StatusOK, gin.H{"status": s})
}
