package dto

import (
	"context"
	"errors"
	"github.com/xpuls-com/xpuls-ml/common/utils"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
	"gorm.io/gorm"
)

type promptsRepository struct{}

var PromptsRepository = promptsRepository{}

// Define the custom type for the enu

func (s *promptsRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.PromptRegistryPrompt{}).Table("prompt_registry_prompts")
}

type ListPromptOptions struct {
	BaseListOption
	SortField string           `query:"sort_field"`
	SortOrder types.SortOrders `query:"sort_order"`
}

func (s *promptsRepository) AddNewPrompt(ctx context.Context,
	opt *types.AddPromptRequest) (*models.PromptRegistryPrompt, error) {

	promptId, err := utils.NewUUIDWithPrefix(types.PROMPT_ID_PREFIX)
	if err != nil {
		return nil, err
	}

	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	prompt := models.PromptRegistryPrompt{
		PromptName: opt.PromptName,
		ProjectID:  opt.ProjectId,
		PromptID:   promptId,
	}
	err = tx.Create(&prompt).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &prompt, err
}

func (s *promptsRepository) UpdatePromptName(ctx context.Context, promptId string, promptName string) error {
	// Update the record
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}
	err := tx.Where("prompt_id = ? and prompt_deleted=false", promptId).UpdateColumn(
		"prompt_name", promptName).Error
	if err != nil {
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	// Return the updated record
	return nil
}

func (s *promptsRepository) GetPromptById(ctx context.Context, promptId string) (*models.PromptRegistryPrompt, error) {
	var promptRegistryPrompt models.PromptRegistryPrompt
	err := getBaseQuery(ctx, s).Where("promptId = ?", promptId).First(&promptRegistryPrompt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	return &promptRegistryPrompt, nil
}
