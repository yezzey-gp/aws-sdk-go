// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appfabriciface provides an interface to enable mocking the AppFabric service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appfabriciface

import (
	"github.com/yezzey-gp/aws-sdk-go/aws"
	"github.com/yezzey-gp/aws-sdk-go/aws/request"
	"github.com/yezzey-gp/aws-sdk-go/service/appfabric"
)

// AppFabricAPI provides an interface to enable mocking the
// appfabric.AppFabric service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AppFabric.
//	func myFunc(svc appfabriciface.AppFabricAPI) bool {
//	    // Make svc.BatchGetUserAccessTasks request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := appfabric.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAppFabricClient struct {
//	    appfabriciface.AppFabricAPI
//	}
//	func (m *mockAppFabricClient) BatchGetUserAccessTasks(input *appfabric.BatchGetUserAccessTasksInput) (*appfabric.BatchGetUserAccessTasksOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAppFabricClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AppFabricAPI interface {
	BatchGetUserAccessTasks(*appfabric.BatchGetUserAccessTasksInput) (*appfabric.BatchGetUserAccessTasksOutput, error)
	BatchGetUserAccessTasksWithContext(aws.Context, *appfabric.BatchGetUserAccessTasksInput, ...request.Option) (*appfabric.BatchGetUserAccessTasksOutput, error)
	BatchGetUserAccessTasksRequest(*appfabric.BatchGetUserAccessTasksInput) (*request.Request, *appfabric.BatchGetUserAccessTasksOutput)

	ConnectAppAuthorization(*appfabric.ConnectAppAuthorizationInput) (*appfabric.ConnectAppAuthorizationOutput, error)
	ConnectAppAuthorizationWithContext(aws.Context, *appfabric.ConnectAppAuthorizationInput, ...request.Option) (*appfabric.ConnectAppAuthorizationOutput, error)
	ConnectAppAuthorizationRequest(*appfabric.ConnectAppAuthorizationInput) (*request.Request, *appfabric.ConnectAppAuthorizationOutput)

	CreateAppAuthorization(*appfabric.CreateAppAuthorizationInput) (*appfabric.CreateAppAuthorizationOutput, error)
	CreateAppAuthorizationWithContext(aws.Context, *appfabric.CreateAppAuthorizationInput, ...request.Option) (*appfabric.CreateAppAuthorizationOutput, error)
	CreateAppAuthorizationRequest(*appfabric.CreateAppAuthorizationInput) (*request.Request, *appfabric.CreateAppAuthorizationOutput)

	CreateAppBundle(*appfabric.CreateAppBundleInput) (*appfabric.CreateAppBundleOutput, error)
	CreateAppBundleWithContext(aws.Context, *appfabric.CreateAppBundleInput, ...request.Option) (*appfabric.CreateAppBundleOutput, error)
	CreateAppBundleRequest(*appfabric.CreateAppBundleInput) (*request.Request, *appfabric.CreateAppBundleOutput)

	CreateIngestion(*appfabric.CreateIngestionInput) (*appfabric.CreateIngestionOutput, error)
	CreateIngestionWithContext(aws.Context, *appfabric.CreateIngestionInput, ...request.Option) (*appfabric.CreateIngestionOutput, error)
	CreateIngestionRequest(*appfabric.CreateIngestionInput) (*request.Request, *appfabric.CreateIngestionOutput)

	CreateIngestionDestination(*appfabric.CreateIngestionDestinationInput) (*appfabric.CreateIngestionDestinationOutput, error)
	CreateIngestionDestinationWithContext(aws.Context, *appfabric.CreateIngestionDestinationInput, ...request.Option) (*appfabric.CreateIngestionDestinationOutput, error)
	CreateIngestionDestinationRequest(*appfabric.CreateIngestionDestinationInput) (*request.Request, *appfabric.CreateIngestionDestinationOutput)

	DeleteAppAuthorization(*appfabric.DeleteAppAuthorizationInput) (*appfabric.DeleteAppAuthorizationOutput, error)
	DeleteAppAuthorizationWithContext(aws.Context, *appfabric.DeleteAppAuthorizationInput, ...request.Option) (*appfabric.DeleteAppAuthorizationOutput, error)
	DeleteAppAuthorizationRequest(*appfabric.DeleteAppAuthorizationInput) (*request.Request, *appfabric.DeleteAppAuthorizationOutput)

	DeleteAppBundle(*appfabric.DeleteAppBundleInput) (*appfabric.DeleteAppBundleOutput, error)
	DeleteAppBundleWithContext(aws.Context, *appfabric.DeleteAppBundleInput, ...request.Option) (*appfabric.DeleteAppBundleOutput, error)
	DeleteAppBundleRequest(*appfabric.DeleteAppBundleInput) (*request.Request, *appfabric.DeleteAppBundleOutput)

	DeleteIngestion(*appfabric.DeleteIngestionInput) (*appfabric.DeleteIngestionOutput, error)
	DeleteIngestionWithContext(aws.Context, *appfabric.DeleteIngestionInput, ...request.Option) (*appfabric.DeleteIngestionOutput, error)
	DeleteIngestionRequest(*appfabric.DeleteIngestionInput) (*request.Request, *appfabric.DeleteIngestionOutput)

	DeleteIngestionDestination(*appfabric.DeleteIngestionDestinationInput) (*appfabric.DeleteIngestionDestinationOutput, error)
	DeleteIngestionDestinationWithContext(aws.Context, *appfabric.DeleteIngestionDestinationInput, ...request.Option) (*appfabric.DeleteIngestionDestinationOutput, error)
	DeleteIngestionDestinationRequest(*appfabric.DeleteIngestionDestinationInput) (*request.Request, *appfabric.DeleteIngestionDestinationOutput)

	GetAppAuthorization(*appfabric.GetAppAuthorizationInput) (*appfabric.GetAppAuthorizationOutput, error)
	GetAppAuthorizationWithContext(aws.Context, *appfabric.GetAppAuthorizationInput, ...request.Option) (*appfabric.GetAppAuthorizationOutput, error)
	GetAppAuthorizationRequest(*appfabric.GetAppAuthorizationInput) (*request.Request, *appfabric.GetAppAuthorizationOutput)

	GetAppBundle(*appfabric.GetAppBundleInput) (*appfabric.GetAppBundleOutput, error)
	GetAppBundleWithContext(aws.Context, *appfabric.GetAppBundleInput, ...request.Option) (*appfabric.GetAppBundleOutput, error)
	GetAppBundleRequest(*appfabric.GetAppBundleInput) (*request.Request, *appfabric.GetAppBundleOutput)

	GetIngestion(*appfabric.GetIngestionInput) (*appfabric.GetIngestionOutput, error)
	GetIngestionWithContext(aws.Context, *appfabric.GetIngestionInput, ...request.Option) (*appfabric.GetIngestionOutput, error)
	GetIngestionRequest(*appfabric.GetIngestionInput) (*request.Request, *appfabric.GetIngestionOutput)

	GetIngestionDestination(*appfabric.GetIngestionDestinationInput) (*appfabric.GetIngestionDestinationOutput, error)
	GetIngestionDestinationWithContext(aws.Context, *appfabric.GetIngestionDestinationInput, ...request.Option) (*appfabric.GetIngestionDestinationOutput, error)
	GetIngestionDestinationRequest(*appfabric.GetIngestionDestinationInput) (*request.Request, *appfabric.GetIngestionDestinationOutput)

	ListAppAuthorizations(*appfabric.ListAppAuthorizationsInput) (*appfabric.ListAppAuthorizationsOutput, error)
	ListAppAuthorizationsWithContext(aws.Context, *appfabric.ListAppAuthorizationsInput, ...request.Option) (*appfabric.ListAppAuthorizationsOutput, error)
	ListAppAuthorizationsRequest(*appfabric.ListAppAuthorizationsInput) (*request.Request, *appfabric.ListAppAuthorizationsOutput)

	ListAppAuthorizationsPages(*appfabric.ListAppAuthorizationsInput, func(*appfabric.ListAppAuthorizationsOutput, bool) bool) error
	ListAppAuthorizationsPagesWithContext(aws.Context, *appfabric.ListAppAuthorizationsInput, func(*appfabric.ListAppAuthorizationsOutput, bool) bool, ...request.Option) error

	ListAppBundles(*appfabric.ListAppBundlesInput) (*appfabric.ListAppBundlesOutput, error)
	ListAppBundlesWithContext(aws.Context, *appfabric.ListAppBundlesInput, ...request.Option) (*appfabric.ListAppBundlesOutput, error)
	ListAppBundlesRequest(*appfabric.ListAppBundlesInput) (*request.Request, *appfabric.ListAppBundlesOutput)

	ListAppBundlesPages(*appfabric.ListAppBundlesInput, func(*appfabric.ListAppBundlesOutput, bool) bool) error
	ListAppBundlesPagesWithContext(aws.Context, *appfabric.ListAppBundlesInput, func(*appfabric.ListAppBundlesOutput, bool) bool, ...request.Option) error

	ListIngestionDestinations(*appfabric.ListIngestionDestinationsInput) (*appfabric.ListIngestionDestinationsOutput, error)
	ListIngestionDestinationsWithContext(aws.Context, *appfabric.ListIngestionDestinationsInput, ...request.Option) (*appfabric.ListIngestionDestinationsOutput, error)
	ListIngestionDestinationsRequest(*appfabric.ListIngestionDestinationsInput) (*request.Request, *appfabric.ListIngestionDestinationsOutput)

	ListIngestionDestinationsPages(*appfabric.ListIngestionDestinationsInput, func(*appfabric.ListIngestionDestinationsOutput, bool) bool) error
	ListIngestionDestinationsPagesWithContext(aws.Context, *appfabric.ListIngestionDestinationsInput, func(*appfabric.ListIngestionDestinationsOutput, bool) bool, ...request.Option) error

	ListIngestions(*appfabric.ListIngestionsInput) (*appfabric.ListIngestionsOutput, error)
	ListIngestionsWithContext(aws.Context, *appfabric.ListIngestionsInput, ...request.Option) (*appfabric.ListIngestionsOutput, error)
	ListIngestionsRequest(*appfabric.ListIngestionsInput) (*request.Request, *appfabric.ListIngestionsOutput)

	ListIngestionsPages(*appfabric.ListIngestionsInput, func(*appfabric.ListIngestionsOutput, bool) bool) error
	ListIngestionsPagesWithContext(aws.Context, *appfabric.ListIngestionsInput, func(*appfabric.ListIngestionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*appfabric.ListTagsForResourceInput) (*appfabric.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *appfabric.ListTagsForResourceInput, ...request.Option) (*appfabric.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*appfabric.ListTagsForResourceInput) (*request.Request, *appfabric.ListTagsForResourceOutput)

	StartIngestion(*appfabric.StartIngestionInput) (*appfabric.StartIngestionOutput, error)
	StartIngestionWithContext(aws.Context, *appfabric.StartIngestionInput, ...request.Option) (*appfabric.StartIngestionOutput, error)
	StartIngestionRequest(*appfabric.StartIngestionInput) (*request.Request, *appfabric.StartIngestionOutput)

	StartUserAccessTasks(*appfabric.StartUserAccessTasksInput) (*appfabric.StartUserAccessTasksOutput, error)
	StartUserAccessTasksWithContext(aws.Context, *appfabric.StartUserAccessTasksInput, ...request.Option) (*appfabric.StartUserAccessTasksOutput, error)
	StartUserAccessTasksRequest(*appfabric.StartUserAccessTasksInput) (*request.Request, *appfabric.StartUserAccessTasksOutput)

	StopIngestion(*appfabric.StopIngestionInput) (*appfabric.StopIngestionOutput, error)
	StopIngestionWithContext(aws.Context, *appfabric.StopIngestionInput, ...request.Option) (*appfabric.StopIngestionOutput, error)
	StopIngestionRequest(*appfabric.StopIngestionInput) (*request.Request, *appfabric.StopIngestionOutput)

	TagResource(*appfabric.TagResourceInput) (*appfabric.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *appfabric.TagResourceInput, ...request.Option) (*appfabric.TagResourceOutput, error)
	TagResourceRequest(*appfabric.TagResourceInput) (*request.Request, *appfabric.TagResourceOutput)

	UntagResource(*appfabric.UntagResourceInput) (*appfabric.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *appfabric.UntagResourceInput, ...request.Option) (*appfabric.UntagResourceOutput, error)
	UntagResourceRequest(*appfabric.UntagResourceInput) (*request.Request, *appfabric.UntagResourceOutput)

	UpdateAppAuthorization(*appfabric.UpdateAppAuthorizationInput) (*appfabric.UpdateAppAuthorizationOutput, error)
	UpdateAppAuthorizationWithContext(aws.Context, *appfabric.UpdateAppAuthorizationInput, ...request.Option) (*appfabric.UpdateAppAuthorizationOutput, error)
	UpdateAppAuthorizationRequest(*appfabric.UpdateAppAuthorizationInput) (*request.Request, *appfabric.UpdateAppAuthorizationOutput)

	UpdateIngestionDestination(*appfabric.UpdateIngestionDestinationInput) (*appfabric.UpdateIngestionDestinationOutput, error)
	UpdateIngestionDestinationWithContext(aws.Context, *appfabric.UpdateIngestionDestinationInput, ...request.Option) (*appfabric.UpdateIngestionDestinationOutput, error)
	UpdateIngestionDestinationRequest(*appfabric.UpdateIngestionDestinationInput) (*request.Request, *appfabric.UpdateIngestionDestinationOutput)
}

var _ AppFabricAPI = (*appfabric.AppFabric)(nil)
