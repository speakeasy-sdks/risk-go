// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"net/http"
)

type UpdateRequest struct {
	StepAPIUpdateIn shared.StepAPIUpdateIn `request:"mediaType=application/json"`
	// The unique ID of the step
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateRequest) GetStepAPIUpdateIn() shared.StepAPIUpdateIn {
	if o == nil {
		return shared.StepAPIUpdateIn{}
	}
	return o.StepAPIUpdateIn
}

func (o *UpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	StepAPIOut *shared.StepAPIOut
}

func (o *UpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateResponse) GetStepAPIOut() *shared.StepAPIOut {
	if o == nil {
		return nil
	}
	return o.StepAPIOut
}
