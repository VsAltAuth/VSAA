package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResolveServerHost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"valid": 1,  // VS og server appears to always set this value to 1
		"host":  "", // Related to VS in-house hosting, which we do not provide
	})
}

/*
 It appears that when trying to connect to a gameserver which starts with "vh" client sends a
 request to this endpoint to get the real hostname. We will return an empty value
 which the client should interpret as "server not found"
*/
