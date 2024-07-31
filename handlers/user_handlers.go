package handlers

import (
	"belajarGoLang/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	users = make(map[string]models.User)
	mu    sync.Mutex
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	users[user.Email] = user
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mu.Lock()
	user, exists := users[credentials.Email]
	mu.Unlock()
	if !exists || user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	mu.Lock()
	user, exists := users[id]
	mu.Unlock()
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	mu.Lock()
	for _, user := range users {
		if user.Name == username {
			mu.Unlock()
			c.JSON(http.StatusOK, user)
			return
		}
	}
	mu.Unlock()

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func GetUserRoles(c *gin.Context) {
	roles := []string{"ADMIN", "USER", "SYSTEM"}
	c.JSON(http.StatusOK, roles)
}
