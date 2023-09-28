package models

import (
	"encoding/json"
	"time"
)

type LangChainRuns struct {
	ChainID            string          `json:"chain_id" gorm:"primary_key;type:varchar(200)"`
	ProjectID          string          `json:"project_id" gorm:"type:varchar(200);not null"`
	ModelInfo          json.RawMessage `json:"model_info" gorm:"type:jsonb"`
	Labels             json.RawMessage `json:"labels" gorm:"type:jsonb;default:'{}'"`
	Runtime            json.RawMessage `json:"runtime" gorm:"type:jsonb"`
	ChainTrackedAt     *time.Time      `json:"chain_tracked_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	FirstStepStartTime *time.Time      `json:"first_step_start_time" gorm:"type:timestamp"`
	LastStepEndTime    *time.Time      `json:"last_step_end_time" gorm:"type:timestamp"`
	TotalTokens        *int            `json:"total_tokens" gorm:"type:integer"`
	PromptTokens       *int            `json:"prompt_tokens" gorm:"type:integer"`
	CompletionTokens   *int            `json:"completion_tokens" gorm:"type:integer"`
}
