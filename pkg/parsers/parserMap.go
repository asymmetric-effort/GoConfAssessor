package parsers

import (
	"fmt"
	"strings"
)

type ParserMap map[string]Parser

var SupportedParsers = ParserMap{
	"cisco-ios": CiscoParser{},
	"juniper":   NotImplemented{},
	"f5":        NotImplemented{},
	"arista":    NotImplemented{},
	"yaml":      NotImplemented{},
	"json":      NotImplemented{},
	"ini":       NotImplemented{},
	"text":      NotImplemented{},
}

func IsValidParser(p string) error {
	thisParser := strings.ToLower(strings.TrimSpace(p))
	if _, ok := SupportedParsers[thisParser]; !ok {
		return fmt.Errorf("unknown or unsupported parser (%s)", thisParser)
	}
	return nil
}
