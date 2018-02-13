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

// ASN a s n
// swagger:model ASN
type ASN struct {

	// The ASN to retrieve
	//
	// Autonomous System Numbers per APNIC for identifying a service provider
	// Required: true
	ASN *int64 `json:"asn"`

	// Related cachegroup name
	Cachegroup string `json:"cachegroup,omitempty"`

	// Related cachegroup id
	CachegroupID int64 `json:"cachegroupId,omitempty"`

	// ID of the ASN
	// Required: true
	ID *int64 `json:"id"`

	// LastUpdated
	LastUpdated string `json:"lastUpdated,omitempty"`
}

// Validate validates this a s n
func (m *ASN) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateASN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ASN) validateASN(formats strfmt.Registry) error {

	if err := validate.Required("asn", "body", m.ASN); err != nil {
		return err
	}

	return nil
}

func (m *ASN) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ASN) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ASN) UnmarshalBinary(b []byte) error {
	var res ASN
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
