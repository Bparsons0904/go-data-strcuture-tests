package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
)

const testArrayLength int = 50000
const numberOfTests int = 100

var testCreateOrders [][]int
var testRemoveOrders [][]int

func main() {
	for i := 0; i < numberOfTests; i++ {
		testCreateOrder, _ := faker.RandomInt(testArrayLength)
		testRemoveOrder, _ := faker.RandomInt(testArrayLength)
		testCreateOrders = append(testCreateOrders, testCreateOrder)
		testRemoveOrders = append(testRemoveOrders, testRemoveOrder)
	}

	fmt.Printf("\n%-30s %-10s %-10s %-10s %-10s %-10s\n", "Type", "Mean", "Min", "Max", "Median", "Std Dev")
	fmt.Println(strings.Repeat("-", 82))

	wg := &sync.WaitGroup{}
	wg.Add(6)

	go func() {
		testUnsortedArray()
		wg.Done()
	}()

	go func() {
		testUnsortedSlice()
		wg.Done()
	}()

	go func() {
		testLinkedList()
		wg.Done()
	}()

	go func() {
		testDoubleLinkedList()
		wg.Done()
	}()

	go func() {
		testMap()
		wg.Done()
	}()

	go func() {
		testBinaryTree()
		wg.Done()
	}()

	wg.Wait()
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

func formatDuration(d time.Duration) string {
	microseconds := d.Microseconds()
	if microseconds >= 10000 {
		return fmt.Sprintf("%.2fms", d.Seconds()*1000)
	}
	return fmt.Sprintf("%dµs", microseconds)
}

func printStatistics(label string, createStats Statistics, removeStats Statistics, combinedStats Statistics) {
	fmt.Printf("%-30s %-10s %-10s %-10s %-10s %-10.2f\n",
		label+" Create",
		formatDuration(createStats.mean),
		formatDuration(createStats.min),
		formatDuration(createStats.max),
		formatDuration(createStats.median),
		createStats.standardDeviation)

	fmt.Printf("%-30s %-10s %-10s %-10s %-10s %-10.2f\n",
		label+" Remove",
		formatDuration(removeStats.mean),
		formatDuration(removeStats.min),
		formatDuration(removeStats.max),
		formatDuration(removeStats.median),
		removeStats.standardDeviation)

	fmt.Printf("\033[1m%-30s %-10s %-10s %-10s %-10s %-10.2f\033[0m\n",
		label+" Combined",
		formatDuration(combinedStats.mean),
		formatDuration(combinedStats.min),
		formatDuration(combinedStats.max),
		formatDuration(combinedStats.median),
		combinedStats.standardDeviation)

	fmt.Println(strings.Repeat("-", 82))
}
