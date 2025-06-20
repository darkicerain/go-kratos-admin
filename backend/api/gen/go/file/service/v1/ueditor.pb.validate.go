// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: file/service/v1/ueditor.proto

package servicev1

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

// Validate checks the field values on UEditorRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UEditorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UEditorRequestMultiError,
// or nil if none found.
func (m *UEditorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Action

	// no validation rules for Encode

	// no validation rules for Start

	// no validation rules for Size

	if len(errors) > 0 {
		return UEditorRequestMultiError(errors)
	}

	return nil
}

// UEditorRequestMultiError is an error wrapping multiple validation errors
// returned by UEditorRequest.ValidateAll() if the designated constraints
// aren't met.
type UEditorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorRequestMultiError) AllErrors() []error { return m }

// UEditorRequestValidationError is the validation error returned by
// UEditorRequest.Validate if the designated constraints aren't met.
type UEditorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorRequestValidationError) ErrorName() string { return "UEditorRequestValidationError" }

// Error satisfies the builtin error interface
func (e UEditorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorRequestValidationError{}

// Validate checks the field values on UEditorResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UEditorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorResponseMultiError, or nil if none found.
func (m *UEditorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UEditorResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UEditorResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UEditorResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.ImageActionName != nil {
		// no validation rules for ImageActionName
	}

	if m.ImageFieldName != nil {
		// no validation rules for ImageFieldName
	}

	if m.ImageMaxSize != nil {
		// no validation rules for ImageMaxSize
	}

	if m.ImageCompressEnable != nil {
		// no validation rules for ImageCompressEnable
	}

	if m.ImageCompressBorder != nil {
		// no validation rules for ImageCompressBorder
	}

	if m.ImageInsertAlign != nil {
		// no validation rules for ImageInsertAlign
	}

	if m.ImageUrlPrefix != nil {
		// no validation rules for ImageUrlPrefix
	}

	if m.ImagePathFormat != nil {
		// no validation rules for ImagePathFormat
	}

	if m.ScrawlActionName != nil {
		// no validation rules for ScrawlActionName
	}

	if m.ScrawlFieldName != nil {
		// no validation rules for ScrawlFieldName
	}

	if m.ScrawlMaxSize != nil {
		// no validation rules for ScrawlMaxSize
	}

	if m.ScrawlUrlPrefix != nil {
		// no validation rules for ScrawlUrlPrefix
	}

	if m.ScrawlInsertAlign != nil {
		// no validation rules for ScrawlInsertAlign
	}

	if m.ScrawlPathFormat != nil {
		// no validation rules for ScrawlPathFormat
	}

	if m.SnapscreenActionName != nil {
		// no validation rules for SnapscreenActionName
	}

	if m.SnapscreenUrlPrefix != nil {
		// no validation rules for SnapscreenUrlPrefix
	}

	if m.SnapscreenInsertAlign != nil {
		// no validation rules for SnapscreenInsertAlign
	}

	if m.SnapscreenPathFormat != nil {
		// no validation rules for SnapscreenPathFormat
	}

	if m.CatcherActionName != nil {
		// no validation rules for CatcherActionName
	}

	if m.CatcherFieldName != nil {
		// no validation rules for CatcherFieldName
	}

	if m.CatcherUrlPrefix != nil {
		// no validation rules for CatcherUrlPrefix
	}

	if m.CatcherMaxSize != nil {
		// no validation rules for CatcherMaxSize
	}

	if m.CatcherPathFormat != nil {
		// no validation rules for CatcherPathFormat
	}

	if m.VideoActionName != nil {
		// no validation rules for VideoActionName
	}

	if m.VideoFieldName != nil {
		// no validation rules for VideoFieldName
	}

	if m.VideoUrlPrefix != nil {
		// no validation rules for VideoUrlPrefix
	}

	if m.VideoMaxSize != nil {
		// no validation rules for VideoMaxSize
	}

	if m.VideoPathFormat != nil {
		// no validation rules for VideoPathFormat
	}

	if m.FileActionName != nil {
		// no validation rules for FileActionName
	}

	if m.FileFieldName != nil {
		// no validation rules for FileFieldName
	}

	if m.FileUrlPrefix != nil {
		// no validation rules for FileUrlPrefix
	}

	if m.FileMaxSize != nil {
		// no validation rules for FileMaxSize
	}

	if m.FilePathFormat != nil {
		// no validation rules for FilePathFormat
	}

	if m.ImageManagerActionName != nil {
		// no validation rules for ImageManagerActionName
	}

	if m.ImageManagerListSize != nil {
		// no validation rules for ImageManagerListSize
	}

	if m.ImageManagerUrlPrefix != nil {
		// no validation rules for ImageManagerUrlPrefix
	}

	if m.ImageManagerInsertAlign != nil {
		// no validation rules for ImageManagerInsertAlign
	}

	if m.ImageManagerListPath != nil {
		// no validation rules for ImageManagerListPath
	}

	if m.FileManagerActionName != nil {
		// no validation rules for FileManagerActionName
	}

	if m.FileManagerUrlPrefix != nil {
		// no validation rules for FileManagerUrlPrefix
	}

	if m.FileManagerListSize != nil {
		// no validation rules for FileManagerListSize
	}

	if m.FileManagerListPath != nil {
		// no validation rules for FileManagerListPath
	}

	if m.FormulaConfig != nil {

		if all {
			switch v := interface{}(m.GetFormulaConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UEditorResponseValidationError{
						field:  "FormulaConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UEditorResponseValidationError{
						field:  "FormulaConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFormulaConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UEditorResponseValidationError{
					field:  "FormulaConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.State != nil {
		// no validation rules for State
	}

	if m.Start != nil {
		// no validation rules for Start
	}

	if m.Total != nil {
		// no validation rules for Total
	}

	if len(errors) > 0 {
		return UEditorResponseMultiError(errors)
	}

	return nil
}

// UEditorResponseMultiError is an error wrapping multiple validation errors
// returned by UEditorResponse.ValidateAll() if the designated constraints
// aren't met.
type UEditorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorResponseMultiError) AllErrors() []error { return m }

// UEditorResponseValidationError is the validation error returned by
// UEditorResponse.Validate if the designated constraints aren't met.
type UEditorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorResponseValidationError) ErrorName() string { return "UEditorResponseValidationError" }

// Error satisfies the builtin error interface
func (e UEditorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorResponseValidationError{}

// Validate checks the field values on UEditorUploadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorUploadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorUploadRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorUploadRequestMultiError, or nil if none found.
func (m *UEditorUploadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorUploadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Action != nil {
		// no validation rules for Action
	}

	if m.File != nil {
		// no validation rules for File
	}

	if m.SourceFileName != nil {
		// no validation rules for SourceFileName
	}

	if m.Mime != nil {
		// no validation rules for Mime
	}

	if len(errors) > 0 {
		return UEditorUploadRequestMultiError(errors)
	}

	return nil
}

// UEditorUploadRequestMultiError is an error wrapping multiple validation
// errors returned by UEditorUploadRequest.ValidateAll() if the designated
// constraints aren't met.
type UEditorUploadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorUploadRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorUploadRequestMultiError) AllErrors() []error { return m }

// UEditorUploadRequestValidationError is the validation error returned by
// UEditorUploadRequest.Validate if the designated constraints aren't met.
type UEditorUploadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorUploadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorUploadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorUploadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorUploadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorUploadRequestValidationError) ErrorName() string {
	return "UEditorUploadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorUploadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorUploadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorUploadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorUploadRequestValidationError{}

// Validate checks the field values on UEditorUploadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorUploadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorUploadResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorUploadResponseMultiError, or nil if none found.
func (m *UEditorUploadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorUploadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UEditorUploadResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UEditorUploadResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UEditorUploadResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.State != nil {
		// no validation rules for State
	}

	if m.Url != nil {
		// no validation rules for Url
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if m.Original != nil {
		// no validation rules for Original
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Size != nil {
		// no validation rules for Size
	}

	if len(errors) > 0 {
		return UEditorUploadResponseMultiError(errors)
	}

	return nil
}

// UEditorUploadResponseMultiError is an error wrapping multiple validation
// errors returned by UEditorUploadResponse.ValidateAll() if the designated
// constraints aren't met.
type UEditorUploadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorUploadResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorUploadResponseMultiError) AllErrors() []error { return m }

// UEditorUploadResponseValidationError is the validation error returned by
// UEditorUploadResponse.Validate if the designated constraints aren't met.
type UEditorUploadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorUploadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorUploadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorUploadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorUploadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorUploadResponseValidationError) ErrorName() string {
	return "UEditorUploadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorUploadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorUploadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorUploadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorUploadResponseValidationError{}

// Validate checks the field values on UEditorListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorListResponseMultiError, or nil if none found.
func (m *UEditorListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for State

	// no validation rules for Start

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UEditorListResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UEditorListResponseValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UEditorListResponseValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UEditorListResponseMultiError(errors)
	}

	return nil
}

// UEditorListResponseMultiError is an error wrapping multiple validation
// errors returned by UEditorListResponse.ValidateAll() if the designated
// constraints aren't met.
type UEditorListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorListResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorListResponseMultiError) AllErrors() []error { return m }

// UEditorListResponseValidationError is the validation error returned by
// UEditorListResponse.Validate if the designated constraints aren't met.
type UEditorListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorListResponseValidationError) ErrorName() string {
	return "UEditorListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorListResponseValidationError{}

// Validate checks the field values on UEditorResponse_FormulaConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorResponse_FormulaConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorResponse_FormulaConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UEditorResponse_FormulaConfigMultiError, or nil if none found.
func (m *UEditorResponse_FormulaConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorResponse_FormulaConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ImageUrlTemplate

	if len(errors) > 0 {
		return UEditorResponse_FormulaConfigMultiError(errors)
	}

	return nil
}

// UEditorResponse_FormulaConfigMultiError is an error wrapping multiple
// validation errors returned by UEditorResponse_FormulaConfig.ValidateAll()
// if the designated constraints aren't met.
type UEditorResponse_FormulaConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorResponse_FormulaConfigMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorResponse_FormulaConfigMultiError) AllErrors() []error { return m }

// UEditorResponse_FormulaConfigValidationError is the validation error
// returned by UEditorResponse_FormulaConfig.Validate if the designated
// constraints aren't met.
type UEditorResponse_FormulaConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorResponse_FormulaConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorResponse_FormulaConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorResponse_FormulaConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorResponse_FormulaConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorResponse_FormulaConfigValidationError) ErrorName() string {
	return "UEditorResponse_FormulaConfigValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorResponse_FormulaConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorResponse_FormulaConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorResponse_FormulaConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorResponse_FormulaConfigValidationError{}

// Validate checks the field values on UEditorResponse_Item with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorResponse_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorResponse_Item with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorResponse_ItemMultiError, or nil if none found.
func (m *UEditorResponse_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorResponse_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	// no validation rules for Mtime

	if len(errors) > 0 {
		return UEditorResponse_ItemMultiError(errors)
	}

	return nil
}

// UEditorResponse_ItemMultiError is an error wrapping multiple validation
// errors returned by UEditorResponse_Item.ValidateAll() if the designated
// constraints aren't met.
type UEditorResponse_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorResponse_ItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorResponse_ItemMultiError) AllErrors() []error { return m }

// UEditorResponse_ItemValidationError is the validation error returned by
// UEditorResponse_Item.Validate if the designated constraints aren't met.
type UEditorResponse_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorResponse_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorResponse_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorResponse_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorResponse_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorResponse_ItemValidationError) ErrorName() string {
	return "UEditorResponse_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorResponse_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorResponse_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorResponse_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorResponse_ItemValidationError{}

// Validate checks the field values on UEditorUploadResponse_Item with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorUploadResponse_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorUploadResponse_Item with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorUploadResponse_ItemMultiError, or nil if none found.
func (m *UEditorUploadResponse_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorUploadResponse_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for State

	if m.Url != nil {
		// no validation rules for Url
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if m.Original != nil {
		// no validation rules for Original
	}

	if m.Type != nil {
		// no validation rules for Type
	}

	if m.Size != nil {
		// no validation rules for Size
	}

	if len(errors) > 0 {
		return UEditorUploadResponse_ItemMultiError(errors)
	}

	return nil
}

// UEditorUploadResponse_ItemMultiError is an error wrapping multiple
// validation errors returned by UEditorUploadResponse_Item.ValidateAll() if
// the designated constraints aren't met.
type UEditorUploadResponse_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorUploadResponse_ItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorUploadResponse_ItemMultiError) AllErrors() []error { return m }

// UEditorUploadResponse_ItemValidationError is the validation error returned
// by UEditorUploadResponse_Item.Validate if the designated constraints aren't met.
type UEditorUploadResponse_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorUploadResponse_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorUploadResponse_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorUploadResponse_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorUploadResponse_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorUploadResponse_ItemValidationError) ErrorName() string {
	return "UEditorUploadResponse_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorUploadResponse_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorUploadResponse_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorUploadResponse_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorUploadResponse_ItemValidationError{}

// Validate checks the field values on UEditorListResponse_Item with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UEditorListResponse_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UEditorListResponse_Item with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UEditorListResponse_ItemMultiError, or nil if none found.
func (m *UEditorListResponse_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *UEditorListResponse_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	// no validation rules for Mtime

	if len(errors) > 0 {
		return UEditorListResponse_ItemMultiError(errors)
	}

	return nil
}

// UEditorListResponse_ItemMultiError is an error wrapping multiple validation
// errors returned by UEditorListResponse_Item.ValidateAll() if the designated
// constraints aren't met.
type UEditorListResponse_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UEditorListResponse_ItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UEditorListResponse_ItemMultiError) AllErrors() []error { return m }

// UEditorListResponse_ItemValidationError is the validation error returned by
// UEditorListResponse_Item.Validate if the designated constraints aren't met.
type UEditorListResponse_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UEditorListResponse_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UEditorListResponse_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UEditorListResponse_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UEditorListResponse_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UEditorListResponse_ItemValidationError) ErrorName() string {
	return "UEditorListResponse_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e UEditorListResponse_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUEditorListResponse_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UEditorListResponse_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UEditorListResponse_ItemValidationError{}
