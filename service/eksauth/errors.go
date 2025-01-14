// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eksauth

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have permissions to perform the requested operation. The IAM principal
	// making the request must have at least one IAM permissions policy attached
	// that grants the required permissions. For more information, see Access management
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html) in the IAM
	// User Guide.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeExpiredTokenException for service response error code
	// "ExpiredTokenException".
	//
	// The specified Kubernetes service account token is expired.
	ErrCodeExpiredTokenException = "ExpiredTokenException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// These errors are usually caused by a server-side issue.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// This exception is thrown if the request contains a semantic error. The precise
	// meaning will depend on the API, and will be documented in the error message.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidTokenException for service response error code
	// "InvalidTokenException".
	//
	// The specified Kubernetes service account token is invalid.
	ErrCodeInvalidTokenException = "InvalidTokenException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is unavailable. Back off and retry the operation.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied because your request rate is too high. Reduce the
	// frequency of requests.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":       newErrorAccessDeniedException,
	"ExpiredTokenException":       newErrorExpiredTokenException,
	"InternalServerException":     newErrorInternalServerException,
	"InvalidParameterException":   newErrorInvalidParameterException,
	"InvalidRequestException":     newErrorInvalidRequestException,
	"InvalidTokenException":       newErrorInvalidTokenException,
	"ResourceNotFoundException":   newErrorResourceNotFoundException,
	"ServiceUnavailableException": newErrorServiceUnavailableException,
	"ThrottlingException":         newErrorThrottlingException,
}
