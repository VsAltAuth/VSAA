package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VsAltAuth/VSAA/services"
	"github.com/VsAltAuth/VSAA/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")
	db := services.DBInit()
	services.InitCacheService(db, 5*time.Minute, 10*time.Minute)
	router := gin.Default()
	//router.TrustedPlatform = gin.PlatformCloudflare
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hi"}) })
	router.POST("/resolveplayername", utils.ResolveUIDByPlayername)
	router.POST("/resolveplayeruid", utils.ResolvePlayernameByUID)

	router.Run("localhost:8080")
}
