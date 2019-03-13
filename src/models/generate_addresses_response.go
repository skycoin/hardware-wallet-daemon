// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GenerateAddressesResponse generate addresses response
// swagger:model GenerateAddressesResponse
type GenerateAddressesResponse struct {

	// data
	Data *GenerateAddressesResponseData `json:"data,omitempty"`
}

// Validate validates this generate addresses response
func (m *GenerateAddressesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateAddressesResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenerateAddressesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateAddressesResponse) UnmarshalBinary(b []byte) error {
	var res GenerateAddressesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GenerateAddressesResponseData generate addresses response data
// swagger:model GenerateAddressesResponseData
type GenerateAddressesResponseData struct {

	// addresses
	Addresses []string `json:"addresses"`
}

// Validate validates this generate addresses response data
func (m *GenerateAddressesResponseData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenerateAddressesResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateAddressesResponseData) UnmarshalBinary(b []byte) error {
	var res GenerateAddressesResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
