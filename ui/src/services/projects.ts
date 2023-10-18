import {mlServerApi} from "@/services/apiConfig";

export const fetchProjects = async (page: number, limit: number): Promise<Project[]> => {
    const response = await mlServerApi.get(`/v1/project/?page=${page}&limit=${limit}`);
    return response.data;
};

export const isProjectSlugAvailable = async (project_slug: string): Promise<boolean> => {
    const response = await mlServerApi.get(
        `/v1/project/is-slug-available?project_slug=${project_slug}`);
    return response.data;
};

export const createProject = async () => {
    const response = await mlServerApi.get(`/v1/project/create`);
    return response.data;
};