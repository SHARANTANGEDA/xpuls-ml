package dto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"

	"gorm.io/gorm"
)

type langChainRunRepository struct{}

var LangChainRunRepository = langChainRunRepository{}

func (s *langChainRunRepository) getBaseDB(ctx context.Context) *gorm.DB {
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

type ListLangChainRunOption struct {
	BaseListOption
}

type LangChainFilterKeys struct {
	ProjectId   string         `json:"project_id" gorm:"column:project_id"`
	RunTimeKeys pq.StringArray `json:"runtime_keys" gorm:"column:runtime_keys"`
	LabelKeys   pq.StringArray `json:"label_keys" gorm:"column:label_keys"`
}

func (s *langChainRunRepository) AddNewRun(ctx context.Context, opt *models.LangChainRuns) (*models.LangChainRuns, error) {

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

func (s *langChainRunRepository) UpdateRun(ctx context.Context, opt *models.LangChainRuns) (*models.LangChainRuns, error) {
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

func (s *langChainRunRepository) GetById(ctx context.Context, runId string) (*models.LangChainRuns, error) {
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

func (s *langChainRunRepository) GetRunsInProject(ctx context.Context, opt *ListLangChainRunOption, projectId string) ([]*models.LangChainRuns, error) {
	var total int64

	query := getBaseQuery(ctx, s).Select("*").Where("project_id = ?", projectId)
	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	fmt.Println(total, projectId)
	langChainRuns := make([]*models.LangChainRuns, 0)
	query = opt.BindQueryWithLimit(query)

	err = query.Find(&langChainRuns).Error
	if err != nil {
		return nil, err
	}
	return langChainRuns, err
}

func (s *langChainRunRepository) GetRunFilterKeys(ctx context.Context, opt *ListLangChainRunOption,
	projectId string) (*LangChainFilterKeys, error) {
	var result *LangChainFilterKeys
	err := mustGetSession(ctx).Model(&LangChainFilterKeys{}).Table(
		"view_langchain_run_filter_keys").Select("*").Where(
		"project_id = ?", projectId).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

//
