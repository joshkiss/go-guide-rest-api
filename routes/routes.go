package routes

import (
	"go-guide/rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getOneEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	authenticated.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	// server.POST("/events", middlewares.Authenticate, createEvent)
	// server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	// server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)

	// Users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
