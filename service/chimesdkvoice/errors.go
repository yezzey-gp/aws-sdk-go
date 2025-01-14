// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chimesdkvoice

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have the permissions needed to run this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input parameters don't match the service's restrictions.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Multiple instances of the same request were made simultaneously.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The client is permanently forbidden from making the request.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeGoneException for service response error code
	// "GoneException".
	//
	// Access to the target resource is no longer available at the origin server.
	// This condition is likely to be permanent.
	ErrCodeGoneException = "GoneException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The requested resource couldn't be found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeResourceLimitExceededException for service response error code
	// "ResourceLimitExceededException".
	//
	// The request exceeds the resource limit.
	ErrCodeResourceLimitExceededException = "ResourceLimitExceededException"

	// ErrCodeServiceFailureException for service response error code
	// "ServiceFailureException".
	//
	// The service encountered an unexpected error.
	ErrCodeServiceFailureException = "ServiceFailureException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is currently unavailable.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottledClientException for service response error code
	// "ThrottledClientException".
	//
	// The number of customer requests exceeds the request rate limit.
	ErrCodeThrottledClientException = "ThrottledClientException"

	// ErrCodeUnauthorizedClientException for service response error code
	// "UnauthorizedClientException".
	//
	// The client isn't authorized to request a resource.
	ErrCodeUnauthorizedClientException = "UnauthorizedClientException"

	// ErrCodeUnprocessableEntityException for service response error code
	// "UnprocessableEntityException".
	//
	// A well-formed request couldn't be followed due to semantic errors.
	ErrCodeUnprocessableEntityException = "UnprocessableEntityException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":          newErrorAccessDeniedException,
	"BadRequestException":            newErrorBadRequestException,
	"ConflictException":              newErrorConflictException,
	"ForbiddenException":             newErrorForbiddenException,
	"GoneException":                  newErrorGoneException,
	"NotFoundException":              newErrorNotFoundException,
	"ResourceLimitExceededException": newErrorResourceLimitExceededException,
	"ServiceFailureException":        newErrorServiceFailureException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
	"ThrottledClientException":       newErrorThrottledClientException,
	"UnauthorizedClientException":    newErrorUnauthorizedClientException,
	"UnprocessableEntityException":   newErrorUnprocessableEntityException,
}
