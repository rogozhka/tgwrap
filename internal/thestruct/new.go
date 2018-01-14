package thestruct

func ParseLiteral(s string) (*Tags, error) {

	o := Tags{}

	arr, err := explodeStructTag(s)
	if err != nil {
		return nil, err
	}

	if err := o.index(arr); err != nil {
		return nil, err
	}

	return &o, nil
}
