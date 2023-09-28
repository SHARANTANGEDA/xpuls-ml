package models

import (
	"encoding/json"
	"github.com/lib/pq"
	"time"
)

type LangChainRunSteps struct {
	RunStepID                    string          `json:"run_step_id" gorm:"primary_key;type:varchar(200)"`
	RunName                      string          `json:"run_name" gorm:"type:varchar(200);not null"`
	RunType                      string          `json:"run_type" gorm:"type:varchar(200);not null"`
	ExecutionOrder               int             `json:"execution_order" gorm:"type:integer;not null"`
	PromptTemplate               string          `json:"prompt_template" gorm:"type:text"`
	PromptContent                string          `json:"prompt_content" gorm:"type:text"`
	PromptTemplateInputVariables pq.StringArray  `json:"prompt_template_input_variables" gorm:"type:varchar(250)[]"`
	PromptInput                  string          `json:"prompt_input" gorm:"type:text"`
	PromptChatHistory            string          `json:"prompt_chat_history" gorm:"type:text"`
	PromptAgentScratchpad        string          `json:"prompt_agent_scratchpad" gorm:"type:text"`
	PromptOutput                 string          `json:"prompt_output" gorm:"type:text"`
	EventStartTime               *time.Time      `json:"event_start_time" gorm:"type:timestamp;not null"`
	EventEndTime                 *time.Time      `json:"event_end_time" gorm:"type:timestamp"`
	ParentStepID                 string          `json:"parent_step_id" gorm:"type:varchar(200)"`
	TokenUsage                   json.RawMessage `json:"token_usage" gorm:"type:jsonb"`
	BeginJSON                    json.RawMessage `json:"begin_json" gorm:"type:jsonb"`
	EndJSON                      json.RawMessage `json:"end_json" gorm:"type:jsonb"`
	TrackedAt                    time.Time       `json:"tracked_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	ChainID                      string          `json:"chain_id" gorm:"type:varchar(200);not null"`
}
