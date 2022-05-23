// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GatewayDescription gateway description
// Example: Sample Gateway description
//
// swagger:model gateway_description
type GatewayDescription string

// Validate validates this gateway description
func (m GatewayDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gateway description based on context it is used
func (m GatewayDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
