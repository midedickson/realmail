package realmail

// Entry Point to start the validation trip process
func StartValidation(email, destination string) (*ValidationBus, error) {
	err := validateValidationDestinationContext(destination)
	if err != nil {
		return nil, err
	}
	return newValidationTrip(email, destination).start(), nil
}
