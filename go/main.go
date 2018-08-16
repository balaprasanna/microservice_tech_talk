package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const APIVERSION string = "v1"
const TIMEFORMAT string = "2006-01-02 15:04:05.9999999"

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", Root)

	v1 := router.Group(fmt.Sprintf("/api/%s", APIVERSION))
	{
		v1.GET("/time", GetTime)
		v1.GET("/source", GetSource)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	router.Run(":" + PORT)
}

//Root: Handler for /
func Root(c *gin.Context) {
	data := gin.H{
		"status": "ok",
	}
	c.JSON(200, data)
}

//GetTime: Handler for /api/v1/time
func GetTime(c *gin.Context) {
	data := gin.H{
		"time": time.Now().Format(TIMEFORMAT),
	}
	c.JSON(200, data)
}

//GetSource: Handler for /api/v1/source
func GetSource(c *gin.Context) {
	data := gin.H{}

	c.JSON(200, data)
}
