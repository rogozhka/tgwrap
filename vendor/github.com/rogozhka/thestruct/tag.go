package thestruct

//
// Tag represents single entity
// inside struct tag literal
//
type Tag struct {
	//
	// Section contains "json"
	// in `json:name,omitempty`
	//
	Section string

	//
	// Value contains "name"
	// in `json:name,omitempty`
	//
	Value string

	//
	// Options contains ["omitempty","another-one"]
	// in `json:name,omitempty,another-one`
	//
	Options []string

	mapOpt map[string]bool
}

//
// IsOpt returns true if option is present
// inside Options slice
//
func (p *Tag) IsOpt(name string) bool {
	if _, is := p.mapOpt[name]; is {
		return true
	}

	return false
}
