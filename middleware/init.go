package middleware

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var middlewaresMap = make(map[string]func(p sync.Map) gin.HandlerFunc)

func registerMiddleware(name string, middleware func(p sync.Map) gin.HandlerFunc) {
	middlewaresMap[name] = middleware
}

func GetMiddleware(name string) func(p sync.Map) gin.HandlerFunc {
	return middlewaresMap[name]
}
