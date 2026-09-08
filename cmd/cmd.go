package cmd

import (
	"flag"

	"github.com/VsAltAuth/VSAA/server"
)

func StartApplication() error {
	initSrvrFlagPtr := flag.Bool("startserver", false, "use this argument if you wish to start the http server")

	flag.Parse()

	if *initSrvrFlagPtr == true {
		server.InitServerInstance()
	}
	return nil
}
