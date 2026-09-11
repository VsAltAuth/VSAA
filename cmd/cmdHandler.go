package cmd

import (
	"fmt"

	"github.com/VsAltAuth/VSAA/utils"
	"github.com/google/uuid"
)

func RegisterNewUser(name string, pass string, email string) error {
	if name == "" {
		return fmt.Errorf("Username must be set!")
	} else if pass == "" {
		return fmt.Errorf("Password must be set!")
	} else if email == "" {
		return fmt.Errorf("Email must be set!")
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
	fmt.Printf("Created new user with uid %s, playername %s", newuser.UID, newuser.Playername)
	return nil
}
