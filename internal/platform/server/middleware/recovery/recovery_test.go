package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	engine.Use(Middleware())
	engine.GET("/test-middleware", func(context *gin.Context) {
		panic("something unexpected")
	})

	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		engine.ServeHTTP(httpRecorder, req)
	})
}
