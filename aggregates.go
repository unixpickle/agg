package main

import "math"

type Aggregate func(ch <-chan float64) float64

var Aggregates = map[string]Aggregate{
	"sum":      aggregateSum,
	"mean":     aggregateMean,
	"variance": aggregateVariance,
	"stddev":   aggregateStddev,
	"geommean": GeometricMean,
}

var AggregateUsage = map[string]string{
	"sum":      "basic sum",
	"mean":     "arithmetic mean",
	"variance": "variance (with Bessel's correction)",
	"stddev":   "standard deviation (with Bessel's correction)",
	"geommean": "geometric mean",
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
