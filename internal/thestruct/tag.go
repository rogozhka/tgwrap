package thestruct

type Tag struct {
	//
	// json:name,omitempty
	//
	Section string
	Value   string
	Options []string
	mapOpt  map[string]bool
}

func (p *Tag) IsOpt(name string) bool {
	if _, is := p.mapOpt[name]; is {
		return true
	}

	return false
}
