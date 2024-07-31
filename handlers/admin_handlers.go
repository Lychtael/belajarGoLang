package handlers

import (
	"belajarGoLang/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	admins  = make(map[string]models.Admin)
	muAdmin sync.Mutex
)

func RegisterAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	muAdmin.Lock()
	admins[admin.Email] = admin
	muAdmin.Unlock()

	c.JSON(http.StatusOK, gin.H{"message": "Admin registered successfully"})
}

func LoginAdmin(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	muAdmin.Lock()
	admin, exists := admins[credentials.Email]
	muAdmin.Unlock()
	if !exists || admin.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
