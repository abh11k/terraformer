// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the request to create a hosted
// zone.
type GetReusableDelegationSetLimitInput struct {
	_ struct{} `type:"structure"`

	// The ID of the delegation set that you want to get the limit for.
	//
	// DelegationSetId is a required field
	DelegationSetId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// Specify MAX_ZONES_BY_REUSABLE_DELEGATION_SET to get the maximum number of
	// hosted zones that you can associate with the specified reusable delegation
	// set.
	//
	// Type is a required field
	Type ReusableDelegationSetLimitType `location:"uri" locationName:"Type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetReusableDelegationSetLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReusableDelegationSetLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReusableDelegationSetLimitInput"}

	if s.DelegationSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DelegationSetId"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetReusableDelegationSetLimitInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DelegationSetId != nil {
		v := *s.DelegationSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Type", v, metadata)
	}
	return nil
}

// A complex type that contains the requested limit.
type GetReusableDelegationSetLimitOutput struct {
	_ struct{} `type:"structure"`

	// The current number of hosted zones that you can associate with the specified
	// reusable delegation set.
	//
	// Count is a required field
	Count *int64 `type:"long" required:"true"`

	// The current setting for the limit on hosted zones that you can associate
	// with the specified reusable delegation set.
	//
	// Limit is a required field
	Limit *ReusableDelegationSetLimit `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetReusableDelegationSetLimitOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetReusableDelegationSetLimitOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Count != nil {
		v := *s.Count

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Count", protocol.Int64Value(v), metadata)
	}
	if s.Limit != nil {
		v := s.Limit

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Limit", v, metadata)
	}
	return nil
}

const opGetReusableDelegationSetLimit = "GetReusableDelegationSetLimit"

// GetReusableDelegationSetLimitRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the maximum number of hosted zones that you can associate with the specified
// reusable delegation set.
//
// For the default limit, see Limits (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a
// case (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53).
//
//    // Example sending a request using GetReusableDelegationSetLimitRequest.
//    req := client.GetReusableDelegationSetLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetReusableDelegationSetLimit
func (c *Client) GetReusableDelegationSetLimitRequest(input *GetReusableDelegationSetLimitInput) GetReusableDelegationSetLimitRequest {
	op := &aws.Operation{
		Name:       opGetReusableDelegationSetLimit,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/reusabledelegationsetlimit/{Id}/{Type}",
	}

	if input == nil {
		input = &GetReusableDelegationSetLimitInput{}
	}

	req := c.newRequest(op, input, &GetReusableDelegationSetLimitOutput{})
	return GetReusableDelegationSetLimitRequest{Request: req, Input: input, Copy: c.GetReusableDelegationSetLimitRequest}
}

// GetReusableDelegationSetLimitRequest is the request type for the
// GetReusableDelegationSetLimit API operation.
type GetReusableDelegationSetLimitRequest struct {
	*aws.Request
	Input *GetReusableDelegationSetLimitInput
	Copy  func(*GetReusableDelegationSetLimitInput) GetReusableDelegationSetLimitRequest
}

// Send marshals and sends the GetReusableDelegationSetLimit API request.
func (r GetReusableDelegationSetLimitRequest) Send(ctx context.Context) (*GetReusableDelegationSetLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReusableDelegationSetLimitResponse{
		GetReusableDelegationSetLimitOutput: r.Request.Data.(*GetReusableDelegationSetLimitOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReusableDelegationSetLimitResponse is the response type for the
// GetReusableDelegationSetLimit API operation.
type GetReusableDelegationSetLimitResponse struct {
	*GetReusableDelegationSetLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReusableDelegationSetLimit request.
func (r *GetReusableDelegationSetLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
