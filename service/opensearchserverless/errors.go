// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opensearchserverless

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// When creating a resource, thrown when a resource with the same name already
	// exists or is being created. When deleting a resource, thrown when the resource
	// is not in the ACTIVE or FAILED state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Thrown when an error internal to the service occurs while processing a request.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeOcuLimitExceededException for service response error code
	// "OcuLimitExceededException".
	//
	// Thrown when the collection you're attempting to create results in a number
	// of search or indexing OCUs that exceeds the account limit.
	ErrCodeOcuLimitExceededException = "OcuLimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Thrown when accessing or deleting a resource that does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// Thrown when you attempt to create more resources than the service allows
	// based on service quotas.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Thrown when the HTTP request contains invalid input or is missing required
	// input.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"OcuLimitExceededException":     newErrorOcuLimitExceededException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ValidationException":           newErrorValidationException,
}
