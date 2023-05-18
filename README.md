# realmail - Only ~~valid~~ real emails

Realmail is a configurable 📨 email validator written in Golang.
You can verify email via Regex, DNS, SMTP and even more, hopefully :-)
Spend less on marketing, filter fake users by making sure emails are not only valid but real as well.

> Please Note: Work is still in progress, not taking contributions for now. Thanks

## Table of Contents

- [Synopsis](#synopsis)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
  - [Configuration features](#configuration-features)
    - [Creating configuration](#creating-configuration)
    - [Using configuration](#using-configuration)
  - [Validation features](#validation-features)
    - [Regex validation](#regex-validation)
      - [With default regex pattern](#with-default-regex-pattern)
      - [With custom regex pattern](#with-custom-regex-pattern)
- [Contributing](#contributing)
- [License](#license)
<!-- - [Code of Conduct](#code-of-conduct)
- [Credits](#credits)
- [Versioning](#versioning)
- [Changelog](CHANGELOG.md) -->

## Synopsis

Email validation is a hard enough. As a developer, I have encountered the problem of validating real emails in several business use cases and requirements. There are a number of different ways to validate an email address and they must conform with the best practices and provide proper validation. `realmail` package helps you validate emails via:

- Regex pattern
- Presence of DNS records
- Real existence of email account on a current email server.

**Syntax Checking**: Checks the email addresses via regex pattern.

**Mail Server Existence Check**: Checks the availability of the email address domain using DNS records.

**Mail Existence Check**: Checks if the email address really exists and can receive email via SMTP connections and email-sending emulation techniques.

## Features

- Configurable validator, validate only what you need
- Supporting of internationalized emails ([EAI](https://en.wikipedia.org/wiki/Email_address#Internationalization))
<!-- - Whitelist/blacklist validation layers (coming soon)
- Ability to configure different MX/SMTP validation flows (coming soon)
- Simple SMTP debugger (coming soon) -->

## Requirements

Golang 1.19+

## Installation

Install `realmail`:

```bash
go get github.com/Double-DOS/realmail
go install -i github.com/Double-DOS/realmail
```

Import `realmail` dependency into your code:

```go
package main

import "github.com/Double-DOS/realmail"
```

## Usage

### Configuration features

You can use global package configuration or custom independent configuration. Available configuration options (growing list):

- verifier email
- verifier domain
- email pattern
<!-- - connection timeout
- response timeout
- connection attempts
- default validation type
- validation type for domains
- whitelisted domains
- whitelist validation
- blacklisted domains
- blacklisted mx ip-addresses
- custom DNS gateway
- RFC MX lookup flow
- SMTP port number
- SMTP error body pattern
- SMTP fail fast
- SMTP safe check -->

#### Creating configuration

TO access all the library features, you must first create configuration struct first. Please use `realmail.NewConfiguration()` built-in constructor to create a valid configuration as in the example below:

```go
import "github.com/Double-DOS/realmail"

configuration := realmail.NewConfiguration(
  ConfigurationSetting{
    // Required parameter. Must be an existing email on behalf of which verification will be
    // performed
    VerifierEmail: "verifier@example.com",

    // Optional parameter. Must be an existing domain on behalf of which verification will be
    // performed. By default verifier domain based on verifier email
    VerifierDomain: "somedomain.com",

    // Optional parameter. You can override default regex pattern
    EmailPattern: `\A.+@(.+)\z`,


    // Optional parameter. You can predefine default validation type for
    // realmail.StartValidation("email@email.com", configuration) call without type-parameter
    // Available validation types: "regex", "mx", "mx_blacklist", "smtp"
    // note: only regex validation is currently supported as other validation destinations
    // are still in progress.
    ValidationDestination: "mx",
  },
)
```

#### Using configuration

```go
import "github.com/Double-DOS/realmail"

configuration := realmail.NewConfiguration(realmail.ConfigurationSetting{VerifierEmail: "verifier@example.com"})


// Returns a pointer to a validation bus that holds information about
// the status of the validation, its success or failure
// and validation errors where they occur.
realmail.StartValidation("some@email.com", configuration)

// Returns boolean status of the validation
realmail.IsReal("email@example.com", configuration, "regex")

```

#### Regex validation

<!-- Validation with regex pattern is the first validation level. It uses whitelist/blacklist check before running itself.

```code
[Whitelist/Blacklist] -> [Regex validation]
``` -->

By default this validation not performs strictly following [RFC 5322](https://www.ietf.org/rfc/rfc5322.txt) standard, so you can override `realmail` default regex pattern if you want.

Example of usage:

##### With default regex pattern

```go
import "github.com/Double-DOS/realmail"

configuration := realmail.NewConfiguration(
  realmail.ConfigurationSetting{
    VerifierEmail: "verifier@example.com",
  },
)

realmail.StartValidation("email@example.com", configuration, "regex") // returns pointer to ValidatorResult with validation details and error
```

##### With custom regex pattern

You should define your custom regex pattern in the configuration setting before calling the `StartValidation()` method

```go
import "github.com/Double-DOS/realmail"

configuration := realmail.NewConfiguration(
  realmail.ConfigurationSetting{
    VerifierEmail: "verifier@example.com",
    EmailPattern: `\A(.+)@(.+)\z`,
  },
)

realmail.StartValidation("email@example.com", configuration, "regex") // returns pointer to ValidatorResult with validation details and error
realmail.IsReal("email@example.com", configuration, "regex") // returns true
realmail.IsReal("not_email", configuration, "regex") // returns false
```

## Contributing

Although the first version of the project is still in progress, this project is intended to be a safe, welcoming space for collaboration and contribution.
