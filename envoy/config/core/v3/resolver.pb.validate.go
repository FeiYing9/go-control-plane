// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/resolver.proto

package envoy_config_core_v3

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

// Validate checks the field values on DnsResolverOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsResolverOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UseTcpForDnsLookups

	// no validation rules for NoDefaultSearchDomain

	return nil
}

// DnsResolverOptionsValidationError is the validation error returned by
// DnsResolverOptions.Validate if the designated constraints aren't met.
type DnsResolverOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsResolverOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsResolverOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsResolverOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsResolverOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsResolverOptionsValidationError) ErrorName() string {
	return "DnsResolverOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e DnsResolverOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsResolverOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsResolverOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsResolverOptionsValidationError{}

// Validate checks the field values on DnsResolutionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsResolutionConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetResolvers()) < 1 {
		return DnsResolutionConfigValidationError{
			field:  "Resolvers",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetResolvers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsResolutionConfigValidationError{
					field:  fmt.Sprintf("Resolvers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetDnsResolverOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsResolutionConfigValidationError{
				field:  "DnsResolverOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DnsResolutionConfigValidationError is the validation error returned by
// DnsResolutionConfig.Validate if the designated constraints aren't met.
type DnsResolutionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsResolutionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsResolutionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsResolutionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsResolutionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsResolutionConfigValidationError) ErrorName() string {
	return "DnsResolutionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DnsResolutionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsResolutionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsResolutionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsResolutionConfigValidationError{}
