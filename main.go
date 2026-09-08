package main

import (
	"fmt"
	"time"

	"github.com/VsAltAuth/VSAA/cmd"
	"github.com/VsAltAuth/VSAA/services"
)

func main() {
	fmt.Println("Hi")
	db := services.DBInit()
	services.InitDatabaseService(db)
	services.InitCacheService(5*time.Minute, 10*time.Minute)
	if err := cmd.StartApplication(); err != nil {
		panic(err)
	}
}
