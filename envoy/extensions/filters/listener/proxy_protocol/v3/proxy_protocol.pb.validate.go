// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/listener/proxy_protocol/v3/proxy_protocol.proto

package envoy_extensions_filters_listener_proxy_protocol_v3

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

// Validate checks the field values on ProxyProtocol with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProxyProtocol) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProxyProtocolValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProxyProtocolValidationError is the validation error returned by
// ProxyProtocol.Validate if the designated constraints aren't met.
type ProxyProtocolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocolValidationError) ErrorName() string { return "ProxyProtocolValidationError" }

// Error satisfies the builtin error interface
func (e ProxyProtocolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocolValidationError{}

// Validate checks the field values on ProxyProtocol_KeyValuePair with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProxyProtocol_KeyValuePair) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MetadataNamespace

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		return ProxyProtocol_KeyValuePairValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// ProxyProtocol_KeyValuePairValidationError is the validation error returned
// by ProxyProtocol_KeyValuePair.Validate if the designated constraints aren't met.
type ProxyProtocol_KeyValuePairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocol_KeyValuePairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocol_KeyValuePairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocol_KeyValuePairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocol_KeyValuePairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocol_KeyValuePairValidationError) ErrorName() string {
	return "ProxyProtocol_KeyValuePairValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocol_KeyValuePairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol_KeyValuePair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocol_KeyValuePairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocol_KeyValuePairValidationError{}

// Validate checks the field values on ProxyProtocol_Rule with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProxyProtocol_Rule) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTlvType() >= 256 {
		return ProxyProtocol_RuleValidationError{
			field:  "TlvType",
			reason: "value must be less than 256",
		}
	}

	if v, ok := interface{}(m.GetOnTlvPresent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProxyProtocol_RuleValidationError{
				field:  "OnTlvPresent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProxyProtocol_RuleValidationError is the validation error returned by
// ProxyProtocol_Rule.Validate if the designated constraints aren't met.
type ProxyProtocol_RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocol_RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocol_RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocol_RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocol_RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocol_RuleValidationError) ErrorName() string {
	return "ProxyProtocol_RuleValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocol_RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocol_Rule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocol_RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocol_RuleValidationError{}
