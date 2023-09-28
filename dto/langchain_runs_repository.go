package dto

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"

	"gorm.io/gorm"
)

type langChainRunService struct{}

var LangChainRunService = langChainRunService{}

func (s *langChainRunService) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.LangChainRuns{}).Table("langchain_runs")
}

type TrackLangChainRunOption struct {
	ChainId            string          `json:"chain_id"`
	ProjectID          string          `json:"project_id"`
	ModelInfo          json.RawMessage `json:"model_info"`
	Labels             json.RawMessage `json:"labels"`
	Runtime            json.RawMessage `json:"runtime"`
	FirstStepStartTime *time.Time      `json:"first_step_start_time"`
	LastStepEndTime    *time.Time      `json:"last_step_end_time"`
	TotalTokens        *int            `json:"total_tokens"`
	PromptTokens       *int            `json:"prompt_tokens"`
	CompletionTokens   *int            `json:"completion_tokens"`
}

type ListLangChainOption struct {
	BaseListOption
}

func (s *langChainRunService) AddNewRun(ctx context.Context, opt *models.LangChainRuns) (*models.LangChainRuns, error) {
	//nowPtr := new(time.Time)
	//*nowPtr = time.Now()

	//chainRun := models.LangChainRuns{
	//	ChainID:            opt.ChainID,
	//	ProjectID:          opt.ProjectID,
	//	ModelInfo:          opt.ModelInfo,
	//	Labels:             opt.Labels,
	//	Runtime:            opt.Runtime,
	//	FirstStepStartTime: opt.FirstStepStartTime,
	//	LastStepEndTime:    opt.LastStepEndTime,
	//	TotalTokens:        opt.TotalTokens,
	//	PromptTokens:       opt.PromptTokens,
	//	CompletionTokens:   opt.CompletionTokens,
	//	ChainTrackedAt:     nowPtr,
	//}
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

func (s *langChainRunService) UpdateRun(ctx context.Context, opt *models.LangChainRuns) (*models.LangChainRuns, error) {
	// Update the record
	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Where("chain_id = ?", opt.ChainID).Updates(opt).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	// Return the updated record
	return opt, nil
}

func (s *langChainRunService) GetById(ctx context.Context, runId string) (*models.LangChainRuns, error) {
	var langChainRun models.LangChainRuns
	err := getBaseQuery(ctx, s).Where("chain_id = ?", runId).First(&langChainRun).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}
	//if project.ID == 0 {
	//	return nil, constants.ErrNotFound
	//}
	return &langChainRun, nil
}

//
//func (s *projectService) List(ctx *gin.Context, opt *ListProjectOption) ([]*models.Project, error) {
//	query := getBaseQuery(ctx, s)
//	query = query.Select("*")
//	var total int64
//	err := query.Count(&total).Error
//	if err != nil {
//		return nil, err
//	}
//	projects := make([]*models.Project, 0)
//	query = opt.BindQueryWithLimit(query)
//
//	err = query.Find(&projects).Error
//	if err != nil {
//		return nil, err
//	}
//	return projects, err
//}
