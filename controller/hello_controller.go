package controller

import (
	"gdi_aop/dto"
	"github.com/gin-gonic/gin"
)

//@router /api
type HelloController struct {
}

//@middleware cache
//@router /index [get]
func (c HelloController) Index(ctx *gin.Context,req *dto.HelloRequest) interface{} {

	return "hello"
}
