// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeCodeSigningConfigNotFoundException for service response error code
	// "CodeSigningConfigNotFoundException".
	//
	// The specified code signing configuration does not exist.
	ErrCodeCodeSigningConfigNotFoundException = "CodeSigningConfigNotFoundException"

	// ErrCodeCodeStorageExceededException for service response error code
	// "CodeStorageExceededException".
	//
	// Your Amazon Web Services account has exceeded its maximum total code size.
	// For more information, see Lambda quotas (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html).
	ErrCodeCodeStorageExceededException = "CodeStorageExceededException"

	// ErrCodeCodeVerificationFailedException for service response error code
	// "CodeVerificationFailedException".
	//
	// The code signature failed one or more of the validation checks for signature
	// mismatch or expiry, and the code signing policy is set to ENFORCE. Lambda
	// blocks the deployment.
	ErrCodeCodeVerificationFailedException = "CodeVerificationFailedException"

	// ErrCodeEC2AccessDeniedException for service response error code
	// "EC2AccessDeniedException".
	//
	// Need additional permissions to configure VPC settings.
	ErrCodeEC2AccessDeniedException = "EC2AccessDeniedException"

	// ErrCodeEC2ThrottledException for service response error code
	// "EC2ThrottledException".
	//
	// Amazon EC2 throttled Lambda during Lambda function initialization using the
	// execution role provided for the function.
	ErrCodeEC2ThrottledException = "EC2ThrottledException"

	// ErrCodeEC2UnexpectedException for service response error code
	// "EC2UnexpectedException".
	//
	// Lambda received an unexpected Amazon EC2 client exception while setting up
	// for the Lambda function.
	ErrCodeEC2UnexpectedException = "EC2UnexpectedException"

	// ErrCodeEFSIOException for service response error code
	// "EFSIOException".
	//
	// An error occurred when reading from or writing to a connected file system.
	ErrCodeEFSIOException = "EFSIOException"

	// ErrCodeEFSMountConnectivityException for service response error code
	// "EFSMountConnectivityException".
	//
	// The Lambda function couldn't make a network connection to the configured
	// file system.
	ErrCodeEFSMountConnectivityException = "EFSMountConnectivityException"

	// ErrCodeEFSMountFailureException for service response error code
	// "EFSMountFailureException".
	//
	// The Lambda function couldn't mount the configured file system due to a permission
	// or configuration issue.
	ErrCodeEFSMountFailureException = "EFSMountFailureException"

	// ErrCodeEFSMountTimeoutException for service response error code
	// "EFSMountTimeoutException".
	//
	// The Lambda function made a network connection to the configured file system,
	// but the mount operation timed out.
	ErrCodeEFSMountTimeoutException = "EFSMountTimeoutException"

	// ErrCodeENILimitReachedException for service response error code
	// "ENILimitReachedException".
	//
	// Lambda couldn't create an elastic network interface in the VPC, specified
	// as part of Lambda function configuration, because the limit for network interfaces
	// has been reached. For more information, see Lambda quotas (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html).
	ErrCodeENILimitReachedException = "ENILimitReachedException"

	// ErrCodeInvalidCodeSignatureException for service response error code
	// "InvalidCodeSignatureException".
	//
	// The code signature failed the integrity check. If the integrity check fails,
	// then Lambda blocks deployment, even if the code signing policy is set to
	// WARN.
	ErrCodeInvalidCodeSignatureException = "InvalidCodeSignatureException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One of the parameters in the request is not valid.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRequestContentException for service response error code
	// "InvalidRequestContentException".
	//
	// The request body could not be parsed as JSON.
	ErrCodeInvalidRequestContentException = "InvalidRequestContentException"

	// ErrCodeInvalidRuntimeException for service response error code
	// "InvalidRuntimeException".
	//
	// The runtime or runtime version specified is not supported.
	ErrCodeInvalidRuntimeException = "InvalidRuntimeException"

	// ErrCodeInvalidSecurityGroupIDException for service response error code
	// "InvalidSecurityGroupIDException".
	//
	// The security group ID provided in the Lambda function VPC configuration is
	// not valid.
	ErrCodeInvalidSecurityGroupIDException = "InvalidSecurityGroupIDException"

	// ErrCodeInvalidSubnetIDException for service response error code
	// "InvalidSubnetIDException".
	//
	// The subnet ID provided in the Lambda function VPC configuration is not valid.
	ErrCodeInvalidSubnetIDException = "InvalidSubnetIDException"

	// ErrCodeInvalidZipFileException for service response error code
	// "InvalidZipFileException".
	//
	// Lambda could not unzip the deployment package.
	ErrCodeInvalidZipFileException = "InvalidZipFileException"

	// ErrCodeKMSAccessDeniedException for service response error code
	// "KMSAccessDeniedException".
	//
	// Lambda couldn't decrypt the environment variables because KMS access was
	// denied. Check the Lambda function's KMS permissions.
	ErrCodeKMSAccessDeniedException = "KMSAccessDeniedException"

	// ErrCodeKMSDisabledException for service response error code
	// "KMSDisabledException".
	//
	// Lambda couldn't decrypt the environment variables because the KMS key used
	// is disabled. Check the Lambda function's KMS key settings.
	ErrCodeKMSDisabledException = "KMSDisabledException"

	// ErrCodeKMSInvalidStateException for service response error code
	// "KMSInvalidStateException".
	//
	// Lambda couldn't decrypt the environment variables because the state of the
	// KMS key used is not valid for Decrypt. Check the function's KMS key settings.
	ErrCodeKMSInvalidStateException = "KMSInvalidStateException"

	// ErrCodeKMSNotFoundException for service response error code
	// "KMSNotFoundException".
	//
	// Lambda couldn't decrypt the environment variables because the KMS key was
	// not found. Check the function's KMS key settings.
	ErrCodeKMSNotFoundException = "KMSNotFoundException"

	// ErrCodePolicyLengthExceededException for service response error code
	// "PolicyLengthExceededException".
	//
	// The permissions policy for the resource is too large. For more information,
	// see Lambda quotas (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html).
	ErrCodePolicyLengthExceededException = "PolicyLengthExceededException"

	// ErrCodePreconditionFailedException for service response error code
	// "PreconditionFailedException".
	//
	// The RevisionId provided does not match the latest RevisionId for the Lambda
	// function or alias. Call the GetFunction or the GetAlias API operation to
	// retrieve the latest RevisionId for your resource.
	ErrCodePreconditionFailedException = "PreconditionFailedException"

	// ErrCodeProvisionedConcurrencyConfigNotFoundException for service response error code
	// "ProvisionedConcurrencyConfigNotFoundException".
	//
	// The specified configuration does not exist.
	ErrCodeProvisionedConcurrencyConfigNotFoundException = "ProvisionedConcurrencyConfigNotFoundException"

	// ErrCodeRecursiveInvocationException for service response error code
	// "RecursiveInvocationException".
	//
	// Lambda has detected your function being invoked in a recursive loop with
	// other Amazon Web Services resources and stopped your function's invocation.
	ErrCodeRecursiveInvocationException = "RecursiveInvocationException"

	// ErrCodeRequestTooLargeException for service response error code
	// "RequestTooLargeException".
	//
	// The request payload exceeded the Invoke request body JSON input quota. For
	// more information, see Lambda quotas (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html).
	ErrCodeRequestTooLargeException = "RequestTooLargeException"

	// ErrCodeResourceConflictException for service response error code
	// "ResourceConflictException".
	//
	// The resource already exists, or another operation is in progress.
	ErrCodeResourceConflictException = "ResourceConflictException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The operation conflicts with the resource's availability. For example, you
	// tried to update an event source mapping in the CREATING state, or you tried
	// to delete an event source mapping currently UPDATING.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource specified in the request does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceNotReadyException for service response error code
	// "ResourceNotReadyException".
	//
	// The function is inactive and its VPC connection is no longer available. Wait
	// for the VPC connection to reestablish and try again.
	ErrCodeResourceNotReadyException = "ResourceNotReadyException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// The Lambda service encountered an internal error.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeSnapStartException for service response error code
	// "SnapStartException".
	//
	// The afterRestore() runtime hook (https://docs.aws.amazon.com/lambda/latest/dg/snapstart-runtime-hooks.html)
	// encountered an error. For more information, check the Amazon CloudWatch logs.
	ErrCodeSnapStartException = "SnapStartException"

	// ErrCodeSnapStartNotReadyException for service response error code
	// "SnapStartNotReadyException".
	//
	// Lambda is initializing your function. You can invoke the function when the
	// function state (https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html)
	// becomes Active.
	ErrCodeSnapStartNotReadyException = "SnapStartNotReadyException"

	// ErrCodeSnapStartTimeoutException for service response error code
	// "SnapStartTimeoutException".
	//
	// Lambda couldn't restore the snapshot within the timeout limit.
	ErrCodeSnapStartTimeoutException = "SnapStartTimeoutException"

	// ErrCodeSubnetIPAddressLimitReachedException for service response error code
	// "SubnetIPAddressLimitReachedException".
	//
	// Lambda couldn't set up VPC access for the Lambda function because one or
	// more configured subnets has no available IP addresses.
	ErrCodeSubnetIPAddressLimitReachedException = "SubnetIPAddressLimitReachedException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// The request throughput limit was exceeded. For more information, see Lambda
	// quotas (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html#api-requests).
	ErrCodeTooManyRequestsException = "TooManyRequestsException"

	// ErrCodeUnsupportedMediaTypeException for service response error code
	// "UnsupportedMediaTypeException".
	//
	// The content type of the Invoke request body is not JSON.
	ErrCodeUnsupportedMediaTypeException = "UnsupportedMediaTypeException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"CodeSigningConfigNotFoundException":            newErrorCodeSigningConfigNotFoundException,
	"CodeStorageExceededException":                  newErrorCodeStorageExceededException,
	"CodeVerificationFailedException":               newErrorCodeVerificationFailedException,
	"EC2AccessDeniedException":                      newErrorEC2AccessDeniedException,
	"EC2ThrottledException":                         newErrorEC2ThrottledException,
	"EC2UnexpectedException":                        newErrorEC2UnexpectedException,
	"EFSIOException":                                newErrorEFSIOException,
	"EFSMountConnectivityException":                 newErrorEFSMountConnectivityException,
	"EFSMountFailureException":                      newErrorEFSMountFailureException,
	"EFSMountTimeoutException":                      newErrorEFSMountTimeoutException,
	"ENILimitReachedException":                      newErrorENILimitReachedException,
	"InvalidCodeSignatureException":                 newErrorInvalidCodeSignatureException,
	"InvalidParameterValueException":                newErrorInvalidParameterValueException,
	"InvalidRequestContentException":                newErrorInvalidRequestContentException,
	"InvalidRuntimeException":                       newErrorInvalidRuntimeException,
	"InvalidSecurityGroupIDException":               newErrorInvalidSecurityGroupIDException,
	"InvalidSubnetIDException":                      newErrorInvalidSubnetIDException,
	"InvalidZipFileException":                       newErrorInvalidZipFileException,
	"KMSAccessDeniedException":                      newErrorKMSAccessDeniedException,
	"KMSDisabledException":                          newErrorKMSDisabledException,
	"KMSInvalidStateException":                      newErrorKMSInvalidStateException,
	"KMSNotFoundException":                          newErrorKMSNotFoundException,
	"PolicyLengthExceededException":                 newErrorPolicyLengthExceededException,
	"PreconditionFailedException":                   newErrorPreconditionFailedException,
	"ProvisionedConcurrencyConfigNotFoundException": newErrorProvisionedConcurrencyConfigNotFoundException,
	"RecursiveInvocationException":                  newErrorRecursiveInvocationException,
	"RequestTooLargeException":                      newErrorRequestTooLargeException,
	"ResourceConflictException":                     newErrorResourceConflictException,
	"ResourceInUseException":                        newErrorResourceInUseException,
	"ResourceNotFoundException":                     newErrorResourceNotFoundException,
	"ResourceNotReadyException":                     newErrorResourceNotReadyException,
	"ServiceException":                              newErrorServiceException,
	"SnapStartException":                            newErrorSnapStartException,
	"SnapStartNotReadyException":                    newErrorSnapStartNotReadyException,
	"SnapStartTimeoutException":                     newErrorSnapStartTimeoutException,
	"SubnetIPAddressLimitReachedException":          newErrorSubnetIPAddressLimitReachedException,
	"TooManyRequestsException":                      newErrorTooManyRequestsException,
	"UnsupportedMediaTypeException":                 newErrorUnsupportedMediaTypeException,
}
