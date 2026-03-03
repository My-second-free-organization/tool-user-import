package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode); w := httptest.NewRecorder(); c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/health", nil); HealthCheck(c)
	assert.Equal(t, 200, w.Code); assert.Contains(t, w.Body.String(), "healthy")
}
