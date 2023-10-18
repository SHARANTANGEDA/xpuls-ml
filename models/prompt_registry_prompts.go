package models

import "time"

type PromptRegistryPrompt struct {
	PromptID        string     `json:"prompt_id" gorm:"type:varchar(200);primaryKey"`
	PromptName      string     `json:"prompt_name" gorm:"type:varchar(200);not null"`
	ProjectID       string     `json:"project_id" gorm:"type:varchar(200);not null;foreignKey:ProjectID;references:projects"`
	PromptCreatedAt time.Time  `json:"prompt_created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP AT TIME ZONE 'UTC';not null"`
	PromptDeleted   bool       `json:"prompt_deleted" gorm:"type:boolean;default:false;not null"`
	PromptDeletedAt *time.Time `json:"prompt_deleted_at" gorm:"type:timestamp"`
}
