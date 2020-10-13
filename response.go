package ginx

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
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

func Ginx() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case Response:
					c.JSON(v.StatusCode, v)
					c.Abort()
				case error:
					c.JSON(http.StatusInternalServerError,
						Response{Error: "internal server error"})
					// log.Println(v)
					c.Abort()
				case string:
					c.JSON(http.StatusInternalServerError,
						Response{Error: v})
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

	resp := Response{
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

func CustomResponse(custom CustomInterface) {
	if custom != nil {
		customResponse = custom
	}
}

func Error(err error) {
	panic(err)
}
