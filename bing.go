package ginx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindError(err error) {
	resp := Response{
		StatusCode: http.StatusBadRequest,
		Message:    "bad request",
		Error:      err.Error(),
	}
	panic(resp)
}

func Bind(c *gin.Context, obj interface{}) {
	if err := c.Bind(obj); err != nil {
		bindError(err)
	}
}

func BindJSON(c *gin.Context, obj interface{}) {
	if err := c.BindJSON(obj); err != nil {
		bindError(err)
	}
}

func BindQuery(c *gin.Context, obj interface{}) {
	if err := c.BindQuery(obj); err != nil {
		bindError(err)
	}
}

func BindXML(c *gin.Context, obj interface{}) {
	if err := c.BindXML(obj); err != nil {
		bindError(err)
	}
}

func BindYAML(c *gin.Context, obj interface{}) {
	if err := c.BindYAML(obj); err != nil {
		bindError(err)
	}
}

func ShouldBind(c *gin.Context, obj interface{}) {
	if err := c.ShouldBind(obj); err != nil {
		bindError(err)
	}
}

func ShouldBindJSON(c *gin.Context, obj interface{}) {
	if err := c.ShouldBindJSON(obj); err != nil {
		bindError(err)
	}
}

func ShouldBindQuery(c *gin.Context, obj interface{}) {
	if err := c.ShouldBindQuery(obj); err != nil {
		bindError(err)
	}
}

func ShouldBindXML(c *gin.Context, obj interface{}) {
	if err := c.ShouldBindXML(obj); err != nil {
		bindError(err)
	}
}

func ShouldBindYAML(c *gin.Context, obj interface{}) {
	if err := c.ShouldBindYAML(obj); err != nil {
		bindError(err)
	}
}
