// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmailmessageflow

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGetRawMessageContent = "GetRawMessageContent"

// GetRawMessageContentRequest generates a "aws/request.Request" representing the
// client's request for the GetRawMessageContent operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetRawMessageContent for more information on using the GetRawMessageContent
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetRawMessageContentRequest method.
//    req, resp := client.GetRawMessageContentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *WorkMailMessageFlow) GetRawMessageContentRequest(input *GetRawMessageContentInput) (req *request.Request, output *GetRawMessageContentOutput) {
	op := &request.Operation{
		Name:       opGetRawMessageContent,
		HTTPMethod: "GET",
		HTTPPath:   "/messages/{messageId}",
	}

	if input == nil {
		input = &GetRawMessageContentInput{}
	}

	output = &GetRawMessageContentOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetRawMessageContent API operation for Amazon WorkMail Message Flow.
//
// Retrieves the raw content of an in-transit email message, in MIME format.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon WorkMail Message Flow's
// API operation GetRawMessageContent for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The requested email message is not found.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/workmailmessageflow-2019-05-01/GetRawMessageContent
func (c *WorkMailMessageFlow) GetRawMessageContent(input *GetRawMessageContentInput) (*GetRawMessageContentOutput, error) {
	req, out := c.GetRawMessageContentRequest(input)
	return out, req.Send()
}

// GetRawMessageContentWithContext is the same as GetRawMessageContent with the addition of
// the ability to pass a context and additional request options.
//
// See GetRawMessageContent for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WorkMailMessageFlow) GetRawMessageContentWithContext(ctx aws.Context, input *GetRawMessageContentInput, opts ...request.Option) (*GetRawMessageContentOutput, error) {
	req, out := c.GetRawMessageContentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRawMessageContentInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the email message to retrieve.
	//
	// MessageId is a required field
	MessageId *string `location:"uri" locationName:"messageId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRawMessageContentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRawMessageContentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRawMessageContentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRawMessageContentInput"}
	if s.MessageId == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageId"))
	}
	if s.MessageId != nil && len(*s.MessageId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("MessageId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMessageId sets the MessageId field's value.
func (s *GetRawMessageContentInput) SetMessageId(v string) *GetRawMessageContentInput {
	s.MessageId = &v
	return s
}

type GetRawMessageContentOutput struct {
	_ struct{} `type:"structure" payload:"MessageContent"`

	// The raw content of the email message, in MIME format.
	//
	// MessageContent is a required field
	MessageContent io.ReadCloser `locationName:"messageContent" type:"blob" required:"true"`
}

// String returns the string representation
func (s GetRawMessageContentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRawMessageContentOutput) GoString() string {
	return s.String()
}

// SetMessageContent sets the MessageContent field's value.
func (s *GetRawMessageContentOutput) SetMessageContent(v io.ReadCloser) *GetRawMessageContentOutput {
	s.MessageContent = v
	return s
}