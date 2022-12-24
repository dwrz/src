package stats

import (
	"sort"
	"time"
)

type DateValue struct {
	Date  time.Time
	Value float64
}

type Stats struct {
	Values  []DateValue
	Average *float64
	Max     *DateValue
	Median  *float64
	Min     *DateValue
}

func New(values []DateValue) Stats {
	sort.Slice(values, func(i, j int) bool {
		return values[i].Value < values[j].Value
	})

	var (
		max *DateValue
		min *DateValue
		sum float64
	)
	for i := range values {
		v := values[i]

		if max == nil || v.Value >= max.Value {
			max = &v
		}
		if min == nil || v.Value <= min.Value {
			min = &v
		}

		sum += v.Value
	}

	var stats = Stats{
		Max:    max,
		Min:    min,
		Values: values,
	}
	if len(values) > 0 {
		avg := sum / float64(len(values))
		stats.Average = &avg

		middle := len(values) / 2

		if len(values)%2 != 0 {
			median := values[middle]

			stats.Median = &median.Value
		} else {
			median := (values[middle-1].Value +
				values[middle].Value) / 2

			stats.Median = &median
		}
	}

	return stats
}
