package tgwrap

//
// InputFile is used to store file_id, URL,
// or localFileName to encode as multipart/form-data
//
type InputFile struct {
	url    string
	fileID string
}

//
// InputFileLocal is used to store localFileName
// to encode as multipart/form-data
//
type InputFileLocal struct {
	pathLocal string
}

//
// NewInputFileFromURL is used to create InputFile
// with URL stored inside
//
func NewInputFileFromURL(url string) *InputFile {
	return &InputFile{
		url: url,
	}
}

//
// NewInputFileFromID is used to create InputFile
// with file_id stored inside
//
func NewInputFileFromID(fileID string) *InputFile {
	return &InputFile{
		fileID: fileID,
	}
}

//
// NewInputFileLocal is used to create InputFile
// with localFileName stored inside
//
func NewInputFileLocal(path string) *InputFileLocal {

	return &InputFileLocal{
		pathLocal: path,
	}
}

//
// MarshalText implements TextMarshaler
// and returns fileID or URL
//
func (p *InputFile) MarshalText() (text []byte, err error) {

	if len(p.fileID) > 0 {
		return []byte(p.fileID), nil
	} else if len(p.url) > 0 {
		return []byte(p.url), nil
	}

	return []byte(""), nil
}

//
// String implements stringer interface
//
func (p *InputFile) String() string {

	bytes, err := p.MarshalText()

	if err != nil {
		return ""
	}

	return string(bytes)
}

//
// MarshalText implements TextMarshaler
//
func (p *InputFileLocal) MarshalText() (text []byte, err error) {

	return []byte(p.pathLocal), nil
}

//
// String implements stringer interface
//
func (p *InputFileLocal) String() string {

	bytes, err := p.MarshalText()

	if err != nil {
		return ""
	}

	return string(bytes)
}
