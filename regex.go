package realmail

type regexEngine struct{}

// run function for regexEngine as needed by the validationEngine interface
func (e *regexEngine) run(vb *ValidationBus) *ValidationBus {
	return vb
}
