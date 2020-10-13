package ginx

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	StatusCode int         `json:"-"`
	Error      string      `json:"error,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case response:
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
				}
			}
		}()
		c.Next()
	}
}

func R(statusCode int, args ...interface{}) {
	resp := response{
		StatusCode: statusCode,
	}
	if len(args) >= 1 && args[0] != nil {
		switch v := args[0].(type) {
		case error:
			resp.Error = v.Error()
		default:
			resp.Message = v
		}
	}
	if len(args) >= 2 {
		resp.Data = args[1]
	}
	panic(resp)
}
