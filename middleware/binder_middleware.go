package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()






func BinderMiddleware(method reflect.Value) gin.HandlerFunc {

	return func(c *gin.Context) {

		if method.Type().NumIn() != 2 {
			c.JSON(400, gin.H{"message": "请求参数错误"})
			return
		}
		reqType := method.Type().In(1)
		var err error

		req := reflect.New(reqType.Elem()).Interface()
		// 获取请求参数 Get
		if c.Request.Method == "GET" {
			err := c.ShouldBindQuery(req)
			if err != nil {
				fmt.Println(req)
				c.JSON(400, gin.H{"message": fmt.Sprintf("请求参数错误%v", err)})
				return
			}
		}
		// 绑定参数

		if (c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" ||
			c.Request.Method == "PATCH") && c.Request.ContentLength > 0 {
			err := c.ShouldBindJSON(req)
			if err != nil {
				fmt.Println(req)
				c.JSON(400, gin.H{"message": "请求参数错误" + err.Error()})
				return
			}
		}
		err = validate.Struct(req)
		if err != nil {

			c.JSON(400, gin.H{"message": "请求参数错误" + err.Error()})
			return
		}
		c.Set("REQ-INPUT", req)
		c.Set("URI", c.Request.RequestURI)

		results := method.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(req)})
		if len(results) > 0 {
			if len(results) == 2 {
				if results[1].Interface() != nil {
					c.JSON(400, gin.H{"message": results[1].Interface().(error).Error()})
					return
				}
			}
			if results[0].Interface() != nil {

				c.JSON(400, gin.H{"data": results[0].Interface(), "message": "ok"})
				return
			}
		}
		c.Next()

	}
}
