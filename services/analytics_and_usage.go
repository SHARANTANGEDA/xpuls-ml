package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xpuls-com/xpuls-ml/common/utils"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/types"
)

type analyticsAndUsageService struct{}

var AnalyticsAndUsageService = analyticsAndUsageService{}

func (s *analyticsAndUsageService) GetSqlTimeSeriesData(ctx *gin.Context,
	builder *types.AnalyticsQueryBuilder) (map[string][]int, error) {
	timeSeriesData, err := dto.AnalyticsQueriesRepository.RunSqlQueryTimeSeries(ctx, builder)
	if err != nil {
		return nil, err
	}

	if builder.AggregationType == types.AGGREGATION_CUMULATIVE {
		return utils.PivotDataToDataFrame(timeSeriesData, "cumulative"), nil
	} else {
		return utils.PivotDataToDataFrame(timeSeriesData, 0), nil
	}
}
