package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/VsAltAuth/VSAA/server"
)

func StartApplication() error {
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	mkUsrCmd := flag.NewFlagSet("mkusr", flag.ExitOnError)
	namePtr := mkUsrCmd.String("name", "", "username of the user being created")
	passPtr := mkUsrCmd.String("pass", "", "password of the user being created")
	mailPtr := mkUsrCmd.String("email", "", "email of the user being created")

	if len(os.Args) < 2 {
		fmt.Println("No arguments received.")
		return nil
	}
	switch os.Args[1] {

	case "run":
		runCmd.Parse(os.Args[2:]) // required to make it shut up
		server.InitServerInstance()
	case "mkusr":
		mkUsrCmd.Parse(os.Args[2:])
		err := RegisterNewUser(*namePtr, *passPtr, *mailPtr)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("No arguments reveived.")
	}

	return nil
}
