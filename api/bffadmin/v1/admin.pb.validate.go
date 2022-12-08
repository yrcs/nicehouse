// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/admin.proto

package v1

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
)

import (
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

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Role) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Role with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RoleMultiError, or nil if none found.
func (m *Role) ValidateAll() error {
	return m.validate(true)
}

func (m *Role) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IsSystem

	if len(errors) > 0 {
		return RoleMultiError(errors)
	}

	return nil
}

// RoleMultiError is an error wrapping multiple validation errors returned by
// Role.ValidateAll() if the designated constraints aren't met.
type RoleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleMultiError) AllErrors() []error { return m }

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on GetRoleRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRoleRequestMultiError,
// or nil if none found.
func (m *GetRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetRoleRequestMultiError(errors)
	}

	return nil
}

// GetRoleRequestMultiError is an error wrapping multiple validation errors
// returned by GetRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleRequestMultiError) AllErrors() []error { return m }

// GetRoleRequestValidationError is the validation error returned by
// GetRoleRequest.Validate if the designated constraints aren't met.
type GetRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleRequestValidationError) ErrorName() string { return "GetRoleRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleRequestValidationError{}

// Validate checks the field values on CreateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRoleRequestMultiError, or nil if none found.
func (m *CreateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return CreateRoleRequestMultiError(errors)
	}

	return nil
}

// CreateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRoleRequestMultiError) AllErrors() []error { return m }

// CreateRoleRequestValidationError is the validation error returned by
// CreateRoleRequest.Validate if the designated constraints aren't met.
type CreateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoleRequestValidationError) ErrorName() string {
	return "CreateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoleRequestValidationError{}

// Validate checks the field values on UpdateRoleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleRequestMultiError, or nil if none found.
func (m *UpdateRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if len(errors) > 0 {
		return UpdateRoleRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRoleRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleRequestMultiError) AllErrors() []error { return m }

// UpdateRoleRequestValidationError is the validation error returned by
// UpdateRoleRequest.Validate if the designated constraints aren't met.
type UpdateRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleRequestValidationError) ErrorName() string {
	return "UpdateRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleRequestValidationError{}