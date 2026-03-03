package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuth_MissingHeader(t *testing.T) {
	gin.SetMode(gin.TestMode); w := httptest.NewRecorder(); c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil); Auth("secret")(c)
	assert.Equal(t, 401, w.Code)
}
