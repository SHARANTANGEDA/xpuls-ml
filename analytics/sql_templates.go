package analytics

import (
	"github.com/xpuls-com/xpuls-ml/types"
	"strconv"
	"time"
)

// GetPSQLTimeSeriesQuery Time-series SQL Queries
func GetPSQLTimeSeriesQuery(groupByColumn string, timeGranularity types.TimeGranularity, timeFrom *time.Time,
	timeTo *time.Time, timeDaysAgo *int) string {
	timeFilter := ""
	if timeDaysAgo != nil {
		timeFilter = "WHERE chain_tracked_at >= current_date - INTERVAL '" + strconv.Itoa(*timeDaysAgo) + " days'"
	} else {
		//your_timestamp BETWEEN $1 AND $2
		if timeTo == nil {
			timeFilter = "WHERE chain_tracked_at >= '" + timeFrom.String() + "'"
		} else {
			timeFilter = "WHERE chain_tracked_at BETWEEN '" + timeFrom.String() + "' AND '" + timeTo.String() + "'"
		}
	}

	dateTrunc := "date_trunc('" + string(timeGranularity) + "', chain_tracked_at)"
	dimension := "model_info->>'" + groupByColumn + "'"
	sqlQuery := "SELECT " + dateTrunc + " AS time" +
		", " + dimension + " as dimension" + ", SUM(total_tokens) as metric FROM langchain_runs \n" +
		timeFilter + " \n " +
		//"GROUP BY " + dateTrunc + ", " + dimension + "\n " +
		"GROUP BY time, dimension " +
		"ORDER BY time, dimension"
	return sqlQuery
}

// GetPSQLTimeSeriesCumulativeQuery Time-series SQL Queries
func GetPSQLTimeSeriesCumulativeQuery(groupByColumn string, timeGranularity types.TimeGranularity, timeFrom *time.Time,
	timeTo *time.Time, timeDaysAgo *int) string {

	sqlQuery := "WITH AggregatedData AS ( \n " + GetPSQLTimeSeriesQuery(groupByColumn, timeGranularity, timeFrom,
		timeTo, timeDaysAgo) + " ) \n " +
		`
		SELECT
		    time,
		    dimension,
		    SUM(metric) OVER (PARTITION BY dimension ORDER BY time) AS metric
		FROM AggregatedData
		ORDER BY time, dimension
		`

	/*
		WITH AggregatedData AS (
		    SELECT
		        date_trunc('day', chain_tracked_at) AS time,
		        model_info->>'model_name' AS dimension,
		        SUM(total_tokens) AS daily_total_tokens
		    FROM langchain_runs
		    WHERE chain_tracked_at >= current_date - INTERVAL '45 days'
		    GROUP BY time, dimension
		)

		SELECT
		    time,
		    dimension,
		    SUM(daily_total_tokens) OVER (PARTITION BY dimension ORDER BY time) AS cumulative_metric
		FROM AggregatedData
		ORDER BY time, dimension;

	*/
	return sqlQuery
}
