// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestShareLink Model for representing a public link
// swagger:model restShareLink
type RestShareLink struct {

	// Timestamp after which the share is disabled
	AccessEnd string `json:"AccessEnd,omitempty"`

	// Timestamp of start date for enabling the share (not implemented yet)
	AccessStart string `json:"AccessStart,omitempty"`

	// Current number of downloads
	CurrentDownloads string `json:"CurrentDownloads,omitempty"`

	// Description of the Link (max 1000 chars)
	Description string `json:"Description,omitempty"`

	// Label of the Link (max 500 chars)
	Label string `json:"Label,omitempty"`

	// Unique Hash for accessing the link
	LinkHash string `json:"LinkHash,omitempty"`

	// Full URL for accessing the link
	LinkURL string `json:"LinkUrl,omitempty"`

	// Maximum number of downloads until expiration
	MaxDownloads string `json:"MaxDownloads,omitempty"`

	// Whether a password is required or not to access the link
	PasswordRequired bool `json:"PasswordRequired,omitempty"`

	// Specific permissions for public links
	Permissions []RestShareLinkAccessType `json:"Permissions"`

	// Security policies
	Policies []*ServiceResourcePolicy `json:"Policies"`

	// Whether policies are currently editable or not
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// RestrictToTargetUsers enable users restriction
	RestrictToTargetUsers bool `json:"RestrictToTargetUsers,omitempty"`

	// Nodes in the tree that serve as root to this link
	RootNodes []*TreeNode `json:"RootNodes"`

	// TargetUsers can be used to restrict access
	TargetUsers map[string]RestShareLinkTargetUser `json:"TargetUsers,omitempty"`

	// Temporary user Login used to login automatically when accessing this link
	UserLogin string `json:"UserLogin,omitempty"`

	// Temporary user Uuid used to login automatically when accessing this link
	UserUUID string `json:"UserUuid,omitempty"`

	// Internal identifier of the link
	UUID string `json:"Uuid,omitempty"`

	// Display Template for loading the public link
	ViewTemplateName string `json:"ViewTemplateName,omitempty"`
}

// Validate validates this rest share link
func (m *RestShareLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestShareLink) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {

		if err := m.Permissions[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Permissions" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *RestShareLink) validatePolicies(formats strfmt.Registry) error {

	if swag.IsZero(m.Policies) { // not required
		return nil
	}

	for i := 0; i < len(m.Policies); i++ {
		if swag.IsZero(m.Policies[i]) { // not required
			continue
		}

		if m.Policies[i] != nil {
			if err := m.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestShareLink) validateRootNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.RootNodes) { // not required
		return nil
	}

	for i := 0; i < len(m.RootNodes); i++ {
		if swag.IsZero(m.RootNodes[i]) { // not required
			continue
		}

		if m.RootNodes[i] != nil {
			if err := m.RootNodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RootNodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestShareLink) validateTargetUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetUsers) { // not required
		return nil
	}

	for k := range m.TargetUsers {

		if err := validate.Required("TargetUsers"+"."+k, "body", m.TargetUsers[k]); err != nil {
			return err
		}
		if val, ok := m.TargetUsers[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestShareLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestShareLink) UnmarshalBinary(b []byte) error {
	var res RestShareLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
