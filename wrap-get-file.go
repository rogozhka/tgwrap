package tgwrap

import "fmt"

// GetFile is used to get basic info about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size. On success, a File object is returned.
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
//
// fileID: (string) File identifier to get info about
//
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

	var sender fCommandSender = p.sendJSON

	err := p.getAPIResponse("getFile", sender, dataSend, &resp)
	if err != nil {
		return nil, fmt.Errorf("getAPIResponse ERROR:%v", err)
	}

	return resp.Result, nil
}
