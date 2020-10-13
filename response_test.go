package ginx_test

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/GizmoOAO/ginx"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	app := gin.New()
	app.Use(ginx.Response())
	app.GET("/test", func(c *gin.Context) {
		ginx.R(http.StatusOK, "Hello", "World")
	})
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"Hello","data":"World"}`, w.Body.String())
}

func TestCustomResponse(t *testing.T) {
	app := gin.New()
	app.Use(ginx.Response())

	ginx.CustomResponse(custom{})

	app.GET("/test", func(c *gin.Context) {
		ginx.R(http.StatusOK, "Hello", "World")
	})
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":200,"args":["Hello","World"]}`, w.Body.String())
}

type custom struct {
	StatusCode int           `json:"code"`
	Args       []interface{} `json:"args,omitempty"`
}

func (custom) Custom(code int, args ...interface{}) interface{} {
	return custom{
		StatusCode: code,
		Args:       args,
	}
}
