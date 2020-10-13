package ginx

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type (
	Resp struct {
		StatusCode int         `json:"-"`
		Error      string      `json:"error,omitempty"`
		Message    string      `json:"message,omitempty"`
		Data       interface{} `json:"data,omitempty"`
	}
	CustomInterface interface {
		Custom(int, ...interface{}) interface{}
	}
)

var customResponse CustomInterface

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case Resp:
					c.JSON(v.StatusCode, v)
					c.Abort()
				case error:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "internal server error",
					})
					log.Println(v)
					c.Abort()
				case string:
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": v,
					})
					c.Abort()
				case CustomInterface:
					c.JSON(int(reflect.ValueOf(v).Field(0).Int()), v)
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}

func R(statusCode int, args ...interface{}) {

	if customResponse != nil {
		panic(customResponse.Custom(statusCode, args...))
	}

	resp := Resp{
		StatusCode: statusCode,
	}
	if len(args) >= 1 && args[0] != nil {
		switch v := args[0].(type) {
		case error:
			resp.Error = v.Error()
		case string:
			resp.Message = v
		}
	}
	if len(args) >= 2 {
		resp.Data = args[1]
	}
	panic(resp)
}

type response1 struct {
	StatusCode int         `json:"code"`
	Error      string      `json:"error,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func CustomResponse(custom CustomInterface) {
	if custom != nil {
		customResponse = custom
	}
}
