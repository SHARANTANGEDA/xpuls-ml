package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/models"
	"github.com/xpuls-com/xpuls-ml/types"
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

func (s *langChainRunsServiceV2) GetRunFilterKeyValues(ctx *gin.Context, opt *dto.ListLangChainFilterValuesOption) ([]*string, error) {
	return dto.LangChainRunRepository.GetRunFilterKeyValues(ctx, opt, ctx.Param("project_id"))
}

func (s *langChainRunsServiceV2) AddLangChainRunToQueue(ctx *gin.Context) error {

	jsonData, err := ctx.GetRawData()
	if err != nil {
		return err
	}
	jsonDataStr := string(jsonData)
	runStep, _, err := s.extractFieldsOnStart(jsonDataStr)
	if err != nil {
		return err
	}
	queueItemId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	paramsJsonData, err := json.Marshal(runStep)
	if err != nil {
		return err
	}

	queueItem := models.LLMTracingQueue{
		QueueItemID:         queueItemId.String(),
		QueueItemType:       types.LANGCHAIN_RUN,
		Params:              paramsJsonData,
		Data:                jsonData,
		QueueProcessingCode: types.LANGCHAIN_RUN_BEGIN,
		QueueItemProcessed:  false,
	}

	_, err = dto.LLMTracingQueueRepository.AddItemToQueue(
		ctx,
		&queueItem,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *langChainRunsServiceV2) AddLangChainRunPatchToQueue(ctx *gin.Context) error {
	runId := ctx.Param("id")

	queueItemId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	jsonData, err := ctx.GetRawData()
	if err != nil {
		return err
	}

	paramsJsonData, err := json.Marshal(map[string]string{
		"run_step_id": runId,
	})
	if err != nil {
		return err
	}

	queueItem := models.LLMTracingQueue{
		QueueItemID:         queueItemId.String(),
		QueueItemType:       types.LANGCHAIN_RUN,
		Params:              paramsJsonData,
		Data:                jsonData,
		QueueProcessingCode: types.LANGCHAIN_RUN_END,
		QueueItemProcessed:  false,
	}

	_, err = dto.LLMTracingQueueRepository.AddItemToQueue(
		ctx,
		&queueItem,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *langChainRunsServiceV2) TrackRunFromQueue(ctx context.Context, params json.RawMessage, data json.RawMessage) error {

	var langChainRunStep *models.LangChainRunSteps
	if err := json.Unmarshal(params, &langChainRunStep); err != nil {
		fmt.Println("Error unmarshalling:", err)
		return nil
	}

	jsonDataStr := string(data)
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
		langChainRun, err = dto.LangChainRunRepository.AddOrGetRun(ctx, run)
		if err != nil {
			return err
		}
	} else if langChainRun.ModelInfo == nil && run.ModelInfo != nil {
		langChainRun.ModelInfo = run.ModelInfo
		_, err = dto.LangChainRunRepository.UpdateRun(ctx, langChainRun)
		if err != nil {
			return err
		}
	}

	_, err = dto.LangChainRunStepsRepository.AddNewRunStep(ctx, runStep)
	return err
}

func (s *langChainRunsServiceV2) PatchRunFromQueue(ctx context.Context, params json.RawMessage,
	data json.RawMessage) error {
	var paramsMap map[string]string
	if err := json.Unmarshal(params, &paramsMap); err != nil {
		fmt.Println("Error unmarshalling:", err)
		return nil
	}
	runId := paramsMap["run_step_id"]

	langChainStep, err := dto.LangChainRunStepsRepository.GetById(ctx, runId)
	if err != nil {
		return err
	}

	jsonDataStr := string(data)

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

	modelInfo := json.RawMessage(gjson.Get(jsonData, "serialized.kwargs.llm.kwargs").Raw)

	if modelInfo == nil {
		modelInfo = json.RawMessage(gjson.Get(jsonData, "serialized.kwargs").Raw)
	}

	var run = models.LangChainRuns{
		ChainID:            gjson.Get(jsonData, "extra.metadata.xpuls.run_id").String(),
		ProjectID:          gjson.Get(jsonData, "extra.metadata.xpuls.project_id").String(),
		ModelInfo:          modelInfo,
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

	var promptTokens int
	var completionTokens int
	var totalTokens int
	var usedTokens = 0
	if langChainRun.PromptTokens != nil {
		usedTokens = *langChainRun.PromptTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.prompt_tokens").Int())
	if usedTokens > 0 {
		promptTokens = usedTokens
	}

	if langChainRun.TotalTokens != nil {
		usedTokens = *langChainRun.TotalTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.total_tokens").Int())
	if usedTokens > 0 {
		totalTokens = usedTokens
	}

	if langChainRun.CompletionTokens != nil {
		usedTokens = *langChainRun.CompletionTokens
	} else {
		usedTokens = 0
	}
	usedTokens += int(gjson.Get(jsonData, "outputs.llm_output.token_usage.completion_tokens").Int())
	if usedTokens > 0 {
		completionTokens = usedTokens
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
		PromptTokens:       &promptTokens,
		TotalTokens:        &totalTokens,
		CompletionTokens:   &completionTokens,
	}
	return &runStep, &run, nil
}
