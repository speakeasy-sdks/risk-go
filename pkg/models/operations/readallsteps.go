// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/utils"
	"net/http"
)

type ReadAllStepsRequest struct {
	// The zero-indexed page number (must not be less than 0, defaults to 0)
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)
	Size *int `queryParam:"style=form,explode=true,name=size"`
	// The unique ID of a workflow where, if provided, the response will only contain steps from the identified workflow
	WorkflowID *string `default:"" queryParam:"style=form,explode=true,name=workflow-id"`
}

func (r ReadAllStepsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReadAllStepsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReadAllStepsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ReadAllStepsRequest) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ReadAllStepsRequest) GetWorkflowID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowID
}

type ReadAllStepsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PageModelOutStepAPIOut *shared.PageModelOutStepAPIOut
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReadAllStepsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadAllStepsResponse) GetPageModelOutStepAPIOut() *shared.PageModelOutStepAPIOut {
	if o == nil {
		return nil
	}
	return o.PageModelOutStepAPIOut
}

func (o *ReadAllStepsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadAllStepsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
