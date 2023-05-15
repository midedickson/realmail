package realmail

// This represents a Bus that moves to different validation destinations
// to collect the results at each destination
type ValidationBus struct {
	Success                       bool
	Email, Domain, ValidationType string
	Route                         []string
	validationEngine
}

// This is the main validation trip that decides the level
// of validation needed at every destination
type ValidationTrip struct {
	bus         *ValidationBus
	destination string
}

// Validation engine interface that has varying abilitis
// to power the level of validation that can be done
// based on the given validation destination
type validationEngine interface {
	run(*ValidationBus) *ValidationBus
}

// function to create a new validation trip
func newValidationTrip(email, destination string) *ValidationTrip {
	vt := &ValidationTrip{
		bus: &ValidationBus{
			Email: email,
		},
		destination: destination,
	}
	return vt

}

// function to attach the right validation engine based on the given destination
func (vt *ValidationTrip) attachEngine() *ValidationTrip {
	switch vt.destination {
	case regexDestination:
		vt.bus.validationEngine = &regexEngine{}
	}
	return vt
}

// function to start the validation bus to drive across the different validation destination
func (vb *ValidationBus) drive() {
	vb.validationEngine.run(vb)
}

// function to start the trip, which will kick start the bus
func (vt *ValidationTrip) start() *ValidationBus {
	vt.attachEngine()
	vt.bus.drive()
	return vt.bus
}
