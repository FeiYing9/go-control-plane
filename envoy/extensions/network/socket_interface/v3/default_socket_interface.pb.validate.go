// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/network/socket_interface/v3/default_socket_interface.proto

package envoy_extensions_network_socket_interface_v3

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

// Validate checks the field values on DefaultSocketInterface with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DefaultSocketInterface) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DefaultSocketInterfaceValidationError is the validation error returned by
// DefaultSocketInterface.Validate if the designated constraints aren't met.
type DefaultSocketInterfaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DefaultSocketInterfaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DefaultSocketInterfaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DefaultSocketInterfaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DefaultSocketInterfaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DefaultSocketInterfaceValidationError) ErrorName() string {
	return "DefaultSocketInterfaceValidationError"
}

// Error satisfies the builtin error interface
func (e DefaultSocketInterfaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDefaultSocketInterface.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DefaultSocketInterfaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DefaultSocketInterfaceValidationError{}
