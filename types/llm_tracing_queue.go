package types

type LLMTracingQueueItemTypes string

const (
	LANGCHAIN_RUN LLMTracingQueueItemTypes = "LANGCHAIN_RUN"
)

type LLMTracingQueueProcessingTypes string

const (
	LANGCHAIN_RUN_BEGIN LLMTracingQueueProcessingTypes = "LANGCHAIN_RUN_BEGIN"
	LANGCHAIN_RUN_END   LLMTracingQueueProcessingTypes = "LANGCHAIN_RUN_END"
)
