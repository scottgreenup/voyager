// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResolutionResponseType resolution response type
// swagger:model ResolutionResponseType
type ResolutionResponseType struct {

	// label
	Label string `json:"label,omitempty"`

	// release groups
	ReleaseGroups map[string]map[string]interface{} `json:"releaseGroups,omitempty"`

	// service
	Service string `json:"service,omitempty"`
}

// Validate validates this resolution response type
func (m *ResolutionResponseType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResolutionResponseType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResolutionResponseType) UnmarshalBinary(b []byte) error {
	var res ResolutionResponseType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}