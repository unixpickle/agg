package main

import (
	"fmt"
	"math"
)

type Aggregate func(ch <-chan float64) any

func wrapFloatFn(f func(ch <-chan float64) float64) Aggregate {
	return func(ch <-chan float64) any {
		return f(ch)
	}
}

var Aggregates = map[string]Aggregate{
	"sum":      wrapFloatFn(aggregateSum),
	"mean":     wrapFloatFn(aggregateMean),
	"variance": wrapFloatFn(aggregateVariance),
	"stddev":   wrapFloatFn(aggregateStddev),
	"geommean": wrapFloatFn(GeometricMean),
	"max":      wrapFloatFn(aggregateMax),
	"min":      wrapFloatFn(aggregateMin),
	"meanerr":  aggregateMeanWithErr,
}

var AggregateUsage = map[string]string{
	"sum":      "basic sum",
	"mean":     "arithmetic mean",
	"variance": "variance (with Bessel's correction)",
	"stddev":   "standard deviation (with Bessel's correction)",
	"geommean": "geometric mean",
	"max":      "maximum value",
	"min":      "minimum value",
	"meanerr":  "mean with error margin 1.96*stddev/sqrt(N)",
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

func aggregateMeanWithErr(ch <-chan float64) any {
	meanCh := make(chan float64)
	stdCh := make(chan float64)
	countOut := make(chan int, 1)
	go func() {
		defer close(meanCh)
		defer close(stdCh)
		n := 0
		for x := range ch {
			meanCh <- x
			stdCh <- x
			n += 1
		}
		countOut <- n
	}()
	meanOut := make(chan float64, 1)
	stdOut := make(chan float64, 1)
	go func() {
		meanOut <- aggregateMean(meanCh)
	}()
	go func() {
		stdOut <- aggregateStddev(stdCh)
	}()
	mean := <-meanOut
	std := <-stdOut
	count := <-countOut
	return fmt.Sprintf("%v ± %v", mean, 1.96*std/float64(count))
}
