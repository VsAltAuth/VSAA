package cmd

import (
	"fmt"

	"github.com/VsAltAuth/VSAA/utils"
	"github.com/google/uuid"
)

func RegisterNewUser(name string, pass string, email string) error {
	if name == "" {
		fmt.Println("Username must be set!")
		return nil
	} else if pass == "" {
		fmt.Println("Password must be set!")
		return nil
	} else if email == "" {
		fmt.Println("Email must be set!")
		return nil
	}
	uid := uuid.NewString()
	hashedpass, salt, err := utils.HashPass(pass)
	if err != nil {
		return err
	}
	newuser, err := utils.WriteUser(uid, email, hashedpass, salt, name, "VIV")
	if err != nil {
		return err
	}
	fmt.Printf("Created new user with uid %s, playername %s\n", newuser.UID, newuser.Playername)
	return nil
}
