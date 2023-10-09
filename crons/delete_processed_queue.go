package crons

import (
	"context"
	"github.com/xpuls-com/xpuls-ml/dto"
)

func DeleteProcessedQueueItems(ctx context.Context) error {

	err := dto.LLMTracingQueueRepository.DeleteProcessedItems(ctx)
	if err != nil {
		return err
	}

	return nil
}
