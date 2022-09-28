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

// ActivityObjectType - Collection: CollectionTypes
//   - Application: Actor Types
//   - Article: Objects Types
//   - Accept: Activity Types
//   - Workspace: Pydio Types
//
// swagger:model activityObjectType
type ActivityObjectType string

func NewActivityObjectType(value ActivityObjectType) *ActivityObjectType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ActivityObjectType.
func (m ActivityObjectType) Pointer() *ActivityObjectType {
	return &m
}

const (

	// ActivityObjectTypeBaseObject captures enum value "BaseObject"
	ActivityObjectTypeBaseObject ActivityObjectType = "BaseObject"

	// ActivityObjectTypeActivity captures enum value "Activity"
	ActivityObjectTypeActivity ActivityObjectType = "Activity"

	// ActivityObjectTypeLink captures enum value "Link"
	ActivityObjectTypeLink ActivityObjectType = "Link"

	// ActivityObjectTypeMention captures enum value "Mention"
	ActivityObjectTypeMention ActivityObjectType = "Mention"

	// ActivityObjectTypeCollection captures enum value "Collection"
	ActivityObjectTypeCollection ActivityObjectType = "Collection"

	// ActivityObjectTypeOrderedCollection captures enum value "OrderedCollection"
	ActivityObjectTypeOrderedCollection ActivityObjectType = "OrderedCollection"

	// ActivityObjectTypeCollectionPage captures enum value "CollectionPage"
	ActivityObjectTypeCollectionPage ActivityObjectType = "CollectionPage"

	// ActivityObjectTypeOrderedCollectionPage captures enum value "OrderedCollectionPage"
	ActivityObjectTypeOrderedCollectionPage ActivityObjectType = "OrderedCollectionPage"

	// ActivityObjectTypeApplication captures enum value "Application"
	ActivityObjectTypeApplication ActivityObjectType = "Application"

	// ActivityObjectTypeGroup captures enum value "Group"
	ActivityObjectTypeGroup ActivityObjectType = "Group"

	// ActivityObjectTypeOrganization captures enum value "Organization"
	ActivityObjectTypeOrganization ActivityObjectType = "Organization"

	// ActivityObjectTypePerson captures enum value "Person"
	ActivityObjectTypePerson ActivityObjectType = "Person"

	// ActivityObjectTypeService captures enum value "Service"
	ActivityObjectTypeService ActivityObjectType = "Service"

	// ActivityObjectTypeArticle captures enum value "Article"
	ActivityObjectTypeArticle ActivityObjectType = "Article"

	// ActivityObjectTypeAudio captures enum value "Audio"
	ActivityObjectTypeAudio ActivityObjectType = "Audio"

	// ActivityObjectTypeDocument captures enum value "Document"
	ActivityObjectTypeDocument ActivityObjectType = "Document"

	// ActivityObjectTypeEvent captures enum value "Event"
	ActivityObjectTypeEvent ActivityObjectType = "Event"

	// ActivityObjectTypeImage captures enum value "Image"
	ActivityObjectTypeImage ActivityObjectType = "Image"

	// ActivityObjectTypeNote captures enum value "Note"
	ActivityObjectTypeNote ActivityObjectType = "Note"

	// ActivityObjectTypePage captures enum value "Page"
	ActivityObjectTypePage ActivityObjectType = "Page"

	// ActivityObjectTypePlace captures enum value "Place"
	ActivityObjectTypePlace ActivityObjectType = "Place"

	// ActivityObjectTypeProfile captures enum value "Profile"
	ActivityObjectTypeProfile ActivityObjectType = "Profile"

	// ActivityObjectTypeRelationship captures enum value "Relationship"
	ActivityObjectTypeRelationship ActivityObjectType = "Relationship"

	// ActivityObjectTypeTombstone captures enum value "Tombstone"
	ActivityObjectTypeTombstone ActivityObjectType = "Tombstone"

	// ActivityObjectTypeVideo captures enum value "Video"
	ActivityObjectTypeVideo ActivityObjectType = "Video"

	// ActivityObjectTypeAccept captures enum value "Accept"
	ActivityObjectTypeAccept ActivityObjectType = "Accept"

	// ActivityObjectTypeAdd captures enum value "Add"
	ActivityObjectTypeAdd ActivityObjectType = "Add"

	// ActivityObjectTypeAnnounce captures enum value "Announce"
	ActivityObjectTypeAnnounce ActivityObjectType = "Announce"

	// ActivityObjectTypeArrive captures enum value "Arrive"
	ActivityObjectTypeArrive ActivityObjectType = "Arrive"

	// ActivityObjectTypeBlock captures enum value "Block"
	ActivityObjectTypeBlock ActivityObjectType = "Block"

	// ActivityObjectTypeCreate captures enum value "Create"
	ActivityObjectTypeCreate ActivityObjectType = "Create"

	// ActivityObjectTypeDelete captures enum value "Delete"
	ActivityObjectTypeDelete ActivityObjectType = "Delete"

	// ActivityObjectTypeDislike captures enum value "Dislike"
	ActivityObjectTypeDislike ActivityObjectType = "Dislike"

	// ActivityObjectTypeFlag captures enum value "Flag"
	ActivityObjectTypeFlag ActivityObjectType = "Flag"

	// ActivityObjectTypeFollow captures enum value "Follow"
	ActivityObjectTypeFollow ActivityObjectType = "Follow"

	// ActivityObjectTypeIgnore captures enum value "Ignore"
	ActivityObjectTypeIgnore ActivityObjectType = "Ignore"

	// ActivityObjectTypeInvite captures enum value "Invite"
	ActivityObjectTypeInvite ActivityObjectType = "Invite"

	// ActivityObjectTypeJoin captures enum value "Join"
	ActivityObjectTypeJoin ActivityObjectType = "Join"

	// ActivityObjectTypeLeave captures enum value "Leave"
	ActivityObjectTypeLeave ActivityObjectType = "Leave"

	// ActivityObjectTypeLike captures enum value "Like"
	ActivityObjectTypeLike ActivityObjectType = "Like"

	// ActivityObjectTypeListen captures enum value "Listen"
	ActivityObjectTypeListen ActivityObjectType = "Listen"

	// ActivityObjectTypeMove captures enum value "Move"
	ActivityObjectTypeMove ActivityObjectType = "Move"

	// ActivityObjectTypeOffer captures enum value "Offer"
	ActivityObjectTypeOffer ActivityObjectType = "Offer"

	// ActivityObjectTypeQuestion captures enum value "Question"
	ActivityObjectTypeQuestion ActivityObjectType = "Question"

	// ActivityObjectTypeReject captures enum value "Reject"
	ActivityObjectTypeReject ActivityObjectType = "Reject"

	// ActivityObjectTypeRead captures enum value "Read"
	ActivityObjectTypeRead ActivityObjectType = "Read"

	// ActivityObjectTypeRemove captures enum value "Remove"
	ActivityObjectTypeRemove ActivityObjectType = "Remove"

	// ActivityObjectTypeTentativeReject captures enum value "TentativeReject"
	ActivityObjectTypeTentativeReject ActivityObjectType = "TentativeReject"

	// ActivityObjectTypeTentativeAccept captures enum value "TentativeAccept"
	ActivityObjectTypeTentativeAccept ActivityObjectType = "TentativeAccept"

	// ActivityObjectTypeTravel captures enum value "Travel"
	ActivityObjectTypeTravel ActivityObjectType = "Travel"

	// ActivityObjectTypeUndo captures enum value "Undo"
	ActivityObjectTypeUndo ActivityObjectType = "Undo"

	// ActivityObjectTypeUpdate captures enum value "Update"
	ActivityObjectTypeUpdate ActivityObjectType = "Update"

	// ActivityObjectTypeUpdateComment captures enum value "UpdateComment"
	ActivityObjectTypeUpdateComment ActivityObjectType = "UpdateComment"

	// ActivityObjectTypeUpdateMeta captures enum value "UpdateMeta"
	ActivityObjectTypeUpdateMeta ActivityObjectType = "UpdateMeta"

	// ActivityObjectTypeView captures enum value "View"
	ActivityObjectTypeView ActivityObjectType = "View"

	// ActivityObjectTypeWorkspace captures enum value "Workspace"
	ActivityObjectTypeWorkspace ActivityObjectType = "Workspace"

	// ActivityObjectTypeDigest captures enum value "Digest"
	ActivityObjectTypeDigest ActivityObjectType = "Digest"

	// ActivityObjectTypeFolder captures enum value "Folder"
	ActivityObjectTypeFolder ActivityObjectType = "Folder"

	// ActivityObjectTypeCell captures enum value "Cell"
	ActivityObjectTypeCell ActivityObjectType = "Cell"

	// ActivityObjectTypeShare captures enum value "Share"
	ActivityObjectTypeShare ActivityObjectType = "Share"
)

// for schema
var activityObjectTypeEnum []interface{}

func init() {
	var res []ActivityObjectType
	if err := json.Unmarshal([]byte(`["BaseObject","Activity","Link","Mention","Collection","OrderedCollection","CollectionPage","OrderedCollectionPage","Application","Group","Organization","Person","Service","Article","Audio","Document","Event","Image","Note","Page","Place","Profile","Relationship","Tombstone","Video","Accept","Add","Announce","Arrive","Block","Create","Delete","Dislike","Flag","Follow","Ignore","Invite","Join","Leave","Like","Listen","Move","Offer","Question","Reject","Read","Remove","TentativeReject","TentativeAccept","Travel","Undo","Update","UpdateComment","UpdateMeta","View","Workspace","Digest","Folder","Cell","Share"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activityObjectTypeEnum = append(activityObjectTypeEnum, v)
	}
}

func (m ActivityObjectType) validateActivityObjectTypeEnum(path, location string, value ActivityObjectType) error {
	if err := validate.EnumCase(path, location, value, activityObjectTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this activity object type
func (m ActivityObjectType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActivityObjectTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this activity object type based on context it is used
func (m ActivityObjectType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}