package dto

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"

	"gorm.io/gorm"
)

type projectService struct{}

var ProjectService = projectService{}

func (s *projectService) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx).Model(&models.Project{}).Table("projects")
}

type CreateProjectOption struct {
	ProjectName        string   `json:"project_name"`
	ProjectSlug        string   `json:"project_slug"`
	ProjectDescription string   `json:"project_description"`
	ProjectTags        []string `json:"project_tags" gorm:"type:text[]"`
}

type ListProjectOption struct {
	BaseListOption
}

func (s *projectService) Create(ctx *gin.Context, opt *CreateProjectOption) (*models.Project, error) {
	projectId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	nowPtr := new(time.Time)
	*nowPtr = time.Now()

	project := models.Project{
		ProjectName:        opt.ProjectName,
		ProjectDescription: opt.ProjectDescription,
		ProjectTags:        pq.StringArray(opt.ProjectTags),
		ProjectSlug:        opt.ProjectSlug,
		ProjectId:          projectId.String(),
		ProjectCreatedAt:   nowPtr,
	}

	err = s.getBaseDB(ctx).Create(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (s *projectService) GetById(ctx *gin.Context, id *uint) (*models.Project, error) {
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

func (s *projectService) List(ctx *gin.Context, opt *ListProjectOption) ([]*models.Project, error) {
	query := getBaseQuery(ctx, s)
	query = query.Select("*")
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, err
	}
	projects := make([]*models.Project, 0)
	query = opt.BindQueryWithLimit(query)

	err = query.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, err
}
