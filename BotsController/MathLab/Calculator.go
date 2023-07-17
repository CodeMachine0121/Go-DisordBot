package MathLab

import (
	"time"
)

func CalculateAverageLatency(latencyList []time.Duration) time.Duration {
	var allLatency time.Duration
	for _, latency := range latencyList {
		allLatency = allLatency + latency
	}
	return allLatency / time.Duration(len(latencyList))
}
