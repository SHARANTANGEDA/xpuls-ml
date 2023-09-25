package dto

import (
	"context"
	"github.com/google/uuid"
	"github.com/xpuls-com/xpuls-ml/models"

	"gorm.io/gorm"
)

type projectService struct{}

var ProjectService = projectService{}

func (s *projectService) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.Project{})
}

type CreateProjectOption struct {
	ProjectName        string
	ProjectSlug        string
	ProjectDescription string
	ProjectTags        []string
}

type ListProjectOption struct {
	BaseListOption
}

func (*projectService) Create(ctx context.Context, opt CreateProjectOption) (*models.Project, error) {
	projectId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	project := models.Project{
		ProjectName:        opt.ProjectName,
		ProjectDescription: opt.ProjectDescription,
		ProjectTags:        opt.ProjectTags,
		ProjectSlug:        opt.ProjectSlug,
		ProjectId:          projectId.String(),
	}
	err = mustGetSession(ctx).Create(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (s *projectService) GetById(ctx context.Context, id uint) (*models.Project, error) {
	var project models.Project
	err := getBaseQuery(ctx, s).Where("project_id = ? and project_deleted = false", id).First(&project).Error
	if err != nil {
		return nil, err
	}
	//if project.ID == 0 {
	//	return nil, constants.ErrNotFound
	//}
	return &project, nil
}

func (s *projectService) List(ctx context.Context, opt ListProjectOption) ([]*models.Project, uint, error) {
	query := getBaseQuery(ctx, s)
	query = query.Select("*")
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	projects := make([]*models.Project, 0)
	query = opt.BindQueryWithLimit(query)

	err = query.Find(&projects).Error
	if err != nil {
		return nil, 0, err
	}
	return projects, uint(total), err
}
