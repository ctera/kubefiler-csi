package driver

import (
	"fmt"
)

func ValidateDriverOptions(options *DriverOptions) error {
	if err := validateMode(options.mode); err != nil {
		return fmt.Errorf("invalid mode: %v", err)
	}

	return nil
}

func validateMode(mode Mode) error {
	if mode != AllMode && mode != ControllerMode && mode != NodeMode {
		return fmt.Errorf("mode is not supported (actual: %s, supported: %v)", mode, []Mode{AllMode, ControllerMode, NodeMode})
	}

	return nil
}
