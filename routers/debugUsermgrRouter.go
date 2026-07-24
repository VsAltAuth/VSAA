package routers

import (
	"fmt"
	"net/http"

	"github.com/VsAltAuth/VSAA/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// WIP
// does not verify anything or even check for duplicate entries for now
func RegisterNewUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	playername := c.PostForm("playername")
	uid := uuid.NewString()
	hashedpass, err := utils.BHashPass(password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "failed to hash pass"})
		return
	}
	newuser, err := utils.WriteUser(uid, email, hashedpass, playername, "VIV")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": fmt.Errorf("%s", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valid": 1, "uid": newuser.UID, "playername": newuser.Playername})
}
