package realmail

import (
	"fmt"
	"regexp"
)

// Returns true if the given string is present in slice,
// otherwise returns false.
func isIncluded(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}

// Returns slice of available validation destinations
func availableValidationDestination() []string {
	return []string{regexDestination}
}

// Validates validation type by available values,
// returns error if validation fails
func validateValidationDestinationContext(validationDestination string) error {
	if isIncluded(availableValidationDestination(), validationDestination) {
		return nil
	}

	return fmt.Errorf(
		"%s is invalid validation destination, use one of these: %s",
		validationDestination,
		availableValidationDestination(),
	)
}

// Regex Builder: Simply builds the regeex
func newRegex(regexPattern string) (*regexp.Regexp, error) {
	return regexp.Compile(regexPattern)
}

// Function to compare target strings with regex patterns
func matchRegex(target, regexPattern string) bool {
	regexp, err := newRegex(regexPattern)
	if err != nil {
		return false
	}
	return regexp.MatchString(target)
}

// Helps to confirm and set validation destination from input or sets to default if empty
func setValidationDestinationFromInput(options []string, defaultDestination string) (string, error) {
	if len(options) == 0 {
		return defaultDestination, nil
	}

	validationDestination := options[0]
	return validationDestination, validateValidationDestinationContext(validationDestination)
}

func getEngineByDestination(destination string) validationEngine {
	switch destination {
	case regexDestination:
		return &regexEngine{}
	}

	return &regexEngine{}
}
