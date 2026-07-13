package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GameLogout(c *gin.Context) {
	email := c.PostForm("email")
	sessionkey := c.PostForm("sessionkey")
	user, err := GetUIDBySessionkey(sessionkey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 1}) // client doesn't care what server returns. Og VS server seems to always return this value
		return
	}
	usr, err := GetUserByUID(user.UID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 1}) // see above
		return
	}
	if usr.Email != email {
		c.JSON(http.StatusOK, gin.H{"valid": 1})
		return
	}
	if err = RMSession(sessionkey); err != nil {
		fmt.Printf("Failed to rm session: %s", err)
		c.JSON(http.StatusOK, gin.H{"valid": 0}) // for debug purposes
	}
	c.JSON(http.StatusOK, gin.H{"valid": 1})
}
