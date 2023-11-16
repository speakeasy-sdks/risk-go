// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"net/http"
)

type DeleteApplicationRequest struct {
	// The unique ID of the application
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteApplicationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteApplicationResponse struct {
	// OK
	AppActionResponse *shared.AppActionResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteApplicationResponse) GetAppActionResponse() *shared.AppActionResponse {
	if o == nil {
		return nil
	}
	return o.AppActionResponse
}

func (o *DeleteApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
