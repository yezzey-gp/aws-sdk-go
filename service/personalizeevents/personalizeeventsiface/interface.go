// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package personalizeeventsiface provides an interface to enable mocking the Amazon Personalize Events service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package personalizeeventsiface

import (
	"github.com/yezzey-gp/aws-sdk-go/aws"
	"github.com/yezzey-gp/aws-sdk-go/aws/request"
	"github.com/yezzey-gp/aws-sdk-go/service/personalizeevents"
)

// PersonalizeEventsAPI provides an interface to enable mocking the
// personalizeevents.PersonalizeEvents service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Personalize Events.
//	func myFunc(svc personalizeeventsiface.PersonalizeEventsAPI) bool {
//	    // Make svc.PutActionInteractions request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := personalizeevents.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockPersonalizeEventsClient struct {
//	    personalizeeventsiface.PersonalizeEventsAPI
//	}
//	func (m *mockPersonalizeEventsClient) PutActionInteractions(input *personalizeevents.PutActionInteractionsInput) (*personalizeevents.PutActionInteractionsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockPersonalizeEventsClient{}
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
type PersonalizeEventsAPI interface {
	PutActionInteractions(*personalizeevents.PutActionInteractionsInput) (*personalizeevents.PutActionInteractionsOutput, error)
	PutActionInteractionsWithContext(aws.Context, *personalizeevents.PutActionInteractionsInput, ...request.Option) (*personalizeevents.PutActionInteractionsOutput, error)
	PutActionInteractionsRequest(*personalizeevents.PutActionInteractionsInput) (*request.Request, *personalizeevents.PutActionInteractionsOutput)

	PutActions(*personalizeevents.PutActionsInput) (*personalizeevents.PutActionsOutput, error)
	PutActionsWithContext(aws.Context, *personalizeevents.PutActionsInput, ...request.Option) (*personalizeevents.PutActionsOutput, error)
	PutActionsRequest(*personalizeevents.PutActionsInput) (*request.Request, *personalizeevents.PutActionsOutput)

	PutEvents(*personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error)
	PutEventsWithContext(aws.Context, *personalizeevents.PutEventsInput, ...request.Option) (*personalizeevents.PutEventsOutput, error)
	PutEventsRequest(*personalizeevents.PutEventsInput) (*request.Request, *personalizeevents.PutEventsOutput)

	PutItems(*personalizeevents.PutItemsInput) (*personalizeevents.PutItemsOutput, error)
	PutItemsWithContext(aws.Context, *personalizeevents.PutItemsInput, ...request.Option) (*personalizeevents.PutItemsOutput, error)
	PutItemsRequest(*personalizeevents.PutItemsInput) (*request.Request, *personalizeevents.PutItemsOutput)

	PutUsers(*personalizeevents.PutUsersInput) (*personalizeevents.PutUsersOutput, error)
	PutUsersWithContext(aws.Context, *personalizeevents.PutUsersInput, ...request.Option) (*personalizeevents.PutUsersOutput, error)
	PutUsersRequest(*personalizeevents.PutUsersInput) (*request.Request, *personalizeevents.PutUsersOutput)
}

var _ PersonalizeEventsAPI = (*personalizeevents.PersonalizeEvents)(nil)
