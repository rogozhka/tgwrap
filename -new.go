package tgwrap

type bot struct {
	token string
}

func NewBot(token string) IBot {
	return &bot{
		token: token,
	}
}
