package models

import (
	"time"
)

type Project struct {
	ProjectId          string     `json:"project_id"`
	ProjectName        string     `json:"project_name"`
	ProjectSlug        string     `json:"project_slug"`
	ProjectDescription string     `json:"project_description"`
	ProjectTags        []string   `json:"project_tags"`
	ProjectCreatedAt   *time.Time `json:"project_created_at"`
	ProjectDeletedAt   *time.Time `json:"project_deleted_at"`
	ProjectDeleted     bool       `json:"project_deleted"`
}

func (b *Project) GetName() string {
	return b.ProjectName
}
