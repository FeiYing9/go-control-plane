// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/runtime/v3/rtds.proto

package envoy_service_runtime_v3

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

// Validate checks the field values on RtdsDummy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RtdsDummy) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RtdsDummyValidationError is the validation error returned by
// RtdsDummy.Validate if the designated constraints aren't met.
type RtdsDummyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RtdsDummyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RtdsDummyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RtdsDummyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RtdsDummyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RtdsDummyValidationError) ErrorName() string { return "RtdsDummyValidationError" }

// Error satisfies the builtin error interface
func (e RtdsDummyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRtdsDummy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RtdsDummyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RtdsDummyValidationError{}

// Validate checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Runtime) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return RuntimeValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetLayer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeValidationError{
				field:  "Layer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RuntimeValidationError is the validation error returned by Runtime.Validate
// if the designated constraints aren't met.
type RuntimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeValidationError) ErrorName() string { return "RuntimeValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeValidationError{}
