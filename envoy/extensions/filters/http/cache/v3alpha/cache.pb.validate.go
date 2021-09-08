// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/cache/v3alpha/cache.proto

package envoy_extensions_filters_http_cache_v3alpha

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

// Validate checks the field values on CacheConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CacheConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTypedConfig() == nil {
		return CacheConfigValidationError{
			field:  "TypedConfig",
			reason: "value is required",
		}
	}

	if a := m.GetTypedConfig(); a != nil {

	}

	for idx, item := range m.GetAllowedVaryHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CacheConfigValidationError{
					field:  fmt.Sprintf("AllowedVaryHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetKeyCreatorParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CacheConfigValidationError{
				field:  "KeyCreatorParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MaxBodyBytes

	return nil
}

// CacheConfigValidationError is the validation error returned by
// CacheConfig.Validate if the designated constraints aren't met.
type CacheConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CacheConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CacheConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CacheConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CacheConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CacheConfigValidationError) ErrorName() string { return "CacheConfigValidationError" }

// Error satisfies the builtin error interface
func (e CacheConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCacheConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CacheConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CacheConfigValidationError{}

// Validate checks the field values on CacheConfig_KeyCreatorParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CacheConfig_KeyCreatorParams) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ExcludeScheme

	// no validation rules for ExcludeHost

	for idx, item := range m.GetQueryParametersIncluded() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CacheConfig_KeyCreatorParamsValidationError{
					field:  fmt.Sprintf("QueryParametersIncluded[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetQueryParametersExcluded() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CacheConfig_KeyCreatorParamsValidationError{
					field:  fmt.Sprintf("QueryParametersExcluded[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CacheConfig_KeyCreatorParamsValidationError is the validation error returned
// by CacheConfig_KeyCreatorParams.Validate if the designated constraints
// aren't met.
type CacheConfig_KeyCreatorParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CacheConfig_KeyCreatorParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CacheConfig_KeyCreatorParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CacheConfig_KeyCreatorParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CacheConfig_KeyCreatorParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CacheConfig_KeyCreatorParamsValidationError) ErrorName() string {
	return "CacheConfig_KeyCreatorParamsValidationError"
}

// Error satisfies the builtin error interface
func (e CacheConfig_KeyCreatorParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCacheConfig_KeyCreatorParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CacheConfig_KeyCreatorParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CacheConfig_KeyCreatorParamsValidationError{}
