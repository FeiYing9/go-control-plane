// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v2/api_listener.proto

package envoy_config_listener_v2

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on ApiListener with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ApiListener) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetApiListener()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ApiListenerValidationError{
				field:  "ApiListener",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ApiListenerValidationError is the validation error returned by
// ApiListener.Validate if the designated constraints aren't met.
type ApiListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApiListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApiListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApiListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApiListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApiListenerValidationError) ErrorName() string { return "ApiListenerValidationError" }

// Error satisfies the builtin error interface
func (e ApiListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApiListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApiListenerValidationError{}
