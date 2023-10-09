package models

import (
	"encoding/json"
	"github.com/xpuls-com/xpuls-ml/types"
	"time"
)

type LLMTracingQueue struct {
	QueueItemID         string                               `json:"queue_item_id" db:"queue_item_id" gorm:"type:varchar(200);primary_key"`
	QueueItemType       types.LLMTracingQueueItemTypes       `json:"queue_item_type" db:"queue_item_type" gorm:"type:varchar(200);default:'LANGCHAIN_RUN';not null"`
	ItemStoredAt        *time.Time                           `json:"item_stored_at" db:"item_stored_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	Params              json.RawMessage                      `json:"params" db:"params" gorm:"type:jsonb;default:'{}';not null"`
	Data                json.RawMessage                      `json:"data" db:"data" gorm:"type:jsonb;default:'{}';not null"`
	QueueProcessingCode types.LLMTracingQueueProcessingTypes `json:"queue_processing_code" db:"queue_processing_code" gorm:"type:varchar(200)"`
	QueueItemProcessed  bool                                 `json:"queue_item_processed" db:"queue_item_processed" gorm:"type:boolean;default:false;not null"`
}
