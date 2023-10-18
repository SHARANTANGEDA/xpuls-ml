package dto

import (
	"context"
	"errors"
	"fmt"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
	"gorm.io/gorm"
)

type ListPromptVersionsOption struct {
	BaseListOption
}

type viewPromptVersionsRepository struct{}
type viewLatestPromptVersionRepository struct{}

var ViewPromptVersionsRepository = viewPromptVersionsRepository{}
var ViewLatestPromptVersionRepository = viewLatestPromptVersionRepository{}

type ListLatestPromptsOption struct {
	BaseListOption
	SortField *string           `query:"sort_field"`
	SortOrder *types.SortOrders `query:"sort_order"`
}

// Define the custom type for the enu

func (s *viewPromptVersionsRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.PromptRegistryPromptVersion{}).Table("view_prompt_registry_versions")
}

func (s *viewLatestPromptVersionRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.PromptRegistryPromptVersion{}).Table(
		"view_prompt_registry_latest_version")
}

func (s *viewPromptVersionsRepository) GetLatestPromptById(ctx context.Context, promptId string) (*models.PromptRegistryPrompt, error) {
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

func (s *viewPromptVersionsRepository) ListPromptVersions(
	ctx context.Context, projectId string, promptId string,
	opt *ListPromptVersionsOption) ([]*models.ViewPromptRegistryVersion, error) {
	var total int64

	query := getBaseQuery(ctx, s).Select("*").Where("project_id = ? and prompt_id = ?", projectId,
		promptId).Order(
		"prompt_version_created_at desc")
	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}

	promptRegistriesPromptVersions := make([]*models.ViewPromptRegistryVersion, 0)
	query = opt.BindQueryWithLimit(query)

	err = query.Find(&promptRegistriesPromptVersions).Error
	if err != nil {
		return nil, err
	}
	return promptRegistriesPromptVersions, err
}

func (s *viewPromptVersionsRepository) LatestPromptVersion(
	ctx context.Context, projectId string, promptId string,
) (*models.ViewPromptRegistryVersion, error) {
	var total int64

	query := getBaseQuery(ctx, s).Select("*").Where("project_id = ? and prompt_id = ?", projectId,
		promptId).Order(
		"prompt_version_created_at desc").Limit(1)
	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}
	var promptRegistriesPromptVersion *models.ViewPromptRegistryVersion = nil

	err = query.Find(&promptRegistriesPromptVersion).Error
	if err != nil {
		return nil, err
	}
	return promptRegistriesPromptVersion, err
}

func (s *viewLatestPromptVersionRepository) ListLatestPromptVersions(
	ctx context.Context, projectId string, opt *ListLatestPromptsOption,
) ([]*models.ViewPromptRegistryVersion, error) {
	var total int64

	var sortField string
	var sortOrder types.SortOrders

	defaultSortField := "prompt_version_created_at"
	defaultSortOrder := types.DESC
	if opt.SortField == nil {
		sortField = defaultSortField
	} else {
		sortField = *opt.SortField
	}

	if opt.SortOrder == nil {
		sortOrder = defaultSortOrder
	} else {
		sortOrder = *opt.SortOrder
	}

	query := getBaseQuery(ctx, s).Select("*").Where("project_id = ?", projectId).Order(
		fmt.Sprintf("%s %s", sortField, sortOrder))
	err := query.Count(&total).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// No record found, return nil without error
			return nil, nil
		}
		return nil, err
	}
	latestPrompts := make([]*models.ViewPromptRegistryVersion, 0)

	err = query.Find(&latestPrompts).Error
	query = opt.BindQueryWithLimit(query)

	if err != nil {
		return nil, err
	}
	return latestPrompts, err
}
