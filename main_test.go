package tgwrap

import (
	"fmt"
	"os"
)

func getTokenEnv() string {
	return requireEnv("TGWRAP_TOKEN")
}

func requireEnv(name string) string {
	v := os.Getenv(name)
	if len(v) < 1 {
		panic(fmt.Sprintf("ERROR: env variable is not set:%s", name))
	}

	return v
}
