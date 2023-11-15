package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/go-faker/faker/v4"
)

const testArrayLength int = 10000
const numberOfTests int = 100

var testCreateOrder []int
var testRemoveOrder []int

func main() {
	testCreateOrder, _ = faker.RandomInt(testArrayLength)
	testRemoveOrder, _ = faker.RandomInt(testArrayLength)

	fmt.Printf("\n%-30s %-10s %-10s %-10s %-10s %-10s\n", "Type", "Mean", "Min", "Max", "Median", "Std Dev")
	fmt.Println(strings.Repeat("-", 82))

	testUnsortedArray()
	testUnsortedSlice()
	testLinkedList()
	testDoubleLinkedList()
	testMap()
}

type Statistics struct {
	min               time.Duration
	max               time.Duration
	mean              time.Duration
	median            time.Duration
	standardDeviation float64
}

func getStatistics(times []time.Duration) Statistics {
	var stats Statistics
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})
	stats.min = times[0]
	stats.max = times[numberOfTests-1]
	stats.mean = calculateMean(times)
	stats.median = calculateMedian(times)
	stats.standardDeviation = calculateStandardDeviation(times, stats.mean)
	return stats
}

func calculateMedian(durations []time.Duration) time.Duration {
	middle := len(durations) / 2
	median := durations[middle]
	if len(durations)%2 == 0 {
		median = (median + durations[middle-1]) / 2
	}
	return median
}

func calculateStandardDeviation(durations []time.Duration, mean time.Duration) float64 {
	var sum float64

	for _, duration := range durations {
		difference := duration - mean
		sum += math.Pow(difference.Seconds(), 2)
	}

	meanOfDifferences := sum / float64(len(durations))
	return math.Sqrt(meanOfDifferences)
}

func calculateMean(durations []time.Duration) time.Duration {
	var sum time.Duration

	for _, duration := range durations {
		sum += duration
	}

	return sum / time.Duration(len(durations))
}

func combineStatistics(createStats, removeStats []time.Duration) Statistics {
	var combinedDurations []time.Duration
	for i := 0; i < len(createStats); i++ {
		combinedDurations = append(combinedDurations, createStats[i]+removeStats[i])
	}
	return getStatistics(combinedDurations)
}

func printStatistics(label string, createStats Statistics, removeStats Statistics, combinedStats Statistics) {
	fmt.Printf("%-30s %-10v %-10v %-10v %-10v %-10.2f\n",
		label+" Create",
		createStats.mean.Microseconds(),
		createStats.min.Microseconds(),
		createStats.max.Microseconds(),
		createStats.median.Microseconds(),
		createStats.standardDeviation)

	fmt.Printf("%-30s %-10v %-10v %-10v %-10v %-10.2f\n",
		label+" Remove",
		removeStats.mean.Microseconds(),
		removeStats.min.Microseconds(),
		removeStats.max.Microseconds(),
		removeStats.median.Microseconds(),
		removeStats.standardDeviation)

	fmt.Printf("%-30s %-10v %-10v %-10v %-10v %-10.2f\n",
		label+" Combined",
		combinedStats.mean.Microseconds(),
		combinedStats.min.Microseconds(),
		combinedStats.max.Microseconds(),
		combinedStats.median.Microseconds(),
		combinedStats.standardDeviation)

}
