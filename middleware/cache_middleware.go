package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"

	"strconv"
	"sync"
	"time"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyCache *bytes.Buffer
}

var cache sync.Map

// rewrite Write()
func (w bodyWriter) Write(b []byte) (int, error) {
	w.bodyCache.Write(b)
	return w.ResponseWriter.Write(b)
}

func init() {
	registerMiddleware("cache", CacheMiddleware)
}

func CacheMiddleware(paramMap sync.Map) gin.HandlerFunc {

	var ttl int64

	if v, ok := paramMap.Load("ttl"); ok {
		ttl, _ = strconv.ParseInt(v.(string), 10, 64)
	} else {
		ttl = 10
	}

	return func(c *gin.Context) {

		if val, ok := cache.Load(c.Request.RequestURI); ok {
			c.Writer.Write(val.([]byte))
			c.Abort()
			return
		}
		c.Writer = &bodyWriter{bodyCache: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Next()
		if c.Writer.Status() < 300 {
			cache.Store(c.Request.RequestURI, c.Writer.(*bodyWriter).bodyCache.Bytes())
			go func() {
				time.Sleep(time.Duration(ttl) * time.Second)
				cache.Delete(c.Request.RequestURI)
			}()

		}
	}

}
