// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package freetier

import (
	"fmt"

	"github.com/yezzey-gp/aws-sdk-go/aws"
	"github.com/yezzey-gp/aws-sdk-go/aws/awsutil"
	"github.com/yezzey-gp/aws-sdk-go/aws/request"
	"github.com/yezzey-gp/aws-sdk-go/private/protocol"
)

const opGetFreeTierUsage = "GetFreeTierUsage"

// GetFreeTierUsageRequest generates a "aws/request.Request" representing the
// client's request for the GetFreeTierUsage operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetFreeTierUsage for more information on using the GetFreeTierUsage
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the GetFreeTierUsageRequest method.
//	req, resp := client.GetFreeTierUsageRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/freetier-2023-09-07/GetFreeTierUsage
func (c *FreeTier) GetFreeTierUsageRequest(input *GetFreeTierUsageInput) (req *request.Request, output *GetFreeTierUsageOutput) {
	op := &request.Operation{
		Name:       opGetFreeTierUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetFreeTierUsageInput{}
	}

	output = &GetFreeTierUsageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetFreeTierUsage API operation for AWS Free Tier.
//
// Returns a list of all Free Tier usage objects that match your filters.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Free Tier's
// API operation GetFreeTierUsage for usage and error information.
//
// Returned Error Types:
//
//   - InternalServerException
//     An unexpected error occurred during the processing of your request.
//
//   - ValidationException
//     The input fails to satisfy the constraints specified by an Amazon Web Service.
//
//   - ThrottlingException
//     The request was denied due to request throttling.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/freetier-2023-09-07/GetFreeTierUsage
func (c *FreeTier) GetFreeTierUsage(input *GetFreeTierUsageInput) (*GetFreeTierUsageOutput, error) {
	req, out := c.GetFreeTierUsageRequest(input)
	return out, req.Send()
}

// GetFreeTierUsageWithContext is the same as GetFreeTierUsage with the addition of
// the ability to pass a context and additional request options.
//
// See GetFreeTierUsage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FreeTier) GetFreeTierUsageWithContext(ctx aws.Context, input *GetFreeTierUsageInput, opts ...request.Option) (*GetFreeTierUsageOutput, error) {
	req, out := c.GetFreeTierUsageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// GetFreeTierUsagePages iterates over the pages of a GetFreeTierUsage operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See GetFreeTierUsage method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//	// Example iterating over at most 3 pages of a GetFreeTierUsage operation.
//	pageNum := 0
//	err := client.GetFreeTierUsagePages(params,
//	    func(page *freetier.GetFreeTierUsageOutput, lastPage bool) bool {
//	        pageNum++
//	        fmt.Println(page)
//	        return pageNum <= 3
//	    })
func (c *FreeTier) GetFreeTierUsagePages(input *GetFreeTierUsageInput, fn func(*GetFreeTierUsageOutput, bool) bool) error {
	return c.GetFreeTierUsagePagesWithContext(aws.BackgroundContext(), input, fn)
}

// GetFreeTierUsagePagesWithContext same as GetFreeTierUsagePages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FreeTier) GetFreeTierUsagePagesWithContext(ctx aws.Context, input *GetFreeTierUsageInput, fn func(*GetFreeTierUsageOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *GetFreeTierUsageInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.GetFreeTierUsageRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*GetFreeTierUsageOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

// Contains the specifications for the filters to use for your request.
type DimensionValues struct {
	_ struct{} `type:"structure"`

	// The name of the dimension that you want to filter on.
	//
	// Key is a required field
	Key *string `type:"string" required:"true" enum:"Dimension"`

	// The match options that you can use to filter your results. You can specify
	// only one of these values in the array.
	//
	// MatchOptions is a required field
	MatchOptions []*string `type:"list" required:"true" enum:"MatchOption"`

	// The metadata values you can specify to filter upon, so that the results all
	// match at least one of the specified values.
	//
	// Values is a required field
	Values []*string `min:"1" type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DimensionValues) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s DimensionValues) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DimensionValues) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DimensionValues"}
	if s.Key == nil {
		invalidParams.Add(request.NewErrParamRequired("Key"))
	}
	if s.MatchOptions == nil {
		invalidParams.Add(request.NewErrParamRequired("MatchOptions"))
	}
	if s.Values == nil {
		invalidParams.Add(request.NewErrParamRequired("Values"))
	}
	if s.Values != nil && len(s.Values) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Values", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetKey sets the Key field's value.
func (s *DimensionValues) SetKey(v string) *DimensionValues {
	s.Key = &v
	return s
}

// SetMatchOptions sets the MatchOptions field's value.
func (s *DimensionValues) SetMatchOptions(v []*string) *DimensionValues {
	s.MatchOptions = v
	return s
}

// SetValues sets the Values field's value.
func (s *DimensionValues) SetValues(v []*string) *DimensionValues {
	s.Values = v
	return s
}

// Use Expression to filter in the GetFreeTierUsage API operation.
//
// You can use the following patterns:
//
//   - Simple dimension values (Dimensions root operator)
//
//   - Complex expressions with logical operators (AND, NOT, and OR root operators).
//
// For simple dimension values, you can set the dimension name, values, and
// match type for the filters that you plan to use.
//
// # Example for simple dimension values
//
// You can filter to match exactly for REGION==us-east-1 OR REGION==us-west-1.
//
// The corresponding Expression appears like the following: { "Dimensions":
// { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ], "MatchOptions":
// ["EQUALS"] } }
//
// As shown in the previous example, lists of dimension values are combined
// with OR when you apply the filter.
//
// For complex expressions with logical operators, you can have nested expressions
// to use the logical operators and specify advanced filtering.
//
// # Example for complex expressions with logical operators
//
// You can filter by ((REGION == us-east-1 OR REGION == us-west-1) OR (SERVICE
// CONTAINS AWSLambda)) AND (USAGE_TYPE !CONTAINS DataTransfer).
//
// The corresponding Expression appears like the following: { "And": [ {"Or":
// [ {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1"
// ], "MatchOptions": ["EQUALS"] }}, {"Dimensions": { "Key": "SERVICE", "Values":
// ["AWSLambda"], "MatchOptions": ["CONTAINS"] } } ]}, {"Not": {"Dimensions":
// { "Key": "USAGE_TYPE", "Values": ["DataTransfer"], "MatchOptions": ["CONTAINS"]
// }}} ] }
//
// In the following Contents, you must specify exactly one of the following
// root operators.
type Expression struct {
	_ struct{} `type:"structure"`

	// Return results that match all Expressions that you specified in the array.
	And []*Expression `type:"list"`

	// The specific dimension, values, and match type to filter objects with.
	Dimensions *DimensionValues `type:"structure"`

	// Return results that don’t match the Expression that you specified.
	Not *Expression `type:"structure"`

	// Return results that match any of the Expressions that you specified. in the
	// array.
	Or []*Expression `type:"list"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Expression) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Expression) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Expression) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Expression"}
	if s.Dimensions != nil {
		if err := s.Dimensions.Validate(); err != nil {
			invalidParams.AddNested("Dimensions", err.(request.ErrInvalidParams))
		}
	}
	if s.Not != nil {
		if err := s.Not.Validate(); err != nil {
			invalidParams.AddNested("Not", err.(request.ErrInvalidParams))
		}
	}
	if s.Or != nil {
		for i, v := range s.Or {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Or", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAnd sets the And field's value.
func (s *Expression) SetAnd(v []*Expression) *Expression {
	s.And = v
	return s
}

// SetDimensions sets the Dimensions field's value.
func (s *Expression) SetDimensions(v *DimensionValues) *Expression {
	s.Dimensions = v
	return s
}

// SetNot sets the Not field's value.
func (s *Expression) SetNot(v *Expression) *Expression {
	s.Not = v
	return s
}

// SetOr sets the Or field's value.
func (s *Expression) SetOr(v []*Expression) *Expression {
	s.Or = v
	return s
}

// Consists of a Amazon Web Services Free Tier offer’s metadata and your data
// usage for the offer.
type FreeTierUsage struct {
	_ struct{} `type:"structure"`

	// Describes the actual usage accrued month-to-day (MTD) that you've used so
	// far.
	ActualUsageAmount *float64 `locationName:"actualUsageAmount" type:"double"`

	// The description of the Free Tier offer.
	Description *string `locationName:"description" type:"string"`

	// Describes the forecasted usage by the month that you're expected to use.
	ForecastedUsageAmount *float64 `locationName:"forecastedUsageAmount" type:"double"`

	// Describes the type of the Free Tier offer. For example, the offer can be
	// "12 Months Free", "Always Free", and "Free Trial".
	FreeTierType *string `locationName:"freeTierType" type:"string"`

	// Describes the maximum usage allowed in Free Tier.
	Limit *float64 `locationName:"limit" type:"double"`

	// Describes usageType more granularly with the specific Amazon Web Service
	// API operation. For example, this can be the RunInstances API operation for
	// Amazon Elastic Compute Cloud.
	Operation *string `locationName:"operation" type:"string"`

	// Describes the Amazon Web Services Region for which this offer is applicable
	Region *string `locationName:"region" type:"string"`

	// The name of the Amazon Web Service providing the Free Tier offer. For example,
	// this can be Amazon Elastic Compute Cloud.
	Service *string `locationName:"service" type:"string"`

	// Describes the unit of the usageType, such as Hrs.
	Unit *string `locationName:"unit" type:"string"`

	// Describes the usage details of the offer. For example, this might be Global-BoxUsage:freetrial.
	UsageType *string `locationName:"usageType" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s FreeTierUsage) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s FreeTierUsage) GoString() string {
	return s.String()
}

// SetActualUsageAmount sets the ActualUsageAmount field's value.
func (s *FreeTierUsage) SetActualUsageAmount(v float64) *FreeTierUsage {
	s.ActualUsageAmount = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *FreeTierUsage) SetDescription(v string) *FreeTierUsage {
	s.Description = &v
	return s
}

// SetForecastedUsageAmount sets the ForecastedUsageAmount field's value.
func (s *FreeTierUsage) SetForecastedUsageAmount(v float64) *FreeTierUsage {
	s.ForecastedUsageAmount = &v
	return s
}

// SetFreeTierType sets the FreeTierType field's value.
func (s *FreeTierUsage) SetFreeTierType(v string) *FreeTierUsage {
	s.FreeTierType = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *FreeTierUsage) SetLimit(v float64) *FreeTierUsage {
	s.Limit = &v
	return s
}

// SetOperation sets the Operation field's value.
func (s *FreeTierUsage) SetOperation(v string) *FreeTierUsage {
	s.Operation = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *FreeTierUsage) SetRegion(v string) *FreeTierUsage {
	s.Region = &v
	return s
}

// SetService sets the Service field's value.
func (s *FreeTierUsage) SetService(v string) *FreeTierUsage {
	s.Service = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *FreeTierUsage) SetUnit(v string) *FreeTierUsage {
	s.Unit = &v
	return s
}

// SetUsageType sets the UsageType field's value.
func (s *FreeTierUsage) SetUsageType(v string) *FreeTierUsage {
	s.UsageType = &v
	return s
}

type GetFreeTierUsageInput struct {
	_ struct{} `type:"structure"`

	// An expression that specifies the conditions that you want each FreeTierUsage
	// object to meet.
	Filter *Expression `locationName:"filter" type:"structure"`

	// The maximum number of results to return in the response. MaxResults means
	// that there can be up to the specified number of values, but there might be
	// fewer results based on your filters.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetFreeTierUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetFreeTierUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFreeTierUsageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetFreeTierUsageInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *GetFreeTierUsageInput) SetFilter(v *Expression) *GetFreeTierUsageInput {
	s.Filter = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *GetFreeTierUsageInput) SetMaxResults(v int64) *GetFreeTierUsageInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetFreeTierUsageInput) SetNextToken(v string) *GetFreeTierUsageInput {
	s.NextToken = &v
	return s
}

type GetFreeTierUsageOutput struct {
	_ struct{} `type:"structure"`

	// The list of Free Tier usage objects that meet your filter expression.
	//
	// FreeTierUsages is a required field
	FreeTierUsages []*FreeTierUsage `locationName:"freeTierUsages" type:"list" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetFreeTierUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s GetFreeTierUsageOutput) GoString() string {
	return s.String()
}

// SetFreeTierUsages sets the FreeTierUsages field's value.
func (s *GetFreeTierUsageOutput) SetFreeTierUsages(v []*FreeTierUsage) *GetFreeTierUsageOutput {
	s.FreeTierUsages = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *GetFreeTierUsageOutput) SetNextToken(v string) *GetFreeTierUsageOutput {
	s.NextToken = &v
	return s
}

// An unexpected error occurred during the processing of your request.
type InternalServerException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServerException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InternalServerException) GoString() string {
	return s.String()
}

func newErrorInternalServerException(v protocol.ResponseMetadata) error {
	return &InternalServerException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InternalServerException) Code() string {
	return "InternalServerException"
}

// Message returns the exception's message.
func (s *InternalServerException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InternalServerException) OrigErr() error {
	return nil
}

func (s *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InternalServerException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InternalServerException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The request was denied due to request throttling.
type ThrottlingException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s *ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ThrottlingException) OrigErr() error {
	return nil
}

func (s *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ThrottlingException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ThrottlingException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The input fails to satisfy the constraints specified by an Amazon Web Service.
type ValidationException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ValidationException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ValidationException) GoString() string {
	return s.String()
}

func newErrorValidationException(v protocol.ResponseMetadata) error {
	return &ValidationException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ValidationException) Code() string {
	return "ValidationException"
}

// Message returns the exception's message.
func (s *ValidationException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ValidationException) OrigErr() error {
	return nil
}

func (s *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ValidationException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ValidationException) RequestID() string {
	return s.RespMetadata.RequestID
}

const (
	// DimensionService is a Dimension enum value
	DimensionService = "SERVICE"

	// DimensionOperation is a Dimension enum value
	DimensionOperation = "OPERATION"

	// DimensionUsageType is a Dimension enum value
	DimensionUsageType = "USAGE_TYPE"

	// DimensionRegion is a Dimension enum value
	DimensionRegion = "REGION"

	// DimensionFreeTierType is a Dimension enum value
	DimensionFreeTierType = "FREE_TIER_TYPE"

	// DimensionDescription is a Dimension enum value
	DimensionDescription = "DESCRIPTION"

	// DimensionUsagePercentage is a Dimension enum value
	DimensionUsagePercentage = "USAGE_PERCENTAGE"
)

// Dimension_Values returns all elements of the Dimension enum
func Dimension_Values() []string {
	return []string{
		DimensionService,
		DimensionOperation,
		DimensionUsageType,
		DimensionRegion,
		DimensionFreeTierType,
		DimensionDescription,
		DimensionUsagePercentage,
	}
}

const (
	// MatchOptionEquals is a MatchOption enum value
	MatchOptionEquals = "EQUALS"

	// MatchOptionStartsWith is a MatchOption enum value
	MatchOptionStartsWith = "STARTS_WITH"

	// MatchOptionEndsWith is a MatchOption enum value
	MatchOptionEndsWith = "ENDS_WITH"

	// MatchOptionContains is a MatchOption enum value
	MatchOptionContains = "CONTAINS"

	// MatchOptionGreaterThanOrEqual is a MatchOption enum value
	MatchOptionGreaterThanOrEqual = "GREATER_THAN_OR_EQUAL"
)

// MatchOption_Values returns all elements of the MatchOption enum
func MatchOption_Values() []string {
	return []string{
		MatchOptionEquals,
		MatchOptionStartsWith,
		MatchOptionEndsWith,
		MatchOptionContains,
		MatchOptionGreaterThanOrEqual,
	}
}
