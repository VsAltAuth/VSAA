package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/VsAltAuth/VSAA/server"
)

func StartApplication() error {
	mkUsrCmd := flag.NewFlagSet("mkusr", flag.ExitOnError)
	namePtr := mkUsrCmd.String("name", "", "username of the user being created")
	passPtr := mkUsrCmd.String("pass", "", "password of the user being created")
	mailPtr := mkUsrCmd.String("email", "", "email of the user being created")
	rmUsrCmd := flag.NewFlagSet("mrmusr", flag.ExitOnError)
	uidPtr := rmUsrCmd.String("uid", "", "uid of user to delete")

	if len(os.Args) < 2 {
		fmt.Println("No arguments received.")
		return nil
	}
	switch os.Args[1] {

	case "run":
		server.InitServerInstance()
	case "mkusr":
		mkUsrCmd.Parse(os.Args[2:])
		err := registerNewUser(*namePtr, *passPtr, *mailPtr)
		if err != nil {
			return err
		}
	case "rmusr":
		rmUsrCmd.Parse(os.Args[2:])
		err := removeUser(*uidPtr)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("No arguments reveived.")
	}

	return nil
}
