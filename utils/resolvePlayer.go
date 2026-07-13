package utils

import (
	"fmt"
	"net/http"

	"github.com/VsAltAuth/VSAA/services"
	"github.com/gin-gonic/gin"
)

func ResolveUIDByPlayername(c *gin.Context) {
	playername := c.PostForm("playername")
	user, err := services.CacheService.GetUserByPlayername(playername)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"playeruid": nil, "valid": 1}) // Original VS auth server returns valid as 1 even when the playername is invalid
		return
	}
	c.JSON(http.StatusOK, gin.H{"playeruid": user.UID, "valid": 1})
}

func ResolvePlayernameByUID(c *gin.Context) {
	uid := c.PostForm("uid")
	user, err := services.CacheService.GetUserByUID(uid)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"playername": nil, "valid": 1}) // same as above
		return
	}
	c.JSON(http.StatusOK, gin.H{"playername": user.Playername, "valid": 1})
}
