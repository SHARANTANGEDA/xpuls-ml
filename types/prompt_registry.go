package types

type AddPromptRequest struct {
	//Prompt     string `json:"prompt"`
	ProjectId  string `json:"project_id"`
	PromptName string `json:"prompt_name"`
}

type AddPromptVersionRequest struct {
	PromptContent string `json:"prompt_content"`
	PromptID      string `json:"prompt_id"`
}
