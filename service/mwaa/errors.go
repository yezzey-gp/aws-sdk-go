// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mwaa

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Access to the Apache Airflow Web UI or CLI has been denied due to insufficient
	// permissions. To learn more, see Accessing an Amazon MWAA environment (https://docs.aws.amazon.com/mwaa/latest/userguide/access-policies.html).
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// InternalServerException: An internal error has occurred.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// ResourceNotFoundException: The resource is not available.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// ValidationException: The provided input is not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ValidationException":       newErrorValidationException,
}
