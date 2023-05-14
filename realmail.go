package realmail

// Entry Point to
func StartValidation(email, destination string) (*ValidationBus, error) {
	err := validateValidationTypeContext(destination)
	if err != nil {
		return nil, err
	}
	return newValidationTrip(email, destination).start(), nil
}
