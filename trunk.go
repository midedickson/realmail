package realmail

import "fmt"

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
func availableValidationTypes() []string {
	return []string{regexDestination}
}

// Validates validation type by available values,
// returns error if validation fails
func validateValidationTypeContext(validationDestination string) error {
	if isIncluded(availableValidationTypes(), validationDestination) {
		return nil
	}

	return fmt.Errorf(
		"%s is invalid validation destination, use one of these: %s",
		validationDestination,
		availableValidationTypes(),
	)
}