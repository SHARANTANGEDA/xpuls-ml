interface PromptVersion {
    prompt_version_id: string;
    prompt_content: string;
    prompt_version_created_at: string;
    prompt_id: string;
    prompt_tag: string;
    prompt_name: string;
    project_id: string;
    prompt_created_at: string;
    prompt_deleted: boolean;
    prompt_deleted_at: string;
}

interface AddNewPrompt {
    project_id: string;
    prompt_name: string;
}