package realmail

type ValidationBus struct {
	Success                       bool
	Email, Domain, ValidationType string
	Route                         []string
	validationEngine
}

type ValidationTrip struct {
	bus         *ValidationBus
	destination string
}

type validationEngine interface {
	run(*ValidationBus) *ValidationBus
}

func newValidationTrip(email, destination string) *ValidationTrip {
	vt := &ValidationTrip{
		bus: &ValidationBus{
			Email: email,
		},
		destination: destination,
	}
	return vt

}

func (vt *ValidationTrip) attachEngine() *ValidationTrip {
	switch vt.destination {
	case regexDestination:
		vt.bus.validationEngine = &regexEngine{}
	}
	return vt
}

func (vb *ValidationBus) drive() {
	vb.validationEngine.run(vb)
}

func (vt *ValidationTrip) start() *ValidationBus {
	vt.attachEngine()
	vt.bus.drive()
	return vt.bus
}
