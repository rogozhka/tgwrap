package tgwrap

func (p *bot) GetFile(fileID string) (*File, error) {

	type sendFormat struct {
		FileID string `json:"file_id"`
	}
	dataSend := sendFormat{
		FileID: fileID,
	}
	var resp struct {
		GenericResponse
		Result *File `json:"result"`
	}
	err := p.getAPIResponse(nil, "getFile", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
