
package main
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	
	r.GET("/ping", Ping)
	r.GET("/", Index)
	r.GET("/test", Test)
	
	r.Run(":8000")
}

func Ping(c *gin.Context) {
	c.String(200, "Pong")
}

func Index(c *gin.Context) {
	c.String(200, "Index")
}

func Test(c *gin.Context) {
	c.String(200, "Test")
}