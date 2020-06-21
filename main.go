package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func reverseProxy(c *gin.Context) {

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "0.0.0.0:8081",
	})

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {

	r := gin.Default()
	r.POST("/calculator.sum", reverseProxy)
	r.POST("/calculator.sub", reverseProxy)
	r.POST("/calculator.mul", reverseProxy)
	r.POST("/calculator.div", reverseProxy)

	r.Run(":8080")
}
