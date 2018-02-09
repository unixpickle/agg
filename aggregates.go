package main

import "math"

type Aggregate func(ch <-chan float64) float64

var Aggregates = map[string]Aggregate{
	"sum":      aggregateSum,
	"mean":     aggregateMean,
	"variance": aggregateVariance,
	"stddev":   aggregateStddev,
	"geommean": GeometricMean,
	"max":      aggregateMax,
	"min":      aggregateMin,
}

var AggregateUsage = map[string]string{
	"sum":      "basic sum",
	"mean":     "arithmetic mean",
	"variance": "variance (with Bessel's correction)",
	"stddev":   "standard deviation (with Bessel's correction)",
	"geommean": "geometric mean",
	"max":      "maximum value",
	"min":      "minimum value",
}

func aggregateSum(ch <-chan float64) float64 {
	var sum float64
	for val := range ch {
		sum += val
	}
	return sum
}

func aggregateMean(ch <-chan float64) float64 {
	mean, _ := MeanAndVariance(ch)
	return mean
}

func aggregateVariance(ch <-chan float64) float64 {
	_, variance := MeanAndVariance(ch)
	return variance
}

func aggregateStddev(ch <-chan float64) float64 {
	return math.Sqrt(aggregateVariance(ch))
}

func aggregateMax(ch <-chan float64) float64 {
	max := math.Inf(-1)
	for num := range ch {
		max = math.Max(max, num)
	}
	return max
}

func aggregateMin(ch <-chan float64) float64 {
	min := math.Inf(1)
	for num := range ch {
		min = math.Min(min, num)
	}
	return min
}
