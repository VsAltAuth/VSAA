package main

import (
	"fmt"
	"time"

	"github.com/VsAltAuth/VSAA/server"
	"github.com/VsAltAuth/VSAA/services"
)

func main() {
	fmt.Println("Hi")
	db := services.DBInit()
	services.InitDatabaseService(db)
	services.InitCacheService(5*time.Minute, 10*time.Minute)
	//TODO: only init this if --start-server arg is passed
	server.InitServerInstance()
}
