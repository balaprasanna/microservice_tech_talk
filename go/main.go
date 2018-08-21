package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const APIVERSION string = "v1"
const TIMEFORMAT string = "2006-01-02 15:04:05.9999999"
const SOURCE_URL string = "http://api.pnd.gs/v1/sources/"

type Anything interface{}

type EndPoint struct {
	Popular string `json:"popular"`
	Latest  string `json:"latest"`
}

type Source struct {
	Endpoints EndPoint `json:"endpoints"`
}

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
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "UnKnown"
	}
	data := gin.H{
		"status": "Ok",
		"host":   hostname,
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
	resp, err := http.Get(SOURCE_URL)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer resp.Body.Close()

	var data []Source
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var result = make(map[string][]string)

	for index := 0; index < len(data); index++ {
		result["latest"] = append(result["latest"], data[index].Endpoints.Latest)
		result["popular"] = append(result["popular"], data[index].Endpoints.Popular)
	}

	c.JSON(200, result)
}
