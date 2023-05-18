package realmail

type regexEngine struct{}

// run function for regexEngine as needed by the validationEngine interface
func (e *regexEngine) run(vb *ValidationBus) *ValidationBus {
	if !vb.Configuration.EmailPattern.MatchString(vb.Email) {
		vb.Success = false
		vb.addError(regexDestination, regexErrorText)
	}
	return vb
}
