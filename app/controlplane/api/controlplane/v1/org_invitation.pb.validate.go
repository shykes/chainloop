// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: controlplane/v1/org_invitation.proto

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

// define the regex for a UUID once up-front
var _org_invitation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OrgInvitationServiceCreateRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrgInvitationServiceCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceCreateRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceCreateRequestMultiError, or nil if none found.
func (m *OrgInvitationServiceCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if err := m._validateEmail(m.GetReceiverEmail()); err != nil {
		err = OrgInvitationServiceCreateRequestValidationError{
			field:  "ReceiverEmail",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrgInvitationServiceCreateRequestMultiError(errors)
	}

	return nil
}

func (m *OrgInvitationServiceCreateRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *OrgInvitationServiceCreateRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// OrgInvitationServiceCreateRequestMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceCreateRequest.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceCreateRequestMultiError) AllErrors() []error { return m }

// OrgInvitationServiceCreateRequestValidationError is the validation error
// returned by OrgInvitationServiceCreateRequest.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceCreateRequestValidationError) ErrorName() string {
	return "OrgInvitationServiceCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceCreateRequestValidationError{}

// Validate checks the field values on OrgInvitationServiceCreateResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrgInvitationServiceCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceCreateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceCreateResponseMultiError, or nil if none found.
func (m *OrgInvitationServiceCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrgInvitationServiceCreateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrgInvitationServiceCreateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrgInvitationServiceCreateResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrgInvitationServiceCreateResponseMultiError(errors)
	}

	return nil
}

// OrgInvitationServiceCreateResponseMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceCreateResponseMultiError) AllErrors() []error { return m }

// OrgInvitationServiceCreateResponseValidationError is the validation error
// returned by OrgInvitationServiceCreateResponse.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceCreateResponseValidationError) ErrorName() string {
	return "OrgInvitationServiceCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceCreateResponseValidationError{}

// Validate checks the field values on OrgInvitationServiceRevokeRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrgInvitationServiceRevokeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceRevokeRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceRevokeRequestMultiError, or nil if none found.
func (m *OrgInvitationServiceRevokeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceRevokeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = OrgInvitationServiceRevokeRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrgInvitationServiceRevokeRequestMultiError(errors)
	}

	return nil
}

func (m *OrgInvitationServiceRevokeRequest) _validateUuid(uuid string) error {
	if matched := _org_invitation_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OrgInvitationServiceRevokeRequestMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceRevokeRequest.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceRevokeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceRevokeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceRevokeRequestMultiError) AllErrors() []error { return m }

// OrgInvitationServiceRevokeRequestValidationError is the validation error
// returned by OrgInvitationServiceRevokeRequest.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceRevokeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceRevokeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceRevokeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceRevokeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceRevokeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceRevokeRequestValidationError) ErrorName() string {
	return "OrgInvitationServiceRevokeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceRevokeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceRevokeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceRevokeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceRevokeRequestValidationError{}

// Validate checks the field values on OrgInvitationServiceRevokeResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrgInvitationServiceRevokeResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceRevokeResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceRevokeResponseMultiError, or nil if none found.
func (m *OrgInvitationServiceRevokeResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceRevokeResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OrgInvitationServiceRevokeResponseMultiError(errors)
	}

	return nil
}

// OrgInvitationServiceRevokeResponseMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceRevokeResponse.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceRevokeResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceRevokeResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceRevokeResponseMultiError) AllErrors() []error { return m }

// OrgInvitationServiceRevokeResponseValidationError is the validation error
// returned by OrgInvitationServiceRevokeResponse.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceRevokeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceRevokeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceRevokeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceRevokeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceRevokeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceRevokeResponseValidationError) ErrorName() string {
	return "OrgInvitationServiceRevokeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceRevokeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceRevokeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceRevokeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceRevokeResponseValidationError{}

// Validate checks the field values on OrgInvitationServiceListSentRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrgInvitationServiceListSentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceListSentRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceListSentRequestMultiError, or nil if none found.
func (m *OrgInvitationServiceListSentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceListSentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OrgInvitationServiceListSentRequestMultiError(errors)
	}

	return nil
}

// OrgInvitationServiceListSentRequestMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceListSentRequest.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceListSentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceListSentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceListSentRequestMultiError) AllErrors() []error { return m }

// OrgInvitationServiceListSentRequestValidationError is the validation error
// returned by OrgInvitationServiceListSentRequest.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceListSentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceListSentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceListSentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceListSentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceListSentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceListSentRequestValidationError) ErrorName() string {
	return "OrgInvitationServiceListSentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceListSentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceListSentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceListSentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceListSentRequestValidationError{}

// Validate checks the field values on OrgInvitationServiceListSentResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *OrgInvitationServiceListSentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationServiceListSentResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrgInvitationServiceListSentResponseMultiError, or nil if none found.
func (m *OrgInvitationServiceListSentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationServiceListSentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrgInvitationServiceListSentResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrgInvitationServiceListSentResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrgInvitationServiceListSentResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrgInvitationServiceListSentResponseMultiError(errors)
	}

	return nil
}

// OrgInvitationServiceListSentResponseMultiError is an error wrapping multiple
// validation errors returned by
// OrgInvitationServiceListSentResponse.ValidateAll() if the designated
// constraints aren't met.
type OrgInvitationServiceListSentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationServiceListSentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationServiceListSentResponseMultiError) AllErrors() []error { return m }

// OrgInvitationServiceListSentResponseValidationError is the validation error
// returned by OrgInvitationServiceListSentResponse.Validate if the designated
// constraints aren't met.
type OrgInvitationServiceListSentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationServiceListSentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationServiceListSentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationServiceListSentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationServiceListSentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationServiceListSentResponseValidationError) ErrorName() string {
	return "OrgInvitationServiceListSentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationServiceListSentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationServiceListSentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationServiceListSentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationServiceListSentResponseValidationError{}

// Validate checks the field values on OrgInvitationItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrgInvitationItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrgInvitationItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrgInvitationItemMultiError, or nil if none found.
func (m *OrgInvitationItem) ValidateAll() error {
	return m.validate(true)
}

func (m *OrgInvitationItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrgInvitationItemValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ReceiverEmail

	if all {
		switch v := interface{}(m.GetSender()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "Sender",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "Sender",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSender()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrgInvitationItemValidationError{
				field:  "Sender",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOrganization()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "Organization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrgInvitationItemValidationError{
					field:  "Organization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrganization()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrgInvitationItemValidationError{
				field:  "Organization",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	if len(errors) > 0 {
		return OrgInvitationItemMultiError(errors)
	}

	return nil
}

// OrgInvitationItemMultiError is an error wrapping multiple validation errors
// returned by OrgInvitationItem.ValidateAll() if the designated constraints
// aren't met.
type OrgInvitationItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrgInvitationItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrgInvitationItemMultiError) AllErrors() []error { return m }

// OrgInvitationItemValidationError is the validation error returned by
// OrgInvitationItem.Validate if the designated constraints aren't met.
type OrgInvitationItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrgInvitationItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrgInvitationItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrgInvitationItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrgInvitationItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrgInvitationItemValidationError) ErrorName() string {
	return "OrgInvitationItemValidationError"
}

// Error satisfies the builtin error interface
func (e OrgInvitationItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrgInvitationItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrgInvitationItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrgInvitationItemValidationError{}
