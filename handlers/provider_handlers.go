package handlers

import (
	"belajarGoLang/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	providers  = make(map[string]models.Provider)
	muProvider sync.Mutex
)

func GetProviders(c *gin.Context) {
	muProvider.Lock()
	defer muProvider.Unlock()

	var providerList []models.Provider
	for _, provider := range providers {
		providerList = append(providerList, provider)
	}

	c.JSON(http.StatusOK, providerList)
}
