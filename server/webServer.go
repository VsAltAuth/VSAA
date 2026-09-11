package server

import (
	"net/http"

	"github.com/VsAltAuth/VSAA/routers"
	"github.com/VsAltAuth/VSAA/utils"
	"github.com/gin-gonic/gin"
)

func InitServerInstance() {
	router := gin.Default()
	//router.TrustedPlatform = gin.PlatformCloudflare

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hi"}) })
	router.POST("/resolveplayername", routers.ResolveUIDByPlayername)
	router.POST("/resolveplayeruid", routers.ResolvePlayernameByUID)
	router.POST("/clientvalidate", routers.ClientValidate)
	router.POST("/gamelogout", routers.GameLogout)
	router.POST("/resolveserverhost", routers.ResolveServerHost)
	router.POST("/:v/gamelogin", routers.GameLogin)

	router.GET("/publickeypem", func(c *gin.Context) { c.String(http.StatusOK, utils.GetPublicKey()) })

	router.Run("localhost:8080")

}
