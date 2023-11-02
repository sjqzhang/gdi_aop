package controller

import (
	"gdi_aop/dto"
	"github.com/gin-gonic/gin"
	"time"
)

//@router /api
type HelloController struct {
}

//@middleware cache(ttl=10)
//@router /index [get]
func (c HelloController) Index(ctx *gin.Context, req *dto.HelloRequest) interface{} {
	return time.Now().String()
}
