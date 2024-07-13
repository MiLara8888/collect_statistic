package staticstic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		if c.Request.Header["Content-Type"] == nil || len(c.Request.Header["Content-Type"]) == 0 {
			c.AbortWithStatus(204)
			return
		}

		content_type := c.Request.Header["Content-Type"]
		if content_type[0] != "application/json" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// обработка доступа к сервису
func (t *Statistic) HostMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		host := strings.Split(c.Request.Host, ":")
		if len(host) == 0 {
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}
		h := strings.Trim(strings.ToLower(host[0]), " ")

		if ex, ok := t.hostAlowed[h]; !ok || !ex {
			c.AbortWithStatus(http.StatusBadGateway)
			msg := fmt.Sprintf("host not allowed %s", host)
			log.Println(msg)
			log.Println(errors.New(msg))
			return
		}
		c.Next()
	}
}
