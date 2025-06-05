package parsers

type NotImplemented struct{}

func (parser NotImplemented) Parse() error {
	return nil
}
