// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostUpdatePushTokenPayload post update push token payload
//
// swagger:model postUpdatePushTokenPayload
type PostUpdatePushTokenPayload struct {

	// New push token for given provider.
	// Required: true
	// Max Length: 500
	NewToken *string `json:"newToken"`

	// Old token that can be deleted if present.
	// Max Length: 500
	OldToken *string `json:"oldToken,omitempty"`

	// Identifier of the provider the token is for (eg. "fcm", "apn"). Currently only "fcm" is supported.
	// Required: true
	// Max Length: 500
	Provider *string `json:"provider"`
}

// Validate validates this post update push token payload
func (m *PostUpdatePushTokenPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostUpdatePushTokenPayload) validateNewToken(formats strfmt.Registry) error {

	if err := validate.Required("newToken", "body", m.NewToken); err != nil {
		return err
	}

	if err := validate.MaxLength("newToken", "body", string(*m.NewToken), 500); err != nil {
		return err
	}

	return nil
}

func (m *PostUpdatePushTokenPayload) validateOldToken(formats strfmt.Registry) error {

	if swag.IsZero(m.OldToken) { // not required
		return nil
	}

	if err := validate.MaxLength("oldToken", "body", string(*m.OldToken), 500); err != nil {
		return err
	}

	return nil
}

func (m *PostUpdatePushTokenPayload) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", m.Provider); err != nil {
		return err
	}

	if err := validate.MaxLength("provider", "body", string(*m.Provider), 500); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostUpdatePushTokenPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUpdatePushTokenPayload) UnmarshalBinary(b []byte) error {
	var res PostUpdatePushTokenPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}