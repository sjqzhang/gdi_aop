
## GO AOP 注解示例

```go
package controller
import (
	"gdi_aop/dto"
	"github.com/gin-gonic/gin"
	"time"
)
//@router /api
type HelloController struct {
}

//@middleware log;cache(ttl=10)
//@router /index [get]
func (c HelloController) Index(ctx *gin.Context, req *dto.HelloRequest) interface{} {
	return time.Now().String()
}

```
上面的示例中使用日志和缓存注解 cache设定10秒过期,多个注解之间用;进行分隔
