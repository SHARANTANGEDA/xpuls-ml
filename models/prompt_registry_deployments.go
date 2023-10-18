package models

import (
	"encoding/json"
	"time"
)

type PromptRegistryDeployment struct {
	PromptDeploymentID      string          `json:"prompt_deployment_id" gorm:"type:varchar(200);primaryKey"`
	DeploymentEnv           string          `json:"deployment_env" gorm:"type:varchar(250);not null;foreignKey:DeploymentEnv;references:compute_environments"`
	DeploymentMode          string          `json:"deployment_mode" gorm:"type:varchar(200);default:'ROLLOUT';not null"`
	DeploymentProperties    json.RawMessage `json:"deployment_properties" gorm:"type:jsonb;default:'{}';not null"`
	DeploymentCreatedAt     time.Time       `json:"deployment_created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP AT TIME ZONE 'UTC';not null"`
	DeploymentLastUpdatedAt *time.Time      `json:"deployment_last_updated_at" gorm:"type:timestamp"`
	PromptID                string          `json:"prompt_id" gorm:"type:varchar(200);not null;foreignKey:PromptID;references:prompt_registry_prompts"`
	DeploymentActive        bool            `json:"deployment_active" gorm:"type:boolean;default:true;not null"`
	DeploymentDeleted       bool            `json:"deployment_deleted" gorm:"type:boolean;default:false;not null"`
	DeploymentDeletedAt     *time.Time      `json:"deployment_deleted_at" gorm:"type:timestamp"`
}
