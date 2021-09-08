// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/wasm/v3/wasm.proto

package envoy_extensions_wasm_v3

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

// Validate checks the field values on CapabilityRestrictionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CapabilityRestrictionConfig) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetAllowedCapabilities() {
		_ = val

		// no validation rules for AllowedCapabilities[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CapabilityRestrictionConfigValidationError{
					field:  fmt.Sprintf("AllowedCapabilities[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CapabilityRestrictionConfigValidationError is the validation error returned
// by CapabilityRestrictionConfig.Validate if the designated constraints
// aren't met.
type CapabilityRestrictionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapabilityRestrictionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapabilityRestrictionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapabilityRestrictionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapabilityRestrictionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapabilityRestrictionConfigValidationError) ErrorName() string {
	return "CapabilityRestrictionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CapabilityRestrictionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapabilityRestrictionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapabilityRestrictionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapabilityRestrictionConfigValidationError{}

// Validate checks the field values on SanitizationConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SanitizationConfig) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SanitizationConfigValidationError is the validation error returned by
// SanitizationConfig.Validate if the designated constraints aren't met.
type SanitizationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SanitizationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SanitizationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SanitizationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SanitizationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SanitizationConfigValidationError) ErrorName() string {
	return "SanitizationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e SanitizationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSanitizationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SanitizationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SanitizationConfigValidationError{}

// Validate checks the field values on VmConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *VmConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VmId

	if utf8.RuneCountInString(m.GetRuntime()) < 1 {
		return VmConfigValidationError{
			field:  "Runtime",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "Code",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "Configuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AllowPrecompiled

	// no validation rules for NackOnCodeCacheMiss

	if v, ok := interface{}(m.GetEnvironmentVariables()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "EnvironmentVariables",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// VmConfigValidationError is the validation error returned by
// VmConfig.Validate if the designated constraints aren't met.
type VmConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VmConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VmConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VmConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VmConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VmConfigValidationError) ErrorName() string { return "VmConfigValidationError" }

// Error satisfies the builtin error interface
func (e VmConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVmConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VmConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VmConfigValidationError{}

// Validate checks the field values on EnvironmentVariables with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EnvironmentVariables) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for KeyValues

	return nil
}

// EnvironmentVariablesValidationError is the validation error returned by
// EnvironmentVariables.Validate if the designated constraints aren't met.
type EnvironmentVariablesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvironmentVariablesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvironmentVariablesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvironmentVariablesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvironmentVariablesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvironmentVariablesValidationError) ErrorName() string {
	return "EnvironmentVariablesValidationError"
}

// Error satisfies the builtin error interface
func (e EnvironmentVariablesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvironmentVariables.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvironmentVariablesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvironmentVariablesValidationError{}

// Validate checks the field values on PluginConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PluginConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for RootId

	if v, ok := interface{}(m.GetConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PluginConfigValidationError{
				field:  "Configuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailOpen

	if v, ok := interface{}(m.GetCapabilityRestrictionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PluginConfigValidationError{
				field:  "CapabilityRestrictionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Vm.(type) {

	case *PluginConfig_VmConfig:

		if v, ok := interface{}(m.GetVmConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PluginConfigValidationError{
					field:  "VmConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PluginConfigValidationError is the validation error returned by
// PluginConfig.Validate if the designated constraints aren't met.
type PluginConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PluginConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PluginConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PluginConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PluginConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PluginConfigValidationError) ErrorName() string { return "PluginConfigValidationError" }

// Error satisfies the builtin error interface
func (e PluginConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPluginConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PluginConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PluginConfigValidationError{}

// Validate checks the field values on WasmService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WasmService) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WasmServiceValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Singleton

	return nil
}

// WasmServiceValidationError is the validation error returned by
// WasmService.Validate if the designated constraints aren't met.
type WasmServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WasmServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WasmServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WasmServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WasmServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WasmServiceValidationError) ErrorName() string { return "WasmServiceValidationError" }

// Error satisfies the builtin error interface
func (e WasmServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWasmService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WasmServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WasmServiceValidationError{}
