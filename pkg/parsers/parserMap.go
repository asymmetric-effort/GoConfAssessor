package parsers

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
