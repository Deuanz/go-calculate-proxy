package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func calculate(c *gin.Context) {
	request, _ := c.GetRawData()
	calculateUrl := "http://0.0.0.0:8081" + string(c.Request.RequestURI)
	resp, err := http.Post(calculateUrl, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return
	}
	var result json.RawMessage
	json.NewDecoder(resp.Body).Decode(&result)

	c.JSON(resp.StatusCode, result)
}

func main() {

	r := gin.Default()
	r.POST("/calculator.sum", calculate)
	r.POST("/calculator.sub", calculate)
	r.POST("/calculator.mul", calculate)
	r.POST("/calculator.div", calculate)

	r.Run(":8080")
}
