package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
    })
    
    r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello I am the new Api")
    })
	r.Run(":8080")
}