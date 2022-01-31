// Created by EldersJavas(EldersJavas&gmail.com)
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := gin.Default()
	router.GET("/",func(c *gin.Context) {
		firstname := c.DefaultQuery("p", "") //如果没有值，还可以给一个默认值
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s ", firstname, lastname)
	})
	// Listen and serve on 0.0.0.0:80
	router.Run(":8024")
}
