import {mlServerApi} from "@/services/apiConfig";

export const fetchProjects = async (page: number, limit: number) => {
    const response = await mlServerApi.get(`/v1/project/?page=${page}&limit=${limit}`);
    return response.data;
};