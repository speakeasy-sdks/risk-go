// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WorkflowAPIOut struct {
	// The unique ID of the parent application of the workflow
	ApplicationID *string `json:"applicationId,omitempty"`
	// The unique ID of this Risk Cloud resource
	ID *string `json:"id,omitempty"`
	// The name of the workflow
	Name *string `json:"name,omitempty"`
	// Identifies the type of object this data represents
	Object *string `json:"object,omitempty"`
	// The prefix to be used in the name of every record created from this workflow
	RecordPrefix *string `json:"recordPrefix,omitempty"`
	// The x-coordinate of the workflow in the application builder
	Xpos *int `json:"xpos,omitempty"`
	// The y-coordinate of the workflow in the application builder
	Ypos *int `json:"ypos,omitempty"`
}

func (o *WorkflowAPIOut) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *WorkflowAPIOut) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WorkflowAPIOut) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WorkflowAPIOut) GetObject() *string {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *WorkflowAPIOut) GetRecordPrefix() *string {
	if o == nil {
		return nil
	}
	return o.RecordPrefix
}

func (o *WorkflowAPIOut) GetXpos() *int {
	if o == nil {
		return nil
	}
	return o.Xpos
}

func (o *WorkflowAPIOut) GetYpos() *int {
	if o == nil {
		return nil
	}
	return o.Ypos
}
