package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/models"
)

type langChainRunsStepsService struct{}

var LangChainRunsStepsService = langChainRunsStepsService{}

func (s *langChainRunsStepsService) GetRunsByChainId(ctx *gin.Context) ([]*models.LangChainRunSteps, error) {
	return dto.LangChainRunStepsRepository.GetRunsByChainId(ctx, &dto.ListLangChainRunStepsOption{
		ChainId:   ctx.Param("chain_id"),
		ProjectId: ctx.Param("project_id"),
	})
}
