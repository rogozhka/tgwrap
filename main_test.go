package tgwrap

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getTokenEnv() string {
	return requireEnv("TGWRAP_TOKEN")
}

func requireEnv(name string) string {
	return testEnvStringValue(name)
}

func testEnvStringValue(name string) string {

	tr := strings.TrimSpace(name)

	if v := os.Getenv(tr); len(v) < 1 {
		panic(fmt.Errorf("env not set | %s", tr))
	} else {
		return v
	}
}

func testEnvIsPresent(name string) bool {
	return len(os.Getenv(strings.TrimSpace(name))) > 0
}

func testEnvIntValue(name string) int {
	return intFromString(testEnvStringValue(name))
}

func testEnvUintValue(name string) uint {
	return uint(intFromString(testEnvStringValue(name)))
}

func intFromString(s string) int {

	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return int(res)
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
