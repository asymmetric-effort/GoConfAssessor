package parsers

type Parser interface {
	// Parse - parse the config file
	Parse() error
	// Find - Given an object path, return the value
	Find(objectByPath string) (value string, err error)
	// Pattern - Return all matches for a given pattern
	Pattern(regex string) (value []string, err error)
}

type Record struct {
	// LineNo - line number where record is found
	LineNo int
	// ExactLine - raw string of the line where the match is found
	ExactLine string
	// Value -
	Value string
	Err   error
}
