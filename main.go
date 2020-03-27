package main
import (
	"net/http"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis"
    "fmt"
)

func main() {
	r := gin.Default()
    client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    
    err = client.Set("hit",1, 0).Err()
    if err != nil {
        fmt.Println(err)
    }


	r.GET("/", func(c *gin.Context) {
    
		c.String(http.StatusOK, "OK")
    })
    
    r.GET("/test", func(c *gin.Context) {
        val, err := client.Incr("hit").Result()

        if err != nil {
            fmt.Println(err)
        }

        c.String(http.StatusOK, "The page has been visited for %d",val)

    })
	r.Run(":8080")
}

func getcounter() {
    
}