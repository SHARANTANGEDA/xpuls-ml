package models

import (
	"github.com/lib/pq"
	"time"
)

type Project struct {
	ProjectId          string         `json:"project_id"`
	ProjectName        string         `json:"project_name"`
	ProjectSlug        string         `json:"project_slug"`
	ProjectDescription string         `json:"project_description"`
	ProjectTags        pq.StringArray `json:"project_tags" gorm:"type:text[]"`
	ProjectCreatedAt   *time.Time     `json:"project_created_at"`
	ProjectDeletedAt   *time.Time     `json:"project_deleted_at"`
	ProjectDeleted     bool           `json:"project_deleted"`
}

func (b *Project) GetName() string {
	return b.ProjectName
}
