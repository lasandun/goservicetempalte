package handlers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}

func GreetUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "Guest"
	}
	c.JSON(200, gin.H{"message": "Hello, " + name + "!"})
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"success": "true"})
}
