package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine


func main(){
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes(router)

	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("accept") {

	case "application/josn":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])

	default:
		c.HTML(http.StatusOK, templateName, data)
	}
	
}