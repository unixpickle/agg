package main

import (
	"math"

	"github.com/unixpickle/essentials"
)

func MeanAndVariance(data <-chan float64) (mean, variance float64) {
	var sum float64
	var sqSum float64
	var count float64
	for x := range data {
		sum += x
		sqSum += x * x
		count++
	}
	mean = sum / count
	variance = besselsCorrection(sqSum/count-mean*mean, count)
	return
}

func GeometricMean(data <-chan float64) float64 {
	var logSum float64
	var count float64
	for x := range data {
		if x < 0 {
			essentials.Die("geometric mean does not support value:", x)
		}
		logSum += math.Log(x)
		count++
	}
	return math.Exp(logSum / count)
}

func besselsCorrection(variance, count float64) float64 {
	if count < 2 {
		return variance
	}
	return variance * count / (count - 1)
}
