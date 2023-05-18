package realmail

import (
	"context"
	"fmt"
	"regexp"
)

type ConfigurationSetting struct {
	ctx                                                                context.Context
	EmailPattern, ValidationDestination, VerifierEmail, VerifierDomain string
	RegexEmail                                                         *regexp.Regexp
}

// Assigns default values on empty compulsory fields
func (c *ConfigurationSetting) assignDefaultValues() {
	if c.EmailPattern == emptyString {
		c.EmailPattern = regexEmailPattern
	}
	if c.ValidationDestination == emptyString {
		c.ValidationDestination = defaultDestination
	}

}

// Validates all the configuration settings with their defaults
func (c *ConfigurationSetting) validate() error {
	err := c.validateVerifierEmail()
	if err != nil {
		return err
	}

	c.VerifierDomain, err = c.buildVerifierDomain(c.VerifierEmail, c.VerifierDomain)
	if err != nil {
		return err
	}

	err = c.validateValidationDestination()
	if err != nil {
		return err
	}
	c.RegexEmail, err = newRegex(c.EmailPattern)
	if err != nil {
		return err
	}

	return nil
}

// Validate the verifier email
func (c *ConfigurationSetting) validateVerifierEmail() error {
	if matchRegex(c.VerifierEmail, emailCharsSize) && matchRegex(c.VerifierEmail, regexEmailPattern) {
		return nil
	}
	return fmt.Errorf("%s is not a valid verifier email", c.VerifierEmail)
}

// Validates strings based on patterns and specific use cases
// func (c *ConfigurationSetting) validateStringContext(target, pattern, msg string) error {
// 	if matchRegex(target, pattern) {
// 		return nil
// 	}
// 	return fmt.Errorf("%s is not a valid %s", target, msg)
// }

func (c *ConfigurationSetting) validateValidationDestination() error {
	return validateValidationDestinationContext(c.ValidationDestination)
}

func (c *ConfigurationSetting) validateVerifierDomain(verifierDomain string) (string, error) {
	if matchRegex(verifierDomain, domainCharsSize) && matchRegex(verifierDomain, regexDomainPattern) {
		return verifierDomain, nil
	}
	return verifierDomain, fmt.Errorf("%s is not a valid verifier domain", verifierDomain)

}

func (c *ConfigurationSetting) buildVerifierDomain(verifierEmail, verifierDomain string) (string, error) {
	if verifierDomain == emptyString {
		regex, _ := newRegex(regexEmailPattern)
		domainCaptureGroup := 3
		return regex.FindStringSubmatch(verifierEmail)[domainCaptureGroup], nil
	}

	return c.validateVerifierDomain(verifierDomain)
}
