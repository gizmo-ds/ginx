package ginx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBind(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `form:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		Bind(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader("data=Hello+World"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestBindJSON(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		BindJSON(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"data":"Hello World"}`))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestBindXML(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `xml:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		BindXML(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(
		`<?xml version="1.0" encoding="UTF-8"?><root><data>Hello World</data></root>`,
	))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestBindQuery(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `form:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		BindQuery(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test?data=Hello+World", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestBindYAML(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `yaml:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		BindYAML(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(`data: "Hello World"`))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestShouldBind(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `form:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		ShouldBind(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader("data=Hello+World"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestShouldBindJSON(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		ShouldBindJSON(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(`{"data":"Hello World"}`))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestShouldBindXML(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `xml:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		ShouldBindXML(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(
		`<?xml version="1.0" encoding="UTF-8"?><root><data>Hello World</data></root>`,
	))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestShouldBindQuery(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `form:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		ShouldBindQuery(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test?data=Hello+World", nil)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}

func TestShouldBindYAML(t *testing.T) {
	app := gin.New()
	app.Use(Response())
	type TestType struct {
		Data string `yaml:"data" json:"data" binding:"required"`
	}
	app.POST("/test", func(c *gin.Context) {
		var data TestType
		ShouldBindYAML(c, &data)
		R(http.StatusOK, nil, data)
	})
	req, _ := http.NewRequest(http.MethodPost, "/test", strings.NewReader(`data: "Hello World"`))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"data":{"data":"Hello World"}}`, w.Body.String())
}
