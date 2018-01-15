package tgwrap

type InputFile struct {
	url    string
	fileID string

	fileNameLocal string
}

func (p *InputFile) Name() string {
	return p.fileNameLocal
}

func NewInputFile() *InputFile {
	return &InputFile{}
}

func NewInputFileFromURL(url string) *InputFile {
	return &InputFile{
		url: url,
	}
}

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

func (p *InputFile) SetURL(url string) *InputFile {
	p.reset()
	p.url = url
	return p
}

func (p *InputFile) SetFileID(fileID string) *InputFile {
	p.reset()
	p.fileID = fileID
	return p
}

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

func (p *InputFile) String() string {

	bytes, err := p.MarshalText()

	if err != nil {
		return ""
	}

	return string(bytes)
}
