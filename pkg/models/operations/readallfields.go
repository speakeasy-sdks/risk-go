// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/risk-go/pkg/models/shared"
	"github.com/speakeasy-sdks/risk-go/pkg/utils"
	"net/http"
)

// ReadAllFieldsFieldTypeFieldTypeFilter - A field type where, if provided, the response will only contain fields of the identified field type
type ReadAllFieldsFieldTypeFieldTypeFilter string

const (
	ReadAllFieldsFieldTypeFieldTypeFilterDate            ReadAllFieldsFieldTypeFieldTypeFilter = "DATE"
	ReadAllFieldsFieldTypeFieldTypeFilterUser            ReadAllFieldsFieldTypeFieldTypeFilter = "USER"
	ReadAllFieldsFieldTypeFieldTypeFilterExternalUser    ReadAllFieldsFieldTypeFieldTypeFilter = "EXTERNAL_USER"
	ReadAllFieldsFieldTypeFieldTypeFilterText            ReadAllFieldsFieldTypeFieldTypeFilter = "TEXT"
	ReadAllFieldsFieldTypeFieldTypeFilterTextArea        ReadAllFieldsFieldTypeFieldTypeFilter = "TEXT_AREA"
	ReadAllFieldsFieldTypeFieldTypeFilterNumber          ReadAllFieldsFieldTypeFieldTypeFilter = "NUMBER"
	ReadAllFieldsFieldTypeFieldTypeFilterESignature      ReadAllFieldsFieldTypeFieldTypeFilter = "E_SIGNATURE"
	ReadAllFieldsFieldTypeFieldTypeFilterCheckbox        ReadAllFieldsFieldTypeFieldTypeFilter = "CHECKBOX"
	ReadAllFieldsFieldTypeFieldTypeFilterMultiSelect     ReadAllFieldsFieldTypeFieldTypeFilter = "MULTI_SELECT"
	ReadAllFieldsFieldTypeFieldTypeFilterRadio           ReadAllFieldsFieldTypeFieldTypeFilter = "RADIO"
	ReadAllFieldsFieldTypeFieldTypeFilterSelect          ReadAllFieldsFieldTypeFieldTypeFilter = "SELECT"
	ReadAllFieldsFieldTypeFieldTypeFilterAttachment      ReadAllFieldsFieldTypeFieldTypeFilter = "ATTACHMENT"
	ReadAllFieldsFieldTypeFieldTypeFilterCalculation     ReadAllFieldsFieldTypeFieldTypeFilter = "CALCULATION"
	ReadAllFieldsFieldTypeFieldTypeFilterDateCalculation ReadAllFieldsFieldTypeFieldTypeFilter = "DATE_CALCULATION"
)

func (e ReadAllFieldsFieldTypeFieldTypeFilter) ToPointer() *ReadAllFieldsFieldTypeFieldTypeFilter {
	return &e
}

func (e *ReadAllFieldsFieldTypeFieldTypeFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DATE":
		fallthrough
	case "USER":
		fallthrough
	case "EXTERNAL_USER":
		fallthrough
	case "TEXT":
		fallthrough
	case "TEXT_AREA":
		fallthrough
	case "NUMBER":
		fallthrough
	case "E_SIGNATURE":
		fallthrough
	case "CHECKBOX":
		fallthrough
	case "MULTI_SELECT":
		fallthrough
	case "RADIO":
		fallthrough
	case "SELECT":
		fallthrough
	case "ATTACHMENT":
		fallthrough
	case "CALCULATION":
		fallthrough
	case "DATE_CALCULATION":
		*e = ReadAllFieldsFieldTypeFieldTypeFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReadAllFieldsFieldTypeFieldTypeFilter: %v", v)
	}
}

type ReadAllFieldsRequest struct {
	// The unique ID of an application where, if provided, the response will only contain fields from the identified application
	ApplicationID *string `default:"" queryParam:"style=form,explode=true,name=application-id"`
	// A field type where, if provided, the response will only contain fields of the identified field type
	FieldType *ReadAllFieldsFieldTypeFieldTypeFilter `queryParam:"style=form,explode=true,name=field-type"`
	// The zero-indexed page number (must not be less than 0, defaults to 0)
	Page *int `queryParam:"style=form,explode=true,name=page"`
	// The size of the page and maximum number of items to be returned (must not be less than 1, defaults to 20)
	Size *int `queryParam:"style=form,explode=true,name=size"`
	// The unique ID of a step where, if provided, the response will only contain fields from the identified step
	StepID *string `default:"" queryParam:"style=form,explode=true,name=step-id"`
	// The unique ID of a workflow where, if provided, the response will only contain fields from the identified workflow
	WorkflowID *string `default:"" queryParam:"style=form,explode=true,name=workflow-id"`
}

func (r ReadAllFieldsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReadAllFieldsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReadAllFieldsRequest) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *ReadAllFieldsRequest) GetFieldType() *ReadAllFieldsFieldTypeFieldTypeFilter {
	if o == nil {
		return nil
	}
	return o.FieldType
}

func (o *ReadAllFieldsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ReadAllFieldsRequest) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ReadAllFieldsRequest) GetStepID() *string {
	if o == nil {
		return nil
	}
	return o.StepID
}

func (o *ReadAllFieldsRequest) GetWorkflowID() *string {
	if o == nil {
		return nil
	}
	return o.WorkflowID
}

type ReadAllFieldsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PageModelOutFieldAPIOut *shared.PageModelOutFieldAPIOut
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReadAllFieldsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReadAllFieldsResponse) GetPageModelOutFieldAPIOut() *shared.PageModelOutFieldAPIOut {
	if o == nil {
		return nil
	}
	return o.PageModelOutFieldAPIOut
}

func (o *ReadAllFieldsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReadAllFieldsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
