// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBillExpirationException for service response error code
	// "BillExpirationException".
	//
	// The requested report expired. Update the date interval and try again.
	ErrCodeBillExpirationException = "BillExpirationException"

	// ErrCodeDataUnavailableException for service response error code
	// "DataUnavailableException".
	//
	// The requested data is unavailable.
	ErrCodeDataUnavailableException = "DataUnavailableException"

	// ErrCodeGenerationExistsException for service response error code
	// "GenerationExistsException".
	//
	// A request to generate a recommendation is already in progress.
	ErrCodeGenerationExistsException = "GenerationExistsException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The pagination token is invalid. Try again without a pagination token.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// You made too many calls in a short period of time. Try again later.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeRequestChangedException for service response error code
	// "RequestChangedException".
	//
	// Your request parameters changed between pages. Try again with the old parameters
	// or without a pagination token.
	ErrCodeRequestChangedException = "RequestChangedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified ARN in the request doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You've reached the limit on the number of resources you can create, or exceeded
	// the size of an individual resource.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// Can occur if you specify a number of tags for a resource greater than the
	// maximum 50 user tags per resource.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeUnknownMonitorException for service response error code
	// "UnknownMonitorException".
	//
	// The cost anomaly monitor does not exist for the account.
	ErrCodeUnknownMonitorException = "UnknownMonitorException"

	// ErrCodeUnknownSubscriptionException for service response error code
	// "UnknownSubscriptionException".
	//
	// The cost anomaly subscription does not exist for the account.
	ErrCodeUnknownSubscriptionException = "UnknownSubscriptionException"

	// ErrCodeUnresolvableUsageUnitException for service response error code
	// "UnresolvableUsageUnitException".
	//
	// Cost Explorer was unable to identify the usage unit. Provide UsageType/UsageTypeGroup
	// filter selections that contain matching units, for example: hours.
	ErrCodeUnresolvableUsageUnitException = "UnresolvableUsageUnitException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BillExpirationException":        newErrorBillExpirationException,
	"DataUnavailableException":       newErrorDataUnavailableException,
	"GenerationExistsException":      newErrorGenerationExistsException,
	"InvalidNextTokenException":      newErrorInvalidNextTokenException,
	"LimitExceededException":         newErrorLimitExceededException,
	"RequestChangedException":        newErrorRequestChangedException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ServiceQuotaExceededException":  newErrorServiceQuotaExceededException,
	"TooManyTagsException":           newErrorTooManyTagsException,
	"UnknownMonitorException":        newErrorUnknownMonitorException,
	"UnknownSubscriptionException":   newErrorUnknownSubscriptionException,
	"UnresolvableUsageUnitException": newErrorUnresolvableUsageUnitException,
}
