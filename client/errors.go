package client

import (
	"regexp"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/digitalocean/godo"
	"github.com/pkg/errors"
)

func ErrorClassifier(_ schema.ClientMeta, resourceName string, err error) diag.Diagnostics {
	return classifyError(err, diag.RESOLVING, diag.WithResourceName(resourceName))
}

func classifyError(err error, fallbackType diag.Type, opts ...diag.BaseErrorOption) diag.Diagnostics {

	var ae *godo.ErrorResponse

	if errors.As(err, &ae) {
		switch ae.Message {
		case "permission denied":
			return diag.Diagnostics{
				RedactError(diag.NewBaseError(err,
					diag.ACCESS,
					append(opts,
						diag.WithType(diag.ACCESS),
						diag.WithSeverity(diag.WARNING),
						diag.WithError(errors.New(ae.Message)),
					)...),
				),
			}
		case "API Rate limit exceeded.":
			return diag.Diagnostics{
				RedactError(diag.NewBaseError(err,
					diag.THROTTLE,
					append(opts,
						diag.WithType(diag.THROTTLE),
						diag.WithSeverity(diag.WARNING),
						diag.WithError(errors.New(ae.Message)),
					)...),
				),
			}
		}
	}

	// Take over from SDK and always return diagnostics, redacting PII
	if d, ok := err.(diag.Diagnostic); ok {
		return diag.Diagnostics{
			RedactError(diag.NewBaseError(d, d.Type(), opts...)),
		}
	}

	return diag.Diagnostics{
		RedactError(diag.NewBaseError(err, fallbackType, opts...)),
	}
}

// RedactError redacts a given diagnostic and returns a RedactedDiagnostic containing both original and redacted versions
func RedactError(e diag.Diagnostic) diag.Diagnostic {
	r := diag.NewBaseError(
		nil,
		e.Type(),
		diag.WithSeverity(e.Severity()),
		diag.WithResourceName(e.Description().Resource),
		diag.WithSummary("%s", removePII(e.Description().Summary)),
		diag.WithDetails("%s", removePII(e.Description().Detail)),
	)
	return diag.NewRedactedDiagnostic(e, r)
}

var (
	requestIdRegex = regexp.MustCompile(`\s([Rr]equest:)\"\s[A-Za-z0-9-]+\"`)
)

func removePII(msg string) string {
	msg = requestIdRegex.ReplaceAllString(msg, " ${1} xxxx")
	return msg
}

func IsErrorMessage(err error, message string) bool {
	var ae *godo.ErrorResponse
	if !errors.As(err, &ae) {
		return false
	}
	if message == ae.Message {
		return true
	}
	return false
}

func IsErrorRegex(err error, messageRegex *regexp.Regexp) bool {
	var ae *godo.ErrorResponse
	if !errors.As(err, &ae) {
		return false
	}
	if messageRegex.MatchString(ae.Message) {
		return true
	}
	return false
}
