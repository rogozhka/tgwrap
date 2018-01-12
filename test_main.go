package tgwrap

import (
	"fmt"
	"os"
)

func getTokenEnv() (string, error) {

	token := os.Getenv("TGWRAP_TOKEN")
	if len(token) < 1 {
		return "", fmt.Errorf("Cannot find token env:TGWRAP_TOKEN")
	}

	return token, nil
}
