package thestruct

//
// Tags represents collection
// of Tag entities inside single struct tag literal
//
type Tags struct {
	names   []string
	tagsMap map[string]*Tag
}

//
// Sections is used to list available
// sections inside literal
//
func (p *Tags) Sections() []string {
	return p.names
}

//
// IsSection returns true if given name
// is present as section inside literal
//
// Example: IsSection("json") returns true
// for`json:name,omitempty`
//

func (p *Tags) IsSection(name string) bool {

	if _, is := p.tagsMap[name]; is {
		return true
	}

	return false
}

//
// Tag returns single Tag from collection
// by name as Section
//
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
