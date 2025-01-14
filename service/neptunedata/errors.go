// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptunedata

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Raised in case of an authentication or authorization failure.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Raised when a request is submitted that cannot be processed.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeBulkLoadIdNotFoundException for service response error code
	// "BulkLoadIdNotFoundException".
	//
	// Raised when a specified bulk-load job ID cannot be found.
	ErrCodeBulkLoadIdNotFoundException = "BulkLoadIdNotFoundException"

	// ErrCodeCancelledByUserException for service response error code
	// "CancelledByUserException".
	//
	// Raised when a user cancelled a request.
	ErrCodeCancelledByUserException = "CancelledByUserException"

	// ErrCodeClientTimeoutException for service response error code
	// "ClientTimeoutException".
	//
	// Raised when a request timed out in the client.
	ErrCodeClientTimeoutException = "ClientTimeoutException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Raised when a request attempts to modify data that is concurrently being
	// modified by another process.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConstraintViolationException for service response error code
	// "ConstraintViolationException".
	//
	// Raised when a value in a request field did not satisfy required constraints.
	ErrCodeConstraintViolationException = "ConstraintViolationException"

	// ErrCodeExpiredStreamException for service response error code
	// "ExpiredStreamException".
	//
	// Raised when a request attempts to access an stream that has expired.
	ErrCodeExpiredStreamException = "ExpiredStreamException"

	// ErrCodeFailureByQueryException for service response error code
	// "FailureByQueryException".
	//
	// Raised when a request fails.
	ErrCodeFailureByQueryException = "FailureByQueryException"

	// ErrCodeIllegalArgumentException for service response error code
	// "IllegalArgumentException".
	//
	// Raised when an argument in a request is not supported.
	ErrCodeIllegalArgumentException = "IllegalArgumentException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// Raised when the processing of the request failed unexpectedly.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// Raised when an argument in a request has an invalid value.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidNumericDataException for service response error code
	// "InvalidNumericDataException".
	//
	// Raised when invalid numerical data is encountered when servicing a request.
	ErrCodeInvalidNumericDataException = "InvalidNumericDataException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// Raised when a parameter value is not valid.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLoadUrlAccessDeniedException for service response error code
	// "LoadUrlAccessDeniedException".
	//
	// Raised when access is denied to a specified load URL.
	ErrCodeLoadUrlAccessDeniedException = "LoadUrlAccessDeniedException"

	// ErrCodeMLResourceNotFoundException for service response error code
	// "MLResourceNotFoundException".
	//
	// Raised when a specified machine-learning resource could not be found.
	ErrCodeMLResourceNotFoundException = "MLResourceNotFoundException"

	// ErrCodeMalformedQueryException for service response error code
	// "MalformedQueryException".
	//
	// Raised when a query is submitted that is syntactically incorrect or does
	// not pass additional validation.
	ErrCodeMalformedQueryException = "MalformedQueryException"

	// ErrCodeMemoryLimitExceededException for service response error code
	// "MemoryLimitExceededException".
	//
	// Raised when a request fails because of insufficient memory resources. The
	// request can be retried.
	ErrCodeMemoryLimitExceededException = "MemoryLimitExceededException"

	// ErrCodeMethodNotAllowedException for service response error code
	// "MethodNotAllowedException".
	//
	// Raised when the HTTP method used by a request is not supported by the endpoint
	// being used.
	ErrCodeMethodNotAllowedException = "MethodNotAllowedException"

	// ErrCodeMissingParameterException for service response error code
	// "MissingParameterException".
	//
	// Raised when a required parameter is missing.
	ErrCodeMissingParameterException = "MissingParameterException"

	// ErrCodeParsingException for service response error code
	// "ParsingException".
	//
	// Raised when a parsing issue is encountered.
	ErrCodeParsingException = "ParsingException"

	// ErrCodePreconditionsFailedException for service response error code
	// "PreconditionsFailedException".
	//
	// Raised when a precondition for processing a request is not satisfied.
	ErrCodePreconditionsFailedException = "PreconditionsFailedException"

	// ErrCodeQueryLimitExceededException for service response error code
	// "QueryLimitExceededException".
	//
	// Raised when the number of active queries exceeds what the server can process.
	// The query in question can be retried when the system is less busy.
	ErrCodeQueryLimitExceededException = "QueryLimitExceededException"

	// ErrCodeQueryLimitException for service response error code
	// "QueryLimitException".
	//
	// Raised when the size of a query exceeds the system limit.
	ErrCodeQueryLimitException = "QueryLimitException"

	// ErrCodeQueryTooLargeException for service response error code
	// "QueryTooLargeException".
	//
	// Raised when the body of a query is too large.
	ErrCodeQueryTooLargeException = "QueryTooLargeException"

	// ErrCodeReadOnlyViolationException for service response error code
	// "ReadOnlyViolationException".
	//
	// Raised when a request attempts to write to a read-only resource.
	ErrCodeReadOnlyViolationException = "ReadOnlyViolationException"

	// ErrCodeS3Exception for service response error code
	// "S3Exception".
	//
	// Raised when there is a problem accessing Amazon S3.
	ErrCodeS3Exception = "S3Exception"

	// ErrCodeServerShutdownException for service response error code
	// "ServerShutdownException".
	//
	// Raised when the server shuts down while processing a request.
	ErrCodeServerShutdownException = "ServerShutdownException"

	// ErrCodeStatisticsNotAvailableException for service response error code
	// "StatisticsNotAvailableException".
	//
	// Raised when statistics needed to satisfy a request are not available.
	ErrCodeStatisticsNotAvailableException = "StatisticsNotAvailableException"

	// ErrCodeStreamRecordsNotFoundException for service response error code
	// "StreamRecordsNotFoundException".
	//
	// Raised when stream records requested by a query cannot be found.
	ErrCodeStreamRecordsNotFoundException = "StreamRecordsNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Raised when the rate of requests exceeds the maximum throughput. Requests
	// can be retried after encountering this exception.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeTimeLimitExceededException for service response error code
	// "TimeLimitExceededException".
	//
	// Raised when the an operation exceeds the time limit allowed for it.
	ErrCodeTimeLimitExceededException = "TimeLimitExceededException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Raised when the number of requests being processed exceeds the limit.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// Raised when a request attempts to initiate an operation that is not supported.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":           newErrorAccessDeniedException,
	"BadRequestException":             newErrorBadRequestException,
	"BulkLoadIdNotFoundException":     newErrorBulkLoadIdNotFoundException,
	"CancelledByUserException":        newErrorCancelledByUserException,
	"ClientTimeoutException":          newErrorClientTimeoutException,
	"ConcurrentModificationException": newErrorConcurrentModificationException,
	"ConstraintViolationException":    newErrorConstraintViolationException,
	"ExpiredStreamException":          newErrorExpiredStreamException,
	"FailureByQueryException":         newErrorFailureByQueryException,
	"IllegalArgumentException":        newErrorIllegalArgumentException,
	"InternalFailureException":        newErrorInternalFailureException,
	"InvalidArgumentException":        newErrorInvalidArgumentException,
	"InvalidNumericDataException":     newErrorInvalidNumericDataException,
	"InvalidParameterException":       newErrorInvalidParameterException,
	"LoadUrlAccessDeniedException":    newErrorLoadUrlAccessDeniedException,
	"MLResourceNotFoundException":     newErrorMLResourceNotFoundException,
	"MalformedQueryException":         newErrorMalformedQueryException,
	"MemoryLimitExceededException":    newErrorMemoryLimitExceededException,
	"MethodNotAllowedException":       newErrorMethodNotAllowedException,
	"MissingParameterException":       newErrorMissingParameterException,
	"ParsingException":                newErrorParsingException,
	"PreconditionsFailedException":    newErrorPreconditionsFailedException,
	"QueryLimitExceededException":     newErrorQueryLimitExceededException,
	"QueryLimitException":             newErrorQueryLimitException,
	"QueryTooLargeException":          newErrorQueryTooLargeException,
	"ReadOnlyViolationException":      newErrorReadOnlyViolationException,
	"S3Exception":                     newErrorS3Exception,
	"ServerShutdownException":         newErrorServerShutdownException,
	"StatisticsNotAvailableException": newErrorStatisticsNotAvailableException,
	"StreamRecordsNotFoundException":  newErrorStreamRecordsNotFoundException,
	"ThrottlingException":             newErrorThrottlingException,
	"TimeLimitExceededException":      newErrorTimeLimitExceededException,
	"TooManyRequestsException":        newErrorTooManyRequestsException,
	"UnsupportedOperationException":   newErrorUnsupportedOperationException,
}
