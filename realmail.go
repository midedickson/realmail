package realmail

// Entry Point to start the validation trip process
func StartValidation(email string, config *Configuration, destinations ...string) (*ValidationBus, error) {
	validationDestination, err := setValidationDestinationFromInput(destinations, config.FinalDestination)
	if err != nil {
		return nil, err
	}
	return newValidationTrip(email, validationDestination, config).start(), err
}

// shourtut to the success status of StartValidation function above
func IsReal(email string, config *Configuration, destinations ...string) bool {
	validationDestination, err := setValidationDestinationFromInput(destinations, config.FinalDestination)
	if err != nil {
		return false
	}
	return newValidationTrip(email, validationDestination, config).start().Success
}
