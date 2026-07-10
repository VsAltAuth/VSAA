package main

import (
	"fmt"
	"net/http"

	//"github.com/VsAltAuth/VSAA/services"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi")
	router := gin.Default()
	//router.TrustedPlatform = gin.PlatformCloudflare
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hi"}) })
	router.GET("/resolveplayername")

	router.Run("localhost:8080")
}
