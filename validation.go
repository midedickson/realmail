package realmail

// This represents a Bus that moves to different validation destinations
// to collect the results at each destination
type ValidationBus struct {
	Success                              bool
	Email, Domain, ValidationDestination string
	Route                                []string
	Configuration                        *Configuration
	Errors                               map[string]string
}

// This is the main validation trip that decides the level
// of validation needed at every destination
type ValidationTrip struct {
	bus               *ValidationBus
	destination       string
	routeDependencies []string
}

// Validation engine interface that has varying abilitis
// to power the level of validation that can be done
// based on the given validation destination
type validationEngine interface {
	run(*ValidationBus) *ValidationBus
}

// function to create a new validation trip
func newValidationTrip(email, destination string, config *Configuration) *ValidationTrip {
	vt := &ValidationTrip{
		bus: &ValidationBus{
			Email:                 email,
			ValidationDestination: config.FinalDestination,
			Configuration:         config,
		},
		destination: destination,
	}
	vt.buildRouteDependencies()
	return vt

}

// ************** validation trip methods **************

// validationTrop method to build all the dependencies for the validation destinations
func (vt *ValidationTrip) buildRouteDependencies() {
	dependencies, ok := buildRouteDependency()[vt.destination]
	if ok {
		vt.routeDependencies = dependencies
	}
}

// Function to start the trip, which will kick start the bus
func (vt *ValidationTrip) start() *ValidationBus {
	// vt.attachEngine().bus.drive()
	validationBus := vt.bus
	if vt.routeDependencies == nil {
		getEngineByDestination(validationBus.Configuration.FinalDestination).run(validationBus)
	} else {
		runDependentValidationEngines(vt.routeDependencies, vt.bus)
	}
	return validationBus
}

// ************ end validation trip methods ************

// ************** validation bus methods **************
func (vb *ValidationBus) addError(key, value string) {
	if vb.Errors == nil {
		vb.Errors = map[string]string{}
	}
	vb.Errors[key] = value
}

// *********** end validation bus methods ************

// Function to check all validation dependencies
func runDependentValidationEngines(dependencies []string, vb *ValidationBus) {
	if !vb.Success {
		return
	}
	allDependencies := buildRouteDependency()
	for _, d := range dependencies {
		if len(allDependencies[d]) == 0 {
			getEngineByDestination(d).run(vb)
		} else {
			runDependentValidationEngines(allDependencies[d], vb)
		}
	}
}
