package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/models"
	"time"
)

type langChainRunsServiceV2 struct{}

var LangChainRunsServiceV2 = langChainRunsServiceV2{}

func (s *langChainRunsServiceV2) GetRunsInProject(ctx *gin.Context, opt *dto.ListLangChainRunOption) ([]*models.LangChainRuns, error) {
	return dto.LangChainRunRepository.GetRunsInProject(ctx, opt, ctx.Param("project_id"))
}

func (s *langChainRunsServiceV2) GetRunFilterKeys(ctx *gin.Context, opt *dto.ListLangChainRunOption) (*dto.LangChainFilterKeys, error) {
	return dto.LangChainRunRepository.GetRunFilterKeys(ctx, opt, ctx.Param("project_id"))
}

func (s *langChainRunsServiceV2) TrackRun(ctx *gin.Context) error {
	//_, ctxLocal, df, err := dto.StartTransaction(ctx)
	//if err != nil {
	//	return err
	//}
	//defer func() { df(err) }()

	jsonData, err := ctx.GetRawData()
	if err != nil {
		return err
	}
	jsonDataStr := string(jsonData)
	runStep, run, err := s.extractFieldsOnStart(jsonDataStr)
	if err != nil {
		return err
	}

	langChainRun, err := dto.LangChainRunRepository.GetById(ctx, run.ChainID)
	if err != nil {
		return err
	}
	// Check if run exists
	if langChainRun == nil {
		langChainRun, err = dto.LangChainRunRepository.AddNewRun(ctx, run)
		if err != nil {
			return err
		}
	}

	_, err = dto.LangChainRunStepsRepository.AddNewRunStep(ctx, runStep)
	return err
}

func (s *langChainRunsServiceV2) PatchRun(ctx *gin.Context) error {
	//_, ctxLocal, df, err := dto.StartTransaction(ctx)
	//if err != nil {
	//	return err
	//}
	runId := ctx.Param("id")
	//defer func() { df(err) }()

	jsonData, err := ctx.GetRawData()
	if err != nil {
		return err
	}

	langChainStep, err := dto.LangChainRunStepsRepository.GetById(ctx, runId)
	jsonDataStr := string(jsonData)
	langChainRun, err := dto.LangChainRunRepository.GetById(ctx, langChainStep.ChainID)
	if err != nil {
		return err
	}
	if langChainRun == nil {
		return fmt.Errorf("unable to find langchain run")
	}
	runStep, run, err := s.extractFieldsOnEnd(jsonDataStr, langChainRun)
	if err != nil {
		return err
	}

	langChainRun.CompletionTokens = run.CompletionTokens
	langChainRun.PromptTokens = run.PromptTokens
	langChainRun.TotalTokens = run.TotalTokens
	langChainRun.LastStepEndTime = run.LastStepEndTime
	_, err = dto.LangChainRunRepository.UpdateRun(ctx, langChainRun)
	if err != nil {
		return err
	}

	langChainStep.PromptContent = runStep.PromptContent
	langChainStep.PromptOutput = runStep.PromptOutput
	langChainStep.EventEndTime = runStep.EventEndTime
	langChainStep.TokenUsage = runStep.TokenUsage
	langChainStep.EndJSON = runStep.EndJSON

	_, err = dto.LangChainRunStepsRepository.PatchRunStep(ctx, langChainStep)
	return err
}

func (s *langChainRunsServiceV2) ParseRun(ctx *gin.Context) (*models.LangChainRunSteps, *models.LangChainRuns, error) {
	jsonData, err := ctx.GetRawData()
	if err != nil {
		return nil, nil, err
	}
	jsonDataStr := string(jsonData)
	runStep, run, err := s.extractFieldsOnStart(jsonDataStr)

	return runStep, run, err
}

func (s *langChainRunsServiceV2) getListOfStringFromGJson(jsonData string, path string) []string {
	tagsArray := gjson.Get(jsonData, path)
	var items []string
	tagsArray.ForEach(func(_, item gjson.Result) bool {
		items = append(items, item.String())
		return true
	})
	return items
}

func getGoTime(timeString string) (*time.Time, error) {
	if timeString == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02T15:04:05.999999", timeString)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *langChainRunsServiceV2) extractFieldsOnStart(jsonData string) (*models.LangChainRunSteps, *models.LangChainRuns, error) {

	eventStartTime, err := getGoTime(gjson.Get(jsonData, "events.0.time").String())
	if err != nil {
		return nil, nil, err
	}
	eventEndTime, err := getGoTime(gjson.Get(jsonData, "events.1.time").String())
	if err != nil {
		return nil, nil, err
	}

	var runStep = models.LangChainRunSteps{
		RunStepID:      gjson.Get(jsonData, "id").String(),
		RunName:        gjson.Get(jsonData, "name").String(),
		RunType:        gjson.Get(jsonData, "run_type").String(),
		ExecutionOrder: int(gjson.Get(jsonData, "execution_order").Int()),
		PromptTemplate: gjson.Get(jsonData, "serialized.kwargs.prompt.kwargs.template").String(),
		PromptTemplateInputVariables: s.getListOfStringFromGJson(jsonData,
			"serialized.kwargs.prompt.input_variables.template"),
		PromptInput:           gjson.Get(jsonData, "inputs.input").String(),
		PromptChatHistory:     gjson.Get(jsonData, "inputs.chat_history").String(),
		PromptAgentScratchpad: gjson.Get(jsonData, "inputs.agent_scratchpad").String(),
		EventStartTime:        eventStartTime,
		EventEndTime:          eventEndTime,
		ParentStepID:          gjson.Get(jsonData, "parent_run_id").String(),
		ChainID:               gjson.Get(jsonData, "extra.metadata.xpuls.run_id").String(),
		BeginJSON:             json.RawMessage(jsonData),
	}

	var run = models.LangChainRuns{
		ChainID:            gjson.Get(jsonData, "extra.metadata.xpuls.run_id").String(),
		ProjectID:          gjson.Get(jsonData, "extra.metadata.xpuls.project_id").String(),
		ModelInfo:          json.RawMessage(gjson.Get(jsonData, "serialized.kwargs.llm.kwargs").Raw),
		Labels:             json.RawMessage(gjson.Get(jsonData, "extra.metadata.xpuls.labels").Raw),
		Runtime:            json.RawMessage(gjson.Get(jsonData, "extra.runtime").Raw),
		FirstStepStartTime: eventStartTime,
		LastStepEndTime:    eventEndTime,
	}
	return &runStep, &run, nil
}

func (s *langChainRunsServiceV2) extractFieldsOnEnd(jsonData string,
	langChainRun *models.LangChainRuns) (*models.LangChainRunSteps, *models.LangChainRuns, error) {

	eventStartTime, err := getGoTime(gjson.Get(jsonData, "events.0.time").String())
	if err != nil {
		return nil, nil, err
	}
	eventEndTime, err := getGoTime(gjson.Get(jsonData, "events.1.time").String())
	if err != nil {
		return nil, nil, err
	}

	var promptTokens *int
	var completionTokens *int
	var totalTokens *int
	var usedTokens = 0
	if langChainRun.PromptTokens != nil {
		usedTokens = *langChainRun.PromptTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.prompt_tokens").Int())
	if usedTokens > 0 {
		promptTokens = &usedTokens
	}

	if langChainRun.TotalTokens != nil {
		usedTokens = *langChainRun.TotalTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.total_tokens").Int())
	if usedTokens > 0 {
		totalTokens = &usedTokens
	}

	if langChainRun.CompletionTokens != nil {
		usedTokens = *langChainRun.CompletionTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.completion_tokens").Int())
	if usedTokens > 0 {
		completionTokens = &usedTokens
	}

	var runStep = models.LangChainRunSteps{
		RunStepID:      gjson.Get(jsonData, "id").String(),
		RunName:        gjson.Get(jsonData, "name").String(),
		RunType:        gjson.Get(jsonData, "run_type").String(),
		ExecutionOrder: int(gjson.Get(jsonData, "execution_order").Int()),
		PromptContent:  gjson.Get(jsonData, "inputs.messages.0.0.kwargs.content").String(),
		PromptTemplate: gjson.Get(jsonData, "outputs.generations.0.0.message.kwargs.content").String(),
		PromptOutput:   gjson.Get(jsonData, "outputs.generations.0.0.text").String(),
		EventStartTime: eventStartTime,
		EventEndTime:   eventEndTime,
		ParentStepID:   gjson.Get(jsonData, "parent_run_id").String(),
		TokenUsage:     json.RawMessage(gjson.Get(jsonData, "outputs.llm_output.token_usage").Raw),
		ChainID:        gjson.Get(jsonData, "extra.metadata.xpuls.run_id").String(),
		EndJSON:        json.RawMessage(jsonData),
	}

	var run = models.LangChainRuns{
		ChainID:            langChainRun.ChainID,
		ProjectID:          langChainRun.ProjectID,
		ModelInfo:          langChainRun.ModelInfo,
		Labels:             langChainRun.Labels,
		Runtime:            langChainRun.Runtime,
		FirstStepStartTime: eventStartTime,
		LastStepEndTime:    eventEndTime,
		PromptTokens:       promptTokens,
		TotalTokens:        totalTokens,
		CompletionTokens:   completionTokens,
	}
	return &runStep, &run, nil
}
