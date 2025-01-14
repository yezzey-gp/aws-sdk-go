// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package keyspaces

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Amazon Keyspaces could not complete the requested action. This error may
	// occur if you try to perform an action and the same or a different action
	// is already in progress, or if you try to create a resource that already exists.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Amazon Keyspaces was unable to fully process this request because of an internal
	// server error.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The operation tried to access a keyspace or table that doesn't exist. The
	// resource might not be specified correctly, or its status might not be ACTIVE.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The operation exceeded the service quota for this resource. For more information
	// on service quotas, see Quotas (https://docs.aws.amazon.com/keyspaces/latest/devguide/quotas.html)
	// in the Amazon Keyspaces Developer Guide.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The operation failed due to an invalid or malformed request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ValidationException":           newErrorValidationException,
}
