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

// ActivityObject activity object
// swagger:model activityObject
type ActivityObject struct {

	// Place Properties
	Accuracy float32 `json:"accuracy,omitempty"`

	// Activity Properties
	Actor *ActivityObject `json:"actor,omitempty"`

	// altitude
	Altitude float32 `json:"altitude,omitempty"`

	// any of
	AnyOf *ActivityObject `json:"anyOf,omitempty"`

	// attachment
	Attachment *ActivityObject `json:"attachment,omitempty"`

	// attributed to
	AttributedTo *ActivityObject `json:"attributedTo,omitempty"`

	// audience
	Audience *ActivityObject `json:"audience,omitempty"`

	// bcc
	Bcc *ActivityObject `json:"bcc,omitempty"`

	// bto
	Bto *ActivityObject `json:"bto,omitempty"`

	// cc
	Cc *ActivityObject `json:"cc,omitempty"`

	// closed
	// Format: date-time
	Closed strfmt.DateTime `json:"closed,omitempty"`

	// content
	Content *ActivityObject `json:"content,omitempty"`

	// context
	Context *ActivityObject `json:"context,omitempty"`

	// current
	Current *ActivityObject `json:"current,omitempty"`

	// deleted
	// Format: date-time
	Deleted strfmt.DateTime `json:"deleted,omitempty"`

	// duration
	// Format: date-time
	Duration strfmt.DateTime `json:"duration,omitempty"`

	// end time
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// first
	First *ActivityObject `json:"first,omitempty"`

	// Tombstone Properties
	FormerType ActivityObjectType `json:"formerType,omitempty"`

	// generator
	Generator *ActivityObject `json:"generator,omitempty"`

	// height
	Height int32 `json:"height,omitempty"`

	// Link Properties
	Href string `json:"href,omitempty"`

	// hreflang
	Hreflang string `json:"hreflang,omitempty"`

	// icon
	Icon *ActivityObject `json:"icon,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// image
	Image *ActivityObject `json:"image,omitempty"`

	// in reply to
	InReplyTo *ActivityObject `json:"inReplyTo,omitempty"`

	// instrument
	Instrument *ActivityObject `json:"instrument,omitempty"`

	// Collection Properties
	Items []*ActivityObject `json:"items"`

	// json ld context
	JSONLdContext string `json:"jsonLdContext,omitempty"`

	// last
	Last *ActivityObject `json:"last,omitempty"`

	// latitude
	Latitude float32 `json:"latitude,omitempty"`

	// location
	Location *ActivityObject `json:"location,omitempty"`

	// longitude
	Longitude float32 `json:"longitude,omitempty"`

	// media type
	MediaType string `json:"mediaType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// next
	Next *ActivityObject `json:"next,omitempty"`

	// object
	Object *ActivityObject `json:"object,omitempty"`

	// Question Properties
	OneOf *ActivityObject `json:"oneOf,omitempty"`

	// origin
	Origin *ActivityObject `json:"origin,omitempty"`

	// part of
	PartOf *ActivityObject `json:"partOf,omitempty"`

	// prev
	Prev *ActivityObject `json:"prev,omitempty"`

	// preview
	Preview *ActivityObject `json:"preview,omitempty"`

	// published
	// Format: date-time
	Published strfmt.DateTime `json:"published,omitempty"`

	// radius
	Radius float32 `json:"radius,omitempty"`

	// rel
	Rel string `json:"rel,omitempty"`

	// relationship
	Relationship *ActivityObject `json:"relationship,omitempty"`

	// replies
	Replies *ActivityObject `json:"replies,omitempty"`

	// result
	Result *ActivityObject `json:"result,omitempty"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// Relationship Properties
	Subject *ActivityObject `json:"subject,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tag
	Tag *ActivityObject `json:"tag,omitempty"`

	// target
	Target *ActivityObject `json:"target,omitempty"`

	// to
	To *ActivityObject `json:"to,omitempty"`

	// total items
	TotalItems int32 `json:"totalItems,omitempty"`

	// type
	Type ActivityObjectType `json:"type,omitempty"`

	// units
	Units string `json:"units,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// url
	URL *ActivityObject `json:"url,omitempty"`

	// width
	Width int32 `json:"width,omitempty"`
}

// Validate validates this activity object
func (m *ActivityObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnyOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBcc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClosed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirst(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenerator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInReplyTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstrument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityObject) validateActor(formats strfmt.Registry) error {

	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateAnyOf(formats strfmt.Registry) error {

	if swag.IsZero(m.AnyOf) { // not required
		return nil
	}

	if m.AnyOf != nil {
		if err := m.AnyOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("anyOf")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateAttachment(formats strfmt.Registry) error {

	if swag.IsZero(m.Attachment) { // not required
		return nil
	}

	if m.Attachment != nil {
		if err := m.Attachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateAttributedTo(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributedTo) { // not required
		return nil
	}

	if m.AttributedTo != nil {
		if err := m.AttributedTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributedTo")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateAudience(formats strfmt.Registry) error {

	if swag.IsZero(m.Audience) { // not required
		return nil
	}

	if m.Audience != nil {
		if err := m.Audience.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audience")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateBcc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bcc) { // not required
		return nil
	}

	if m.Bcc != nil {
		if err := m.Bcc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bcc")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateBto(formats strfmt.Registry) error {

	if swag.IsZero(m.Bto) { // not required
		return nil
	}

	if m.Bto != nil {
		if err := m.Bto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bto")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateCc(formats strfmt.Registry) error {

	if swag.IsZero(m.Cc) { // not required
		return nil
	}

	if m.Cc != nil {
		if err := m.Cc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cc")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateClosed(formats strfmt.Registry) error {

	if swag.IsZero(m.Closed) { // not required
		return nil
	}

	if err := validate.FormatOf("closed", "body", "date-time", m.Closed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateCurrent(formats strfmt.Registry) error {

	if swag.IsZero(m.Current) { // not required
		return nil
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.Deleted) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted", "body", "date-time", m.Deleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if err := validate.FormatOf("duration", "body", "date-time", m.Duration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateFirst(formats strfmt.Registry) error {

	if swag.IsZero(m.First) { // not required
		return nil
	}

	if m.First != nil {
		if err := m.First.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("first")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateFormerType(formats strfmt.Registry) error {

	if swag.IsZero(m.FormerType) { // not required
		return nil
	}

	if err := m.FormerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("formerType")
		}
		return err
	}

	return nil
}

func (m *ActivityObject) validateGenerator(formats strfmt.Registry) error {

	if swag.IsZero(m.Generator) { // not required
		return nil
	}

	if m.Generator != nil {
		if err := m.Generator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generator")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateIcon(formats strfmt.Registry) error {

	if swag.IsZero(m.Icon) { // not required
		return nil
	}

	if m.Icon != nil {
		if err := m.Icon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("icon")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateInReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.InReplyTo) { // not required
		return nil
	}

	if m.InReplyTo != nil {
		if err := m.InReplyTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inReplyTo")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateInstrument(formats strfmt.Registry) error {

	if swag.IsZero(m.Instrument) { // not required
		return nil
	}

	if m.Instrument != nil {
		if err := m.Instrument.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instrument")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActivityObject) validateLast(formats strfmt.Registry) error {

	if swag.IsZero(m.Last) { // not required
		return nil
	}

	if m.Last != nil {
		if err := m.Last.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateObject(formats strfmt.Registry) error {

	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateOneOf(formats strfmt.Registry) error {

	if swag.IsZero(m.OneOf) { // not required
		return nil
	}

	if m.OneOf != nil {
		if err := m.OneOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneOf")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateOrigin(formats strfmt.Registry) error {

	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validatePartOf(formats strfmt.Registry) error {

	if swag.IsZero(m.PartOf) { // not required
		return nil
	}

	if m.PartOf != nil {
		if err := m.PartOf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partOf")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validatePrev(formats strfmt.Registry) error {

	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if m.Prev != nil {
		if err := m.Prev.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prev")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validatePreview(formats strfmt.Registry) error {

	if swag.IsZero(m.Preview) { // not required
		return nil
	}

	if m.Preview != nil {
		if err := m.Preview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preview")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validatePublished(formats strfmt.Registry) error {

	if swag.IsZero(m.Published) { // not required
		return nil
	}

	if err := validate.FormatOf("published", "body", "date-time", m.Published.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateRelationship(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationship) { // not required
		return nil
	}

	if m.Relationship != nil {
		if err := m.Relationship.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationship")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateReplies(formats strfmt.Registry) error {

	if swag.IsZero(m.Replies) { // not required
		return nil
	}

	if m.Replies != nil {
		if err := m.Replies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replies")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateStartTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateSubject(formats strfmt.Registry) error {

	if swag.IsZero(m.Subject) { // not required
		return nil
	}

	if m.Subject != nil {
		if err := m.Subject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateTag(formats strfmt.Registry) error {

	if swag.IsZero(m.Tag) { // not required
		return nil
	}

	if m.Tag != nil {
		if err := m.Tag.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tag")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateTo(formats strfmt.Registry) error {

	if swag.IsZero(m.To) { // not required
		return nil
	}

	if m.To != nil {
		if err := m.To.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("to")
			}
			return err
		}
	}

	return nil
}

func (m *ActivityObject) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *ActivityObject) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActivityObject) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if m.URL != nil {
		if err := m.URL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("url")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityObject) UnmarshalBinary(b []byte) error {
	var res ActivityObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}