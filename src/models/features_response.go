// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FeaturesResponse features response
// swagger:model FeaturesResponse
type FeaturesResponse struct {

	// data
	Data *FeaturesResponseData `json:"data,omitempty"`
}

// Validate validates this features response
func (m *FeaturesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturesResponse) validateData(formats strfmt.Registry) error {

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
func (m *FeaturesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturesResponse) UnmarshalBinary(b []byte) error {
	var res FeaturesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FeaturesResponseData features response data
// swagger:model FeaturesResponseData
type FeaturesResponseData struct {

	// features
	Features *FeaturesResponseDataFeatures `json:"features,omitempty"`
}

// Validate validates this features response data
func (m *FeaturesResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FeaturesResponseData) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "features")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FeaturesResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturesResponseData) UnmarshalBinary(b []byte) error {
	var res FeaturesResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FeaturesResponseDataFeatures features response data features
// swagger:model FeaturesResponseDataFeatures
type FeaturesResponseDataFeatures struct {

	// bootloader hash
	BootloaderHash string `json:"bootloader_hash,omitempty"`

	// device id
	DeviceID string `json:"device_id,omitempty"`

	// initialized
	Initialized bool `json:"initialized,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// major version
	MajorVersion int64 `json:"major_version,omitempty"`

	// minor version
	MinorVersion int64 `json:"minor_version,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// needs backup
	NeedsBackup bool `json:"needs_backup,omitempty"`

	// passphrase cached
	PassphraseCached bool `json:"passphrase_cached,omitempty"`

	// passphrase protection
	PassphraseProtection bool `json:"passphrase_protection,omitempty"`

	// patch version
	PatchVersion int64 `json:"patch_version,omitempty"`

	// pin cached
	PinCached bool `json:"pin_cached,omitempty"`

	// pin protection
	PinProtection bool `json:"pin_protection,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`
}

// Validate validates this features response data features
func (m *FeaturesResponseDataFeatures) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FeaturesResponseDataFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FeaturesResponseDataFeatures) UnmarshalBinary(b []byte) error {
	var res FeaturesResponseDataFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
