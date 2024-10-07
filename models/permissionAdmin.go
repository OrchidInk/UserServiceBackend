package models

type PermissionRequest struct {
	AdminID   int32 `json:"admin_id" validate:"required"`
	CanCreate bool  `json:"can_create"`
	CanRead   bool  `json:"can_read"`
	CanUpdate bool  `json:"can_update"`
	CanDelete bool  `json:"can_delete"`
}
