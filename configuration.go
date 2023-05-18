package realmail

import (
	"context"
	"regexp"
)

// will have configuration parameters for each instance of the realmail invocation
type Configuration struct {
	ctx                                             context.Context
	FinalDestination, VerifierEmail, VerifierDomain string
	EmailPattern                                    *regexp.Regexp
}

func NewConfiguration(configSetting ConfigurationSetting) (*Configuration, error) {
	configSetting.assignDefaultValues()

	err := configSetting.validate()
	if err != nil {
		return nil, err
	}
	return &Configuration{
		ctx:              configSetting.ctx,
		FinalDestination: configSetting.ValidationDestination,
		VerifierEmail:    configSetting.VerifierEmail,
		VerifierDomain:   configSetting.VerifierDomain,
		EmailPattern:     configSetting.RegexEmail,
	}, nil
}
