package dto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
	"strings"
	"time"

	"gorm.io/gorm"
)

type langChainRunRepository struct{}

var LangChainRunRepository = langChainRunRepository{}

// Define the custom type for the enum
type LabelCondition string

// Define the possible values for the enum
const (
	EqualsCondition   LabelCondition = "EQUALS"
	ILikeCondition    LabelCondition = "ILIKE"
	NotILikeCondition LabelCondition = "NOT ILIKE"
)

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
	SortField string           `query:"sort_field"`
	SortOrder types.SortOrders `query:"sort_order"`
}

type LangChainFilterKeys struct {
	ProjectId   string         `json:"project_id" gorm:"column:project_id"`
	RunTimeKeys pq.StringArray `json:"runtime_keys" gorm:"column:runtime_keys"`
	LabelKeys   pq.StringArray `json:"label_keys" gorm:"column:label_keys"`
}

type ListLangChainFilterValuesOption struct {
	BaseListOption
	LabelKey    string         `query:"label_key"`
	Condition   LabelCondition `query:"condition"`
	SearchValue string         `query:"search_value"`
}

func (s *langChainRunRepository) AddOrGetRun(ctx context.Context,
	opt *models.LangChainRuns) (*models.LangChainRuns, error) {

	tx := s.getBaseDB(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	err := tx.Create(&opt).Error
	if err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// No record found, return nil without error
			return s.GetById(ctx, opt.ChainID)
		}
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

	query := getBaseQuery(ctx, s).Select("*").Where("project_id = ?", projectId).Order(
		fmt.Sprintf("%s %s", opt.SortField, opt.SortOrder))
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

func (s *langChainRunRepository) GetRunFilterKeyValues(ctx context.Context,
	opt *ListLangChainFilterValuesOption, projectId string) ([]*string, error) {
	result := make([]*string, 0)
	var total int64
	var labelSplit = strings.Split(opt.LabelKey, ".")
	var conditionValueTemplate string
	var sqlCondition string
	switch conditionValueSwitch := opt.Condition; conditionValueSwitch {
	case EqualsCondition:
		sqlCondition = "="
		conditionValueTemplate = fmt.Sprintf("'%s'", opt.SearchValue)
	case ILikeCondition:
		sqlCondition = string(opt.Condition)

		conditionValueTemplate = fmt.Sprintf("'%%%s%%'", opt.SearchValue)
	case NotILikeCondition:
		sqlCondition = string(opt.Condition)
		conditionValueTemplate = fmt.Sprintf("'%%%s%%'", opt.SearchValue)

	}

	query := getBaseQuery(ctx, s).Select(fmt.Sprintf("DISTINCT(%s->>'%s') as label_values",
		labelSplit[0], labelSplit[1])).Where(
		fmt.Sprintf("project_id = ? and (%s->>'%s') %s %s", labelSplit[0], labelSplit[1], sqlCondition,
			conditionValueTemplate), projectId)

	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	query = opt.BindQueryWithLimit(query)

	err = query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

//
