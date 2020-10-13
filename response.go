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

	// 返回值应为 struct, 并且第一个参数必须是 StatusCode, 且类型需要为 int
	CustomResponseHandlerFunc func(int, ...interface{}) interface{}
)

var CustomResponse CustomResponseHandlerFunc

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
				default:
					code := http.StatusInternalServerError
					if v := reflect.ValueOf(v).Field(0); v.Kind() == reflect.Int {
						code = int(v.Int())
					}
					c.JSON(code, v)
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}

func R(statusCode int, args ...interface{}) {

	if CustomResponse != nil {
		panic(CustomResponse(statusCode, args...))
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

func Error(err error) {
	panic(err)
}
