package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
)

type promptRegistryService struct{}

var PromptRegistryService = promptRegistryService{}

func (s *promptRegistryService) AddNewPrompt(ctx *gin.Context, opt *types.AddPromptRequest) (*string, error) {
	prompt, err := dto.PromptsRepository.AddNewPrompt(ctx, opt)
	if err != nil {
		return nil, err
	}

	promptVersionID, err := dto.PromptVersionsRepository.AddNewPromptVersion(
		ctx, &types.AddPromptVersionRequest{
			PromptID:      prompt.PromptID,
			PromptContent: "",
		})
	if err != nil {
		return nil, err
	}

	return promptVersionID, err
}

//func (s *promptRegistryService) UpdatePrompt(ctx *gin.Context, opt *types.AddPromptRequest) (*string, error) {
//	prompt, err := dto.PromptsRepository.AddNewPrompt(ctx, opt)
//	if err != nil {
//		return nil, err
//	}
//
//	promptVersionID, err := dto.PromptVersionsRepository.AddNewPromptVersion(
//		ctx, &types.AddPromptVersionRequest{
//			PromptID:      prompt.PromptID,
//			PromptContent: opt.Prompt,
//		})
//	if err != nil {
//		return nil, err
//	}
//
//	return promptVersionID, err
//}

func (s *promptRegistryService) ListPromptVersions(ctx *gin.Context,
	opt *dto.ListPromptVersionsOption) ([]*models.ViewPromptRegistryVersion, error) {

	promptId := ctx.Param("prompt_id")
	projectId := ctx.Param("project_id")

	promptVersions, err := dto.ViewPromptVersionsRepository.ListPromptVersions(ctx, projectId, promptId, opt)
	if err != nil {
		return nil, err
	}

	return promptVersions, err
}

func (s *promptRegistryService) LatestPromptForPromptId(ctx *gin.Context) (*models.ViewPromptRegistryVersion, error) {

	promptId := ctx.Param("prompt_id")
	projectId := ctx.Param("project_id")

	latestPrompt, err := dto.ViewPromptVersionsRepository.LatestPromptVersion(ctx, projectId, promptId)
	if err != nil {
		return nil, err
	}

	return latestPrompt, err
}

func (s *promptRegistryService) LatestPromptVersionsInProject(ctx *gin.Context,
	opt *dto.ListLatestPromptsOption) ([]*models.ViewPromptRegistryVersion, error) {

	projectId := ctx.Param("project_id")

	latestPrompts, err := dto.ViewLatestPromptVersionRepository.ListLatestPromptVersions(ctx, projectId, opt)
	if err != nil {
		return nil, err
	}

	return latestPrompts, err
}
