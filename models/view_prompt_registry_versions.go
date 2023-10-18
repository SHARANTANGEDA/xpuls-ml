package models

import (
	"time"
)

type ViewPromptRegistryVersion struct {
	PromptVersionID        string     `json:"prompt_version_id" gorm:"column:prompt_version_id"`
	PromptContent          string     `json:"prompt_content" gorm:"column:prompt_content"`
	PromptVersionCreatedAt time.Time  `json:"prompt_version_created_at" gorm:"column:prompt_version_created_at"`
	PromptID               string     `json:"prompt_id" gorm:"column:prompt_id"`
	PromptTag              string     `json:"prompt_tag" gorm:"column:prompt_tag"`
	PromptName             string     `json:"prompt_name" gorm:"column:prompt_name"`
	ProjectID              string     `json:"project_id" gorm:"column:project_id"`
	PromptCreatedAt        time.Time  `json:"prompt_created_at" gorm:"column:prompt_created_at"`
	PromptDeleted          bool       `json:"prompt_deleted" gorm:"column:prompt_deleted"`
	PromptDeletedAt        *time.Time `json:"prompt_deleted_at" gorm:"column:prompt_deleted_at"`
}

// Ensure you set the correct table name if using GORM
//func (ViewPromptRegistryVersion) TableName() string {
//	return "view_prompt_registry_versions"
//}
