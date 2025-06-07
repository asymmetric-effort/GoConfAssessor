package parsers

import "fmt"

type ParserMap map[string]Parser

var SupportedParsers = ParserMap{
	"cisco-ios": &CiscoParser{},
	"juniper":   &NotImplemented{},
	"f5":        &NotImplemented{},
	"arista":    &NotImplemented{},
	"yaml":      &NotImplemented{},
	"json":      &NotImplemented{},
	"ini":       &NotImplemented{},
	"text":      &NotImplemented{},
}

// New returns a Parser for the given parserType.
// It errors if parserType is nil or not one of SupportedParsers.
func New(parserType *string) (Parser, error) {
	if parserType == nil {
		return nil, fmt.Errorf("parser type is required")
	}
	key := *parserType
	p, ok := SupportedParsers[key]
	if !ok {
		return nil, fmt.Errorf("unsupported parser type: %q", key)
	}
	return p, nil
}
