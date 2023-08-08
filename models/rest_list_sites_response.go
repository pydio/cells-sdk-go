// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestListSitesResponse Response with declared sites
//
// swagger:model restListSitesResponse
type RestListSitesResponse struct {

	// sites
	Sites []*InstallProxyConfig `json:"Sites"`
}

// Validate validates this rest list sites response
func (m *RestListSitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestListSitesResponse) validateSites(formats strfmt.Registry) error {
	if swag.IsZero(m.Sites) { // not required
		return nil
	}

	for i := 0; i < len(m.Sites); i++ {
		if swag.IsZero(m.Sites[i]) { // not required
			continue
		}

		if m.Sites[i] != nil {
			if err := m.Sites[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sites" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest list sites response based on the context it is used
func (m *RestListSitesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSites(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestListSitesResponse) contextValidateSites(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sites); i++ {

		if m.Sites[i] != nil {

			if swag.IsZero(m.Sites[i]) { // not required
				return nil
			}

			if err := m.Sites[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sites" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sites" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestListSitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestListSitesResponse) UnmarshalBinary(b []byte) error {
	var res RestListSitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
