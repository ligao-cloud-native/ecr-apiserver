package models

type ResourceGroup struct {
	ResourceGroupID   int    `gorm:"primary_key" json:"resource_group_id,omitempty"`
	ResourceGroupName string `gorm:"size:255" json:"resource_group_name"`
	ResourceGroupUUID string `gorm:"size:255" json:"resource_group_uuid"`
	ResourceGroupENV  string `gorm:"size:255" json:"resource_group_env"`
}
