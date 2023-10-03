package dto

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"

	"gorm.io/gorm"
)

type langChainRunStepsRepository struct{}

var LangChainRunStepsRepository = langChainRunStepsRepository{}

func (s *langChainRunStepsRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.LangChainRunSteps{}).Table("langchain_run_steps")
}

type ListLangChainRunStepsOption struct {
	BaseListOption
	ProjectId string `query:"project_id"`
	ChainId   string `query:"chain_id"`
}

func (s *langChainRunStepsRepository) AddNewRunStep(ctx *gin.Context, opt *models.LangChainRunSteps) (*models.LangChainRunSteps, error) {
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

func (s *langChainRunStepsRepository) PatchRunStep(ctx *gin.Context,
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

func (s *langChainRunStepsRepository) GetById(ctx context.Context, runStepId string) (*models.LangChainRunSteps, error) {
	var langChainRunStep models.LangChainRunSteps
	err := getBaseQuery(ctx, s).Where("run_step_id = ?", runStepId).First(&langChainRunStep).Error
	if err != nil {
		return nil, err
	}

	return &langChainRunStep, nil
}
func (s *langChainRunStepsRepository) GetRunsByChainId(ctx context.Context, opt *ListLangChainRunStepsOption) ([]*models.LangChainRunSteps, error) {
	langChainRunSteps := make([]*models.LangChainRunSteps, 0)
	var total int64
	query := getBaseQuery(ctx, s).Where("chain_id = ?", opt.ChainId)

	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	query = opt.BindQueryWithLimit(query)

	err = query.Find(&langChainRunSteps).Error
	if err != nil {
		return nil, err
	}

	return langChainRunSteps, err
}
