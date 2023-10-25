package types

import "time"

type TimeGranularity string

const (
	TIME_DAY   TimeGranularity = "DAY"
	TIME_MONTH TimeGranularity = "MONTH"
	TIME_HOUR  TimeGranularity = "HOUR"
	TIME_YEAR  TimeGranularity = "YEAR"
)

type AggregationType string

const (
	AGGREGATION_DEFAULT    AggregationType = "default"
	AGGREGATION_CUMULATIVE AggregationType = "cumulative"
)

type ChartType string

const (
	CHART_LINE ChartType = "line"
	CHART_BAR  ChartType = "bar"
	CHART_PIE  ChartType = "pie"
)

type AnalyticsQueryBuilder struct {
	ChartType       ChartType       `json:"chart_type"`
	GroupByColumn   string          `json:"group_by_column"`
	AggregationType AggregationType `json:"aggregation_type" default:"DEFAULT"`
	TimeGranularity TimeGranularity `json:"time_granularity"`
	TimeFrom        *time.Time      `json:"time_from"`
	TimeTo          *time.Time      `json:"time_to"`
	TimeDaysAgo     *int            `json:"time_days_ago"`
}

type TimeSeriesData struct {
	Dimension string     `json:"dimension"`
	Time      *time.Time `json:"time"`
	Metric    int        `json:"metric"`
}

type TimeSeriesDataFrame struct {
	Time   time.Time      `json:"time"`
	Values map[string]int `json:"data"`
}
