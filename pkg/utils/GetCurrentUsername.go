// file: pkg/utils/GetCurrentUsername.go
// (c) 2025 Asymmetric Effort, LLC. <scaldwell@asymmetric-effort.com>

package utils

import (
	"log"
	"os/user"
)

// GetCurrentUsername returns the username of the user executing the current process.
// It returns an error if the username cannot be determined.
func GetCurrentUsername() (username string) {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("error getting username: %v", err)
	}
	return u.Username
}
