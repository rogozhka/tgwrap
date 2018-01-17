package tgwrap

//
// InputFile is used to store file_id, URL,
// or localFileName to encode as multipart/form-data
//
type InputFile struct {
	url    string
	fileID string

	fileNameLocal string
}

//
// Name returns stored localFileName or empty string
// used as an indicator that InputFile
// needs special treatment. In other cases standard
// Marshalling procedures are applied. No need to extract
// URL or file_id so there are no URL() and FileID() methods.
//
func (p *InputFile) Name() string {
	return p.fileNameLocal
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
// NewInputFileLocal is used to create InputFile
// with localFileName stored inside
//
func NewInputFileLocal(path string) *InputFile {

	o := &InputFile{}

	return o.SetFileName(path)
}

func (p *InputFile) reset() *InputFile {
	p.url = ""
	p.fileID = ""
	p.fileNameLocal = ""

	return p
}

//
// Switch InputFile into URL keeping mode
//
func (p *InputFile) SetURL(url string) *InputFile {
	p.reset()
	p.url = url
	return p
}

//
// Switch InputFile into file_id keeping mode
//
func (p *InputFile) SetFileID(fileID string) *InputFile {
	p.reset()
	p.fileID = fileID
	return p
}

//
// Switch InputFile into localFileName keeping mode
//
func (p *InputFile) SetFileName(fileName string) *InputFile {
	p.reset()
	p.fileNameLocal = fileName
	return p
}

//
// MarshalText() implements TextMarshaler
// and returns fileID or URL
//
// Note: fileName is ignored and should be used
// by multipart-form encoder
//
func (p *InputFile) MarshalText() (text []byte, err error) {

	if len(p.fileID) > 0 {
		return []byte(p.fileID), nil
	} else if len(p.url) > 0 {
		return []byte(p.url), nil
	}

	if len(p.fileNameLocal) < 1 {
		return []byte(""), nil
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
