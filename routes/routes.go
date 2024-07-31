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
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/users/:id", handlers.GetUserByID)
	r.GET("/users/username/:username", handlers.GetUserByUsername)
	r.GET("/users/roles", handlers.GetUserRoles)

	// Admin routes
	r.POST("/admin/register", handlers.RegisterAdmin)
	r.POST("/admin/login", handlers.LoginAdmin)

	// Provider routes
	r.GET("/providers", handlers.GetProviders)

	// Render the login page
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	return r
}
