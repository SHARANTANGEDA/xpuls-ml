package dto

import (
	"context"
	"fmt"
	"github.com/xpuls-com/xpuls-ml/analytics"
	"github.com/xpuls-com/xpuls-ml/types"
	"gorm.io/gorm"
)

type AnalyticsQueryOption struct {
	BaseListOption
}

type analyticsQueriesRepository struct{}

var AnalyticsQueriesRepository = analyticsQueriesRepository{}

func (s *analyticsQueriesRepository) getBaseDB(ctx context.Context) *gorm.DB {
	return mustGetSession(ctx)
}

func (s *analyticsQueriesRepository) RunSqlQueryTimeSeries(ctx context.Context,
	builder *types.AnalyticsQueryBuilder) ([]*types.TimeSeriesData, error) {
	//var total int64
	var rawSqlQuery = ""
	if builder.AggregationType == types.AGGREGATION_DEFAULT {
		rawSqlQuery = analytics.GetPSQLTimeSeriesQuery(builder.GroupByColumn, builder.TimeGranularity,
			builder.TimeFrom, builder.TimeTo, builder.TimeDaysAgo)
	} else if builder.AggregationType == types.AGGREGATION_CUMULATIVE {
		rawSqlQuery = analytics.GetPSQLTimeSeriesCumulativeQuery(builder.GroupByColumn, builder.TimeGranularity,
			builder.TimeFrom, builder.TimeTo, builder.TimeDaysAgo)
	} else {
		return nil, fmt.Errorf("invalid aggregation_type")
	}

	query := getBaseQuery(ctx, s).Model(&types.TimeSeriesData{}).Raw(rawSqlQuery)
	//err := query.Count(&total).Error
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		// No record found, return nil without error
	//		return nil, nil
	//	}
	//	return nil, err
	//}
	timeSeriesData := make([]*types.TimeSeriesData, 0)

	err := query.Find(&timeSeriesData).Error
	if err != nil {
		return nil, err
	}

	return timeSeriesData, nil
}
