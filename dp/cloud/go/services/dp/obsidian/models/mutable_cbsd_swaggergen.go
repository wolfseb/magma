// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MutableCbsd mutable cbsd
// swagger:model mutable_cbsd
type MutableCbsd struct {

	// capabilities
	// Required: true
	Capabilities Capabilities `json:"capabilities"`

	// fcc id
	// Required: true
	// Min Length: 1
	FccID string `json:"fcc_id"`

	// frequency preferences
	// Required: true
	FrequencyPreferences FrequencyPreferences `json:"frequency_preferences"`

	// serial number
	// Required: true
	// Min Length: 1
	SerialNumber string `json:"serial_number"`

	// user id
	// Required: true
	// Min Length: 1
	UserID string `json:"user_id"`
}

// Validate validates this mutable cbsd
func (m *MutableCbsd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFccID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequencyPreferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerialNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MutableCbsd) validateCapabilities(formats strfmt.Registry) error {

	if err := m.Capabilities.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("capabilities")
		}
		return err
	}

	return nil
}

func (m *MutableCbsd) validateFccID(formats strfmt.Registry) error {

	if err := validate.RequiredString("fcc_id", "body", string(m.FccID)); err != nil {
		return err
	}

	if err := validate.MinLength("fcc_id", "body", string(m.FccID), 1); err != nil {
		return err
	}

	return nil
}

func (m *MutableCbsd) validateFrequencyPreferences(formats strfmt.Registry) error {

	if err := m.FrequencyPreferences.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("frequency_preferences")
		}
		return err
	}

	return nil
}

func (m *MutableCbsd) validateSerialNumber(formats strfmt.Registry) error {

	if err := validate.RequiredString("serial_number", "body", string(m.SerialNumber)); err != nil {
		return err
	}

	if err := validate.MinLength("serial_number", "body", string(m.SerialNumber), 1); err != nil {
		return err
	}

	return nil
}

func (m *MutableCbsd) validateUserID(formats strfmt.Registry) error {

	if err := validate.RequiredString("user_id", "body", string(m.UserID)); err != nil {
		return err
	}

	if err := validate.MinLength("user_id", "body", string(m.UserID), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MutableCbsd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MutableCbsd) UnmarshalBinary(b []byte) error {
	var res MutableCbsd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
