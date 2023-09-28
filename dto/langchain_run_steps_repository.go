package dto

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"

	"gorm.io/gorm"
)

type langChainRunStepsService struct{}

var LangChainRunStepsService = langChainRunStepsService{}

func (s *langChainRunStepsService) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.LangChainRunSteps{}).Table("langchain_run_steps")
}

type ListLangChainRunStepsOption struct {
	BaseListOption
}

func (s *langChainRunStepsService) AddNewRunStep(ctx *gin.Context, opt *models.LangChainRunSteps) (*models.LangChainRunSteps, error) {
	nowPtr := new(time.Time)
	*nowPtr = time.Now()
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

func (s *langChainRunStepsService) PatchRunStep(ctx *gin.Context,
	opt *models.LangChainRunSteps) (*models.LangChainRunSteps, error) {
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Where("run_step_id = ?", opt.RunStepID).Updates(opt).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Return the updated record
	return opt, nil
}

func (s *langChainRunStepsService) GetById(ctx context.Context, runStepId string) (*models.LangChainRunSteps, error) {
	var langChainRunStep models.LangChainRunSteps
	err := getBaseQuery(ctx, s).Where("run_step_id = ?", runStepId).First(&langChainRunStep).Error
	if err != nil {
		return nil, err
	}

	return &langChainRunStep, nil
}
