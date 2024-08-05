package routes

import (
	"belajarGoLang/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set static folder
	r.Static("/static", "./static")

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// User routes
	r.POST("/login", handlers.LoginHandler)

	// Serve the user page
	r.GET("/pageUser", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pageUser.html", gin.H{})
	})

	// Root URL
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	return r
}
