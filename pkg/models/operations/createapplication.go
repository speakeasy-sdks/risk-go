// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"net/http"
)

type CreateApplicationRequest struct {
	ApplicationAPICreateIn shared.ApplicationAPICreateIn `request:"mediaType=application/json"`
}

func (o *CreateApplicationRequest) GetApplicationAPICreateIn() shared.ApplicationAPICreateIn {
	if o == nil {
		return shared.ApplicationAPICreateIn{}
	}
	return o.ApplicationAPICreateIn
}

type CreateApplicationResponse struct {
	// OK
	ApplicationAPIOut *shared.ApplicationAPIOut
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateApplicationResponse) GetApplicationAPIOut() *shared.ApplicationAPIOut {
	if o == nil {
		return nil
	}
	return o.ApplicationAPIOut
}

func (o *CreateApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
