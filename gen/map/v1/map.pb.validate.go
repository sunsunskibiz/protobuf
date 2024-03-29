// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: map/v1/map.proto

package mapv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on IdWrapper with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IdWrapper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IdWrapper with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdWrapperMultiError, or nil
// if none found.
func (m *IdWrapper) ValidateAll() error {
	return m.validate(true)
}

func (m *IdWrapper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdWrapperMultiError(errors)
	}

	return nil
}

// IdWrapperMultiError is an error wrapping multiple validation errors returned
// by IdWrapper.ValidateAll() if the designated constraints aren't met.
type IdWrapperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdWrapperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdWrapperMultiError) AllErrors() []error { return m }

// IdWrapperValidationError is the validation error returned by
// IdWrapper.Validate if the designated constraints aren't met.
type IdWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdWrapperValidationError) ErrorName() string { return "IdWrapperValidationError" }

// Error satisfies the builtin error interface
func (e IdWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdWrapperValidationError{}

// Validate checks the field values on MapExample with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MapExample) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MapExample with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MapExampleMultiError, or
// nil if none found.
func (m *MapExample) ValidateAll() error {
	return m.validate(true)
}

func (m *MapExample) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetIds()))
		i := 0
		for key := range m.GetIds() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetIds()[key]
			_ = val

			// no validation rules for Ids[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, MapExampleValidationError{
							field:  fmt.Sprintf("Ids[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, MapExampleValidationError{
							field:  fmt.Sprintf("Ids[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return MapExampleValidationError{
						field:  fmt.Sprintf("Ids[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return MapExampleMultiError(errors)
	}

	return nil
}

// MapExampleMultiError is an error wrapping multiple validation errors
// returned by MapExample.ValidateAll() if the designated constraints aren't met.
type MapExampleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MapExampleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MapExampleMultiError) AllErrors() []error { return m }

// MapExampleValidationError is the validation error returned by
// MapExample.Validate if the designated constraints aren't met.
type MapExampleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MapExampleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MapExampleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MapExampleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MapExampleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MapExampleValidationError) ErrorName() string { return "MapExampleValidationError" }

// Error satisfies the builtin error interface
func (e MapExampleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMapExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MapExampleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MapExampleValidationError{}
