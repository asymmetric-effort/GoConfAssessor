package parsers

import "fmt"

type NotImplemented struct{}

func (parser *NotImplemented) Parse() error {
	return fmt.Errorf("not implemented")
}

func (parser *NotImplemented) Find(objectByPath string) (value string, err error) {
	return value, fmt.Errorf("not implemented")
}

// Pattern - Return all matches for a given pattern
func (parser *NotImplemented) Pattern(regex string) (value []string, err error) {
	return value, fmt.Errorf("not implemented")
}
