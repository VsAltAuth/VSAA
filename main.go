package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VsAltAuth/VSAA/routers"
	"github.com/VsAltAuth/VSAA/services"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")
	db := services.DBInit()
	services.InitDatabaseService(db)
	services.InitCacheService(5*time.Minute, 10*time.Minute)
	router := gin.Default()
	//router.TrustedPlatform = gin.PlatformCloudflare
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hi"}) })
	router.POST("/resolveplayername", routers.ResolveUIDByPlayername)
	router.POST("/resolveplayeruid", routers.ResolvePlayernameByUID)
	router.POST("/clientvalidate", routers.ClientValidate)
	router.POST("/gamelogout", routers.GameLogout)

	router.Run("localhost:8080")
}
