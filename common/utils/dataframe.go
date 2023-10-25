package utils

import (
	"github.com/xpuls-com/xpuls-ml/types"
	"sort"
	"time"
)

func indexOf(times []time.Time, t time.Time) int {
	for i, time := range times {
		if time.Equal(t) {
			return i
		}
	}
	return -1
}

func PivotDataToDataFrame(metrics []*types.TimeSeriesData, defaultValue interface{}) map[string][]int {
	// Sort metrics by Time
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Time.Before(*metrics[j].Time)
	})

	// Map to hold data
	dataMap := make(map[time.Time]map[string]int)

	// Populate the map with metrics
	for _, m := range metrics {
		if _, exists := dataMap[*m.Time]; !exists {
			dataMap[*m.Time] = make(map[string]int)
		}
		dataMap[*m.Time][m.Dimension] = m.Metric
	}

	// Extract sorted unique timestamps
	var uniqueTimes []time.Time
	for t := range dataMap {
		uniqueTimes = append(uniqueTimes, t)
	}
	sort.Slice(uniqueTimes, func(i, j int) bool {
		return uniqueTimes[i].Before(uniqueTimes[j])
	})

	// Extract unique dimensions
	dimensionSet := make(map[string]bool)
	for _, m := range metrics {
		dimensionSet[m.Dimension] = true
	}

	var uniqueDimensions []string
	for d := range dimensionSet {
		uniqueDimensions = append(uniqueDimensions, d)
	}

	// Create result structure
	result := make(map[string][]int)
	result["timestamp"] = make([]int, len(uniqueTimes))
	for i, t := range uniqueTimes {
		result["timestamp"][i] = int(t.Unix())
	}

	// Initialize dimensions
	for _, dim := range uniqueDimensions {
		result[dim] = make([]int, len(uniqueTimes))
		if intVal, ok := defaultValue.(int); ok {
			for i := range result[dim] {
				result[dim][i] = intVal
			}
		}
	}

	for i, t := range uniqueTimes {
		for _, dim := range uniqueDimensions {
			if val, exists := dataMap[t][dim]; exists {
				result[dim][i] = val
			} else if i > 0 && defaultValue == "cumulative" {
				result[dim][i] = result[dim][i-1]
			}
		}
	}

	// Adjust for empty or null dimensions
	if _, exists := result[""]; exists {
		result["others"] = result[""]
		delete(result, "")
	}

	return result
}
