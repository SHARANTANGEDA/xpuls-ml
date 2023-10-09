package dto

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/models"
	"gorm.io/gorm"
)

type llmTracingQueueRepository struct{}

var LLMTracingQueueRepository = llmTracingQueueRepository{}

func (s *llmTracingQueueRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.LLMTracingQueue{}).Table("llm_tracing_queue")
}

func (s *llmTracingQueueRepository) AddItemToQueue(ctx *gin.Context,
	opt *models.LLMTracingQueue) (*models.LLMTracingQueue, error) {

	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Create(&opt).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return opt, err
}

func (s *llmTracingQueueRepository) GetUnProcessedItems(ctx context.Context) ([]*models.LLMTracingQueue, error) {
	llmTracingQueue := make([]*models.LLMTracingQueue, 0)
	var total int64
	query := getBaseQuery(ctx, s).Where("queue_item_processed = false")

	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	err = query.Find(&llmTracingQueue).Error
	if err != nil {
		return nil, err
	}

	return llmTracingQueue, err
}

func (s *llmTracingQueueRepository) MarkItemAsProcessed(ctx context.Context,
	queueItem *models.LLMTracingQueue) (*models.LLMTracingQueue, error) {
	// Update the record
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Where("queue_item_id = ?", queueItem.QueueItemID).Updates(&queueItem).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	// Return the updated record
	return queueItem, nil
}

func (s *llmTracingQueueRepository) DeleteProcessedItems(ctx context.Context) error {
	// Start a new transaction
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete the processed records
	if err := tx.Where("queue_item_processed = ?", true).Delete(&models.LLMTracingQueue{}).Error; err != nil {
		tx.Rollback() // Important: If there's an error, rollback the transaction
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
