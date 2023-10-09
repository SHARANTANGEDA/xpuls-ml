package crons

import (
	"context"
	"fmt"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/services"
	"github.com/xpuls-com/xpuls-ml/types"
)

func processQueueItem(ctx context.Context, item *models.LLMTracingQueue) error {
	if item.QueueItemType == types.LANGCHAIN_RUN {
		switch processingCode := item.QueueProcessingCode; processingCode {
		case types.LANGCHAIN_RUN_BEGIN:
			err := services.LangChainRunsServiceV2.TrackRunFromQueue(ctx, item.Params, item.Data)
			if err != nil {
				return err
			}
		case types.LANGCHAIN_RUN_END:
			err := services.LangChainRunsServiceV2.PatchRunFromQueue(ctx, item.Params, item.Data)
			if err != nil {
				return err
			}

		}
	}
	item.QueueItemProcessed = true
	_, err := dto.LLMTracingQueueRepository.MarkItemAsProcessed(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func AddProcessLLMTracingQueue(ctx context.Context) error {

	items, err := dto.LLMTracingQueueRepository.GetUnProcessedItems(ctx)
	if err != nil {
		return err
	}
	// Iterate over items and process them
	for _, item := range items {
		if err := processQueueItem(ctx, item); err != nil {
			fmt.Printf("Unable to parse queue item: %s. Will attempt it again in next run", item.QueueItemID)
			continue
		}
	}

	return nil
}
