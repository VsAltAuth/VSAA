package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/VsAltAuth/VSAA/models"
	"github.com/VsAltAuth/VSAA/services"
	"github.com/VsAltAuth/VSAA/utils"
)

func GameLogout(c *gin.Context) {
	email := c.PostForm("email")
	sessionkey := c.PostForm("sessionkey")
	user, err := utils.GetUIDBySessionkey(sessionkey)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 1}) // client doesn't care what server returns.
		return                                   // Og VS server seems to always return this value
	}
	usr, err := utils.GetUserByUID(user.UID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"valid": 1})
		return
	}
	if usr.Email != email {
		c.JSON(http.StatusOK, gin.H{"valid": 1})
		return
	}
	if err = utils.RMSession(sessionkey); err != nil {
		fmt.Printf("Failed to rm session: %s", err)
		c.JSON(http.StatusOK, gin.H{"valid": 0}) // for debug purposes
	}
	c.JSON(http.StatusOK, gin.H{"valid": 1})
}

func GameLogin(c *gin.Context) {
	v := c.Param("v")
	if v == "/2/" {
		email := c.PostForm("email")
		password := c.PostForm("password")
		gamever := c.PostForm("gameloginversion")
		var user *models.User
		if err := services.Query(services.DatabaseService, "email", email, &user); err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "dberror"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.HashedPass), []byte(password))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "invalidemailorpassword"})
			return
		}
		if email != user.Email {
			c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "invalidemailorpassword"})
			return
		}
		sessionkey := uuid.NewString()
		sessionkeysig, err := utils.Sign(sessionkey)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "signfail"})
			return
		}
		session := models.Session{UID: user.UID, Sessionkey: sessionkey, Gamever: gamever}
		services.WriteNew(services.CacheService, sessionkey, &session)
		services.WriteCache(services.CacheService, user.UID, user)
		c.JSON(http.StatusOK, gin.H{
			"valid":               1,
			"sessionkey":          sessionkey,
			"sessionkeysignature": sessionkeysig,
			"uid":                 user.UID,
			"playername":          user.Playername,
			"entitlements":        user.Entitlements,
			"prelogintoken":       "",    //TODO: figure out what this is
			"hasgameserver":       false, // Relate to VS's in-house hosting which we do not provide
		})
	}
	c.JSON(http.StatusOK, gin.H{"valid": 0, "reason": "Unsupported API version. Currently only v2 is supported"})
}
