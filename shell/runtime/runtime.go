package runtime

import (
	"fmt"
	"os"
	"os/user"
)

var (
	// User's home folder at $HOME, shouldn't be changed
	RootPath string = getHomeDir()
	// The current folder pointer, default to $HOME
	CurrentPath string = getHomeDir()
	// TODO: For now hardocded display value as the line prefix
	Username string = getUserName()
)

func getUserName() string {
	u, err := user.Current()
	if err != nil {
		fmt.Printf("error: could not get user, got error: %s", err)
	}
	return u.Username
}

func getHomeDir() string {
	hd, _ := os.UserHomeDir()
	return hd
}
