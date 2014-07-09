
package main
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	
	r.GET("/ping", Ping)
	r.GET("/", Index)
	r.GET("/test", Test)
	r.GET("/hello/:name", Hello)
	
	r.Run(":8000")
}

func Hello(c *gin.Context) {
	var name = c.Params.ByName("name")
	var message = "Hello " + name
	c.String(200, message)
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