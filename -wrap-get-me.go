package tgwrap

//
// Telegram API method
//
func (p *bot) GetMe() (*User, error) {

	var resp struct {
		GenericResponse

		Result *User `json:"result"`
	}

	err := p.getResponse("getMe", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
