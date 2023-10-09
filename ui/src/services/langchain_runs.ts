import {mlServerApi} from "@/services/apiConfig";

export const fetchLangChainRuns = async (project_id: string, page: number, limit: number, sort_field: string,
                                         sort_order: string) => {
    const response = await mlServerApi.get(
        `/v1/langchain/${project_id}/runs?start=${page}&count=${limit}&sort_field=${sort_field}&sort_order=${sort_order}`);
    return response.data;
};

export const fetchLangChainRunSteps = async (project_id: string, chain_id: string) => {
    const response = await mlServerApi.get(`/v1/langchain/${project_id}/runs/${chain_id}`);
    return response.data;
};

export const fetchLangChainFilterKeys = async (project_id: string, page: number, limit: number) => {
    const response = await mlServerApi.get(
        `/v1/langchain/${project_id}/runs/filters/keys?start=${page}&count=${limit}`);
    return response.data;
};

export const fetchLangChainFilterValues = async (project_id: string, label_key: string, condition: string,
                                                 search_value: string, page: number, limit: number) => {
    const response = await mlServerApi.get(
        `/v1/langchain/${project_id}/runs/filters/values?label_key=${label_key}&condition=${condition}&search_value=${search_value}&start=${page}&count=${limit}`);
    return response.data;
};