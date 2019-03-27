package tgwrap

import (
	"fmt"
	"net/http"
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

func createTestBotFromEnv() *bot {

	var url string

	v := os.Getenv("TGWRAP_API_URL")
	if len(v) < 1 {
		url = TelegramBotAPI
	} else {
		url = v
	}

	return NewBotWithClientAndURL(getTokenEnv(), &http.Client{
		Timeout: DefaultClientTimeout,
	}, url)
}
