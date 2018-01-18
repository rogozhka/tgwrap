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
	return &bot{
		token: token,
	}
}
