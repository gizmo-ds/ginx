package ginx

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	app.GET("/test", func(c *gin.Context) {
		R(http.StatusOK, "Hello", "World")
	})
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Hello","data":"World"}`, w.Body.String())
}
