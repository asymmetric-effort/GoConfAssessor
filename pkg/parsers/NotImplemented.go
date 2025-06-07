package parsers

type NotImplemented struct{}

func (parser *NotImplemented) Parse() error {
	return nil
}

func (parser *NotImplemented) Find(objectByPath string) (value string, err error) {
	return value, err
}

// Pattern - Return all matches for a given pattern
func (parser *NotImplemented) Pattern(regex string) (value []string, err error) {
	return value, err
}
