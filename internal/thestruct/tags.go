package thestruct

type Tags struct {
	names   []string
	tagsMap map[string]*Tag
}

func (p *Tags) Sections() []string {
	return p.names
}

func (p *Tags) IsSection(name string) bool {

	if _, is := p.tagsMap[name]; is {
		return true
	}

	return false
}

func (p *Tags) Tag(name string) *Tag {
	if t, is := p.tagsMap[name]; is {
		return t
	}
	return nil
}

func (p *Tags) index(arr []*Tag) error {

	if p.tagsMap == nil {
		p.tagsMap = make(map[string]*Tag)
	}

	for _, t := range arr {
		p.tagsMap[t.Section] = t
		p.names = append(p.names, t.Section)
	}

	return nil
}
