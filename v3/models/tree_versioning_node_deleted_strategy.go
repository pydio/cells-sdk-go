// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TreeVersioningNodeDeletedStrategy tree versioning node deleted strategy
//
// swagger:model treeVersioningNodeDeletedStrategy
type TreeVersioningNodeDeletedStrategy string

func NewTreeVersioningNodeDeletedStrategy(value TreeVersioningNodeDeletedStrategy) *TreeVersioningNodeDeletedStrategy {
	v := value
	return &v
}

const (

	// TreeVersioningNodeDeletedStrategyKeepAll captures enum value "KeepAll"
	TreeVersioningNodeDeletedStrategyKeepAll TreeVersioningNodeDeletedStrategy = "KeepAll"

	// TreeVersioningNodeDeletedStrategyKeepLast captures enum value "KeepLast"
	TreeVersioningNodeDeletedStrategyKeepLast TreeVersioningNodeDeletedStrategy = "KeepLast"

	// TreeVersioningNodeDeletedStrategyKeepNone captures enum value "KeepNone"
	TreeVersioningNodeDeletedStrategyKeepNone TreeVersioningNodeDeletedStrategy = "KeepNone"
)

// for schema
var treeVersioningNodeDeletedStrategyEnum []interface{}

func init() {
	var res []TreeVersioningNodeDeletedStrategy
	if err := json.Unmarshal([]byte(`["KeepAll","KeepLast","KeepNone"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		treeVersioningNodeDeletedStrategyEnum = append(treeVersioningNodeDeletedStrategyEnum, v)
	}
}

func (m TreeVersioningNodeDeletedStrategy) validateTreeVersioningNodeDeletedStrategyEnum(path, location string, value TreeVersioningNodeDeletedStrategy) error {
	if err := validate.EnumCase(path, location, value, treeVersioningNodeDeletedStrategyEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tree versioning node deleted strategy
func (m TreeVersioningNodeDeletedStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTreeVersioningNodeDeletedStrategyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tree versioning node deleted strategy based on context it is used
func (m TreeVersioningNodeDeletedStrategy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
