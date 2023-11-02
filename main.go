package main

import (
	"gdi_aop/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sjqzhang/gdi"
)

func main() {

	gdi.GenGDIRegisterFile(true)
	ctrls, err := gdi.AutoRegisterByPackagePatten("controller")
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	routerInfo, err := gdi.GetRouterInfoByPatten("controller")
	for _, info := range routerInfo {
		for _, ctrl := range ctrls {
			if ctrl.Type().String() == "*"+info.Controller {
				method := ctrl.MethodByName(info.Handler)
				router.Handle(info.Method, info.Uri, middleware.BinderMiddleware(method))
			}
		}
	}
	router.Run(":8085")
}
