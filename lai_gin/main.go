package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Geektutu")
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s\n", name)
	})

	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "student")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		passwd := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"username": name,
			"password": passwd})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
