import {mlServerApi} from "@/services/apiConfig";

export const getPromptVersions = async (project_id: string, prompt_id: string, page: number=0,
                                        limit: number=20): Promise<PromptVersion[]> => {
    const response = await mlServerApi.get(
        `/v1/registry/${project_id}/prompt/${prompt_id}/versions?start=${page}&count=${limit}`);
    return response.data;
};

export const getLatestPrompt = async (project_id: string, prompt_id: string): Promise<PromptVersion> => {
    const response = await mlServerApi.get(
        `/v1/registry/${project_id}/prompt/${prompt_id}/latest`);
    return response.data;
};

export const addNewPrompt = async (newPrompt: AddNewPrompt) => {
    const response = await mlServerApi.post(
        `/v1/registry/${newPrompt.project_id}/prompt`, newPrompt);
    return response.data;
};

export const getLatestPromptsInProject = async (project_id: string, page: number=0,
                                                limit: number=20): Promise<PromptVersion[]> => {
    const response = await mlServerApi.get(
        `/v1/registry/${project_id}/prompt?start=${page}&count=${limit}`);
    return response.data;
};

