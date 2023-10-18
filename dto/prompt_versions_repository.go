package dto

import (
	"context"
	"github.com/google/uuid"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
	"gorm.io/gorm"
)

type promptVersionsRepository struct{}

var PromptVersionsRepository = promptVersionsRepository{}

func (s *promptVersionsRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.PromptRegistryPromptVersion{}).Table("prompt_registry_prompt_versions")
}

func (s *promptVersionsRepository) AddNewPromptVersion(ctx context.Context,
	opt *types.AddPromptVersionRequest) (*string, error) {
	promptVersionId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	promptVersionIdString := promptVersionId.String()
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	promptVersions := models.PromptRegistryPromptVersion{
		PromptID:        opt.PromptID,
		PromptContent:   opt.PromptContent,
		PromptVersionID: promptVersionId.String(),
	}
	err = tx.Create(&promptVersions).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &promptVersionIdString, err
}
