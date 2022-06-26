package tgwrap

import "context"

func (p *bot) GetFile(fileID string) (*File, error) {
	return p.GetFileContext(nil, fileID)
}

func (p *bot) GetFileContext(ctx context.Context, fileID string) (*File, error) {
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
	err := p.getAPIResponse(ctx, "getFile", p.sendJSON, dataSend, &resp)
	return resp.Result, err
}
