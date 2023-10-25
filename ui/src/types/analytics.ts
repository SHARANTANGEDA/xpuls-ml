interface TokenUsageQueryBuilder {
    chart_type: string;
    group_by_column: string;
    aggregation_type: string;
    time_granularity: string;
    time_from?: string | null;
    time_to?: string | null;
    time_days_ago: number;
}

interface UsageDataResponse {
    timestamp: number[];
    [key: string]: number[];
}
