// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"net/http"
)

type UpdateWorkflowRequest struct {
	WorkflowAPIUpdateIn shared.WorkflowAPIUpdateIn `request:"mediaType=application/json"`
	// The unique ID of the workflow
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateWorkflowRequest) GetWorkflowAPIUpdateIn() shared.WorkflowAPIUpdateIn {
	if o == nil {
		return shared.WorkflowAPIUpdateIn{}
	}
	return o.WorkflowAPIUpdateIn
}

func (o *UpdateWorkflowRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateWorkflowResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WorkflowAPIOut *shared.WorkflowAPIOut
}

func (o *UpdateWorkflowResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWorkflowResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWorkflowResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWorkflowResponse) GetWorkflowAPIOut() *shared.WorkflowAPIOut {
	if o == nil {
		return nil
	}
	return o.WorkflowAPIOut
}
