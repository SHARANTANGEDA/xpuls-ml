package models

import (
	"time"
)

type PromptRegistryPromptVersion struct {
	PromptVersionID        string    `json:"prompt_version_id" gorm:"type:varchar(200);primaryKey"`
	PromptContent          string    `json:"prompt_content" gorm:"type:text;not null"`
	PromptVersionCreatedAt time.Time `json:"prompt_version_created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP AT TIME ZONE 'UTC';not null"`
	PromptID               string    `json:"prompt_id" gorm:"type:varchar(200);not null;foreignKey:PromptID;references:prompt_registry_prompts"`
}
