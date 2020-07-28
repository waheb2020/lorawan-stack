// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _billing_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on Billing with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Billing) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = BillingFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "provider":
			if m.Provider == nil {
				return BillingValidationError{
					field:  "provider",
					reason: "value is required",
				}
			}
			if len(subs) == 0 {
				subs = []string{
					"stripe",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "stripe":
					w, ok := m.Provider.(*Billing_Stripe_)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetStripe()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return BillingValidationError{
								field:  "stripe",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return BillingValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// BillingValidationError is the validation error returned by
// Billing.ValidateFields if the designated constraints aren't met.
type BillingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BillingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BillingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BillingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BillingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BillingValidationError) ErrorName() string { return "BillingValidationError" }

// Error satisfies the builtin error interface
func (e BillingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBilling.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BillingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BillingValidationError{}

// ValidateFields checks the field values on Billing_Stripe with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Billing_Stripe) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = Billing_StripeFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "customer_id":

			if utf8.RuneCountInString(m.GetCustomerID()) < 1 {
				return Billing_StripeValidationError{
					field:  "customer_id",
					reason: "value length must be at least 1 runes",
				}
			}

		case "plan_id":

			if utf8.RuneCountInString(m.GetPlanID()) < 1 {
				return Billing_StripeValidationError{
					field:  "plan_id",
					reason: "value length must be at least 1 runes",
				}
			}

		case "subscription_id":

			if utf8.RuneCountInString(m.GetSubscriptionID()) < 1 {
				return Billing_StripeValidationError{
					field:  "subscription_id",
					reason: "value length must be at least 1 runes",
				}
			}

		case "subscription_item_id":

			if utf8.RuneCountInString(m.GetSubscriptionItemID()) < 1 {
				return Billing_StripeValidationError{
					field:  "subscription_item_id",
					reason: "value length must be at least 1 runes",
				}
			}

		default:
			return Billing_StripeValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// Billing_StripeValidationError is the validation error returned by
// Billing_Stripe.ValidateFields if the designated constraints aren't met.
type Billing_StripeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Billing_StripeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Billing_StripeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Billing_StripeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Billing_StripeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Billing_StripeValidationError) ErrorName() string { return "Billing_StripeValidationError" }

// Error satisfies the builtin error interface
func (e Billing_StripeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBilling_Stripe.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Billing_StripeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Billing_StripeValidationError{}