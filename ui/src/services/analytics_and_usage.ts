import {mlServerApi} from "@/services/apiConfig";

export const getTokenUsage = async (builder: TokenUsageQueryBuilder): Promise<UsageDataResponse> => {
    const response = await mlServerApi.post(
        `/v1/analytics/usage`, builder);
    return response.data;
};
