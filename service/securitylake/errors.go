// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securitylake

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action. Access denied errors
	// appear when Amazon Security Lake explicitly or implicitly denies an authorization
	// request. An explicit denial occurs when a policy contains a Deny statement
	// for the specific Amazon Web Services action. An implicit denial occurs when
	// there is no applicable Deny statement and also no applicable Allow statement.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The request is malformed or contains an error such as an invalid parameter
	// value or a missing required parameter.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Occurs when a conflict with a previous successful write is detected. This
	// generally occurs when the previous write did not have time to propagate to
	// the host serving the current request. A retry (with appropriate backoff logic)
	// is the recommended response to this exception.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Internal service exceptions are sometimes caused by transient issues. Before
	// you start troubleshooting, perform the operation again.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The limit on the number of requests per second was exceeded.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"BadRequestException":       newErrorBadRequestException,
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
}
