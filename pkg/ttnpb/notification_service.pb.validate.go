// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _notification_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on Notification with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Notification) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = NotificationFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id":
			// no validation rules for Id
		case "created_at":

			if v, ok := interface{}(m.GetCreatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NotificationValidationError{
						field:  "created_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "entity_ids":

			if v, ok := interface{}(m.GetEntityIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NotificationValidationError{
						field:  "entity_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "notification_type":
			// no validation rules for NotificationType
		case "data":

			if v, ok := interface{}(m.GetData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NotificationValidationError{
						field:  "data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "sender_ids":

			if v, ok := interface{}(m.GetSenderIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NotificationValidationError{
						field:  "sender_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "receivers":

		case "email":
			// no validation rules for Email
		case "status":
			// no validation rules for Status
		case "status_updated_at":

			if v, ok := interface{}(m.GetStatusUpdatedAt()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NotificationValidationError{
						field:  "status_updated_at",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return NotificationValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// NotificationValidationError is the validation error returned by
// Notification.ValidateFields if the designated constraints aren't met.
type NotificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationValidationError) ErrorName() string { return "NotificationValidationError" }

// Error satisfies the builtin error interface
func (e NotificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationValidationError{}

// ValidateFields checks the field values on CreateNotificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateNotificationRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CreateNotificationRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "entity_ids":

			if m.GetEntityIds() == nil {
				return CreateNotificationRequestValidationError{
					field:  "entity_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetEntityIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CreateNotificationRequestValidationError{
						field:  "entity_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "notification_type":

			if l := utf8.RuneCountInString(m.GetNotificationType()); l < 1 || l > 100 {
				return CreateNotificationRequestValidationError{
					field:  "notification_type",
					reason: "value length must be between 1 and 100 runes, inclusive",
				}
			}

		case "data":

			if v, ok := interface{}(m.GetData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CreateNotificationRequestValidationError{
						field:  "data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "sender_ids":

			if v, ok := interface{}(m.GetSenderIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CreateNotificationRequestValidationError{
						field:  "sender_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "receivers":

			if len(m.GetReceivers()) < 1 {
				return CreateNotificationRequestValidationError{
					field:  "receivers",
					reason: "value must contain at least 1 item(s)",
				}
			}

			_CreateNotificationRequest_Receivers_Unique := make(map[NotificationReceiver]struct{}, len(m.GetReceivers()))

			for idx, item := range m.GetReceivers() {
				_, _ = idx, item

				if _, exists := _CreateNotificationRequest_Receivers_Unique[item]; exists {
					return CreateNotificationRequestValidationError{
						field:  fmt.Sprintf("receivers[%v]", idx),
						reason: "repeated value must contain unique items",
					}
				} else {
					_CreateNotificationRequest_Receivers_Unique[item] = struct{}{}
				}

				if _, ok := NotificationReceiver_name[int32(item)]; !ok {
					return CreateNotificationRequestValidationError{
						field:  fmt.Sprintf("receivers[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		case "email":
			// no validation rules for Email
		default:
			return CreateNotificationRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CreateNotificationRequestValidationError is the validation error returned by
// CreateNotificationRequest.ValidateFields if the designated constraints
// aren't met.
type CreateNotificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNotificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNotificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNotificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNotificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNotificationRequestValidationError) ErrorName() string {
	return "CreateNotificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNotificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNotificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNotificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNotificationRequestValidationError{}

// ValidateFields checks the field values on CreateNotificationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CreateNotificationResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CreateNotificationResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "id":
			// no validation rules for Id
		default:
			return CreateNotificationResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CreateNotificationResponseValidationError is the validation error returned
// by CreateNotificationResponse.ValidateFields if the designated constraints
// aren't met.
type CreateNotificationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateNotificationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateNotificationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateNotificationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateNotificationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateNotificationResponseValidationError) ErrorName() string {
	return "CreateNotificationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateNotificationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateNotificationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateNotificationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateNotificationResponseValidationError{}

// ValidateFields checks the field values on ListNotificationsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotificationsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListNotificationsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "receiver_ids":

			if m.GetReceiverIds() == nil {
				return ListNotificationsRequestValidationError{
					field:  "receiver_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetReceiverIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ListNotificationsRequestValidationError{
						field:  "receiver_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "status":

			_ListNotificationsRequest_Status_Unique := make(map[NotificationStatus]struct{}, len(m.GetStatus()))

			for idx, item := range m.GetStatus() {
				_, _ = idx, item

				if _, exists := _ListNotificationsRequest_Status_Unique[item]; exists {
					return ListNotificationsRequestValidationError{
						field:  fmt.Sprintf("status[%v]", idx),
						reason: "repeated value must contain unique items",
					}
				} else {
					_ListNotificationsRequest_Status_Unique[item] = struct{}{}
				}

				if _, ok := NotificationStatus_name[int32(item)]; !ok {
					return ListNotificationsRequestValidationError{
						field:  fmt.Sprintf("status[%v]", idx),
						reason: "value must be one of the defined enum values",
					}
				}

			}

		case "limit":

			if m.GetLimit() > 1000 {
				return ListNotificationsRequestValidationError{
					field:  "limit",
					reason: "value must be less than or equal to 1000",
				}
			}

		case "page":
			// no validation rules for Page
		default:
			return ListNotificationsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListNotificationsRequestValidationError is the validation error returned by
// ListNotificationsRequest.ValidateFields if the designated constraints
// aren't met.
type ListNotificationsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotificationsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotificationsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotificationsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotificationsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotificationsRequestValidationError) ErrorName() string {
	return "ListNotificationsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotificationsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotificationsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotificationsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotificationsRequestValidationError{}

// ValidateFields checks the field values on ListNotificationsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListNotificationsResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ListNotificationsResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "notifications":

			for idx, item := range m.GetNotifications() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ListNotificationsResponseValidationError{
							field:  fmt.Sprintf("notifications[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return ListNotificationsResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ListNotificationsResponseValidationError is the validation error returned by
// ListNotificationsResponse.ValidateFields if the designated constraints
// aren't met.
type ListNotificationsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListNotificationsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListNotificationsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListNotificationsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListNotificationsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListNotificationsResponseValidationError) ErrorName() string {
	return "ListNotificationsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListNotificationsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListNotificationsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListNotificationsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListNotificationsResponseValidationError{}

// ValidateFields checks the field values on UpdateNotificationStatusRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *UpdateNotificationStatusRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = UpdateNotificationStatusRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "receiver_ids":

			if m.GetReceiverIds() == nil {
				return UpdateNotificationStatusRequestValidationError{
					field:  "receiver_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetReceiverIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return UpdateNotificationStatusRequestValidationError{
						field:  "receiver_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "ids":

			if l := len(m.GetIds()); l < 1 || l > 1000 {
				return UpdateNotificationStatusRequestValidationError{
					field:  "ids",
					reason: "value must contain between 1 and 1000 items, inclusive",
				}
			}

			_UpdateNotificationStatusRequest_Ids_Unique := make(map[string]struct{}, len(m.GetIds()))

			for idx, item := range m.GetIds() {
				_, _ = idx, item

				if _, exists := _UpdateNotificationStatusRequest_Ids_Unique[item]; exists {
					return UpdateNotificationStatusRequestValidationError{
						field:  fmt.Sprintf("ids[%v]", idx),
						reason: "repeated value must contain unique items",
					}
				} else {
					_UpdateNotificationStatusRequest_Ids_Unique[item] = struct{}{}
				}

				if utf8.RuneCountInString(item) != 36 {
					return UpdateNotificationStatusRequestValidationError{
						field:  fmt.Sprintf("ids[%v]", idx),
						reason: "value length must be 36 runes",
					}

				}

			}

		case "status":

			if _, ok := NotificationStatus_name[int32(m.GetStatus())]; !ok {
				return UpdateNotificationStatusRequestValidationError{
					field:  "status",
					reason: "value must be one of the defined enum values",
				}
			}

		default:
			return UpdateNotificationStatusRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// UpdateNotificationStatusRequestValidationError is the validation error
// returned by UpdateNotificationStatusRequest.ValidateFields if the
// designated constraints aren't met.
type UpdateNotificationStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateNotificationStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateNotificationStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateNotificationStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateNotificationStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateNotificationStatusRequestValidationError) ErrorName() string {
	return "UpdateNotificationStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateNotificationStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateNotificationStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateNotificationStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateNotificationStatusRequestValidationError{}
