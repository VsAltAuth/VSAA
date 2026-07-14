package routers

import (
	"net/http"

	"github.com/VsAltAuth/VSAA/utils"
	"github.com/gin-gonic/gin"
)

func ClientValidate(c *gin.Context) {
	sessionKey := c.PostForm("sessionkey")
	uid := c.PostForm("uid")
	user, err := utils.GetUIDBySessionkey(sessionKey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "nosession"})
		return
	}
	if user.UID == uid {
		usr, err := utils.GetUserByUID(uid)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"valid": 1, "reason": "", "entitlements": usr.Entitlements, "hasgameserver": false}) // hasgameserver likely reffers to VS's in-house hosting service. We don't provide one so always return false
	} else {
		c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "noaccount"})
	}
}
