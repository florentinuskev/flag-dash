package project

import (
	"context"

	"github.com/blastertwist/flag-dash/internal/dto"
)

type Service interface {
	NewProject(ctx context.Context, newProjectReq *dto.NewProjectRequest) (*dto.NewProjectReponse, error)
	EditProject(ctx context.Context, editProjectReq *dto.EditProjectRequest) (*dto.EditProjectResponse, error)
	DeleteProject(ctx context.Context, deleteProjectReq *dto.DeleteProjectRequest) error
	GetProject(ctx context.Context, getProjectReq *dto.GetProjectRequest) (*dto.GetProjectResponse, error)
	GetProjects(ctx context.Context, getProjectsReq *dto.GetProjectsRequest) (*dto.GetProjectsResponse, error)
}