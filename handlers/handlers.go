package handlers

import (
	"belajarGoLang/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var loginDetails struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var roleCode string
	query := `
        SELECT role_id
        FROM USERS u
        JOIN ROLES r ON u.role_id = r.id
        WHERE u.email = ? AND u.password = ?
    `
	if err := config.DB.Raw(query, loginDetails.Email, loginDetails.Password).Scan(&roleCode).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if roleCode != "1" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can login"})
		return
	}

	// Successful login, redirect to user page
	c.Redirect(http.StatusFound, "/pageUser")
}
