//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

type PublicUserproperty struct {
	UserpropertyID int64 `db:"userproperty_id,pk" json:"userpropertyId" form:"userpropertyId" validate:"required"`
	UserID         int64 `db:"user_id,use_zero" json:"userId" form:"userId" validate:"required"`
	PropertyID     int64 `db:"property_id,use_zero" json:"propertyId" form:"propertyId" validate:"required"`
}
