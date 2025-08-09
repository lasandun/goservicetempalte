package routes

import (
	"github.com/gin-gonic/gin"
	"test.com/microservice/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.HelloWorld)
	r.GET("/greet", handlers.GreetUser)
	r.GET("/health", handlers.Health)
}
