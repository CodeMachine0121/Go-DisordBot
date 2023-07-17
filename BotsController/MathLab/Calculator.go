package MathLab

import (
	"math"
	"time"
)

func CalculateAverageLatency(latencyList []time.Duration) time.Duration {
	var allLatency time.Duration
	for _, latency := range latencyList {
		allLatency = allLatency + latency
	}
	return allLatency / time.Duration(len(latencyList))
}

func CalculateStandardDeviationLatency(latencyList []time.Duration, averageLatency time.Duration) time.Duration {

	var varianceLatency float64
	for _, latency := range latencyList {
		varianceLatency += math.Pow(float64(latency-averageLatency), 2)
	}

	varianceLatency /= float64(len(latencyList))
	return time.Duration(math.Sqrt(varianceLatency))
}
