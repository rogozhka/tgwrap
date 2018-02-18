package tgwrap

type bot struct {
	token string
}

//
// NewBot creates new object with associated token inside
//
// note: object is object, new is new
// and this is stupid comment to make linter happy
//
func NewBot(token string) IBot {
	return createBot(token)
}

//
// createBot is used for testing purposes
// to test new methods still unexposed in IBot
//
func createBot(token string) *bot {
	return &bot{
		token: token,
	}
}
