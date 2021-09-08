// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/udp/udp_proxy/v2alpha/udp_proxy.proto

package envoy_config_filter_udp_udp_proxy_v2alpha

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

// Validate checks the field values on UdpProxyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UdpProxyConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return UdpProxyConfigValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpProxyConfigValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.RouteSpecifier.(type) {

	case *UdpProxyConfig_Cluster:

		if len(m.GetCluster()) < 1 {
			return UdpProxyConfigValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 bytes",
			}
		}

	default:
		return UdpProxyConfigValidationError{
			field:  "RouteSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// UdpProxyConfigValidationError is the validation error returned by
// UdpProxyConfig.Validate if the designated constraints aren't met.
type UdpProxyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfigValidationError) ErrorName() string { return "UdpProxyConfigValidationError" }

// Error satisfies the builtin error interface
func (e UdpProxyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfigValidationError{}
