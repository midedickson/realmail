package realmail

const (
	// validation destination types
	regexDestination   = "regex"
	defaultDestination = regexDestination

	// regex patterns

	domainCharsSize              = `\A.{4,255}\z`
	emailCharsSize               = `\A.{6,255}\z`
	regexDomainPattern           = `(?i)[\p{L}0-9]+([\-.]{1}[\p{L}0-9]+)*\.\p{L}{2,63}`
	regexEmailPattern            = `(\A([\p{L}0-9]+[\W\w]*)@(` + regexDomainPattern + `)\z)`
	regexDomainFromEmail         = `\A.+@(.+)\z`
	regexSMTPErrorBodyPattern    = `(?i).*550{1}.*(user|account|customer|mailbox).*`
	regexPortNumber              = `(6553[0-5]|655[0-2]\d|65[0-4](\d){2}|6[0-4](\d){3}|[1-5](\d){4}|[1-9](\d){0,3})`
	regexIpAddress               = `((\d|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])\.){3}(\d|[1-9]\d|1\d{2}|2[0-4]\d|25[0-5])`
	regexIpAddressPattern        = `\A` + regexIpAddress + `\z`
	regexDNSServerAddressPattern = `\A` + regexIpAddress + `(:` + regexPortNumber + `)?\z`

	// shortcuts
	emptyString = ""

	// regex destination specific
	regexErrorText = "email does not match regex expression"
)
