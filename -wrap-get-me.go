package tgwrap

//
// Telegram API method
//
func (p *bot) GetMe() (*User, error) {

	var resp struct {
		GenericResponse

		Result *User `json:"result"`
	}

	err := p.getAPIResponse("getMe", p.sendJSON, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
