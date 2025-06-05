package manifest

import "fmt"

func (descriptor *Assertion) validateSource() (err error) {
	if err = descriptor.validateSourcePath(); err != nil {
		return err
	}
	if err = descriptor.validateSourcePattern(); err != nil {
		return err
	}
	if descriptor.Source.Pattern != "" && descriptor.Source.Path != "" {
		return fmt.Errorf("Source.Path and Source.Pattern cannot both be used in an assertion.  "+
			"See assertion '%s'", descriptor.Label)
	}
	return nil
}

func (descriptor *Assertion) validateSourcePath() (err error) {
	// ToDo: validate paths
	return err
}

func (descriptor *Assertion) validateSourcePattern() (err error) {
	// ToDo: validate patterns
	return err
}
