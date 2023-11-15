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

const testArrayLength int = 10000
const numberOfTests int = 100

var testCreateOrders [][]int
var testRemoveOrders [][]int

func main() {
	initTestData()
	printTestHeader()
	runTestsConcurrently()
}

func initTestData() {
	testCreateOrders = make([][]int, numberOfTests)
	testRemoveOrders = make([][]int, numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		testCreateOrders[i], _ = faker.RandomInt(testArrayLength)
		testRemoveOrders[i], _ = faker.RandomInt(testArrayLength)
	}
}

func printTestHeader() {
	fmt.Println("\nTest Array Length:", testArrayLength)
	fmt.Println("Number of Tests:", numberOfTests)
	fmt.Println("Test Run Time:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("\n%-30s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-15s\n",
		"Type", "Mean", "Median", "Min", "Max", "95th Perc", "Ops/s", "Std Dev", "Total")
	fmt.Println(strings.Repeat("-", 115))
}

func runTestsConcurrently() {
	wg := &sync.WaitGroup{}
	tests := []func(){
		testUnsortedArray,
		testUnsortedSlice,
		testLinkedList,
		testDoubleLinkedList,
		testMap,
		testBinaryTree,
	}

	wg.Add(len(tests))
	for _, test := range tests {
		go func(testFunc func()) {
			testFunc()
			wg.Done()
		}(test)
	}
	wg.Wait()
}

type Statistics struct {
	min               time.Duration
	max               time.Duration
	mean              time.Duration
	median            time.Duration
	percentile95      time.Duration
	standardDeviation float64
	throughput        float64
	durations         []time.Duration
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
	stats.percentile95 = calculatePercentile(times, 95)
	stats.throughput = calculateThroughput(times)
	stats.standardDeviation = calculateStandardDeviation(times, stats.mean)
	stats.durations = times
	return stats
}

func calculatePercentile(durations []time.Duration, percentile float64) time.Duration {
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })
	index := int(percentile / 100.0 * float64(len(durations)))
	if index >= len(durations) {
		index = len(durations) - 1
	}
	return durations[index]
}

func calculateThroughput(durations []time.Duration) float64 {
	var totalDuration time.Duration
	for _, d := range durations {
		totalDuration += d
	}
	return float64(len(durations)) / totalDuration.Seconds()
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
	if microseconds >= 1000 {
		return fmt.Sprintf("%.2fms", d.Seconds()*1000)
	}
	return fmt.Sprintf("%dµs", microseconds)
}

func formatTotalDuration(durations []time.Duration) string {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	if total.Seconds() >= 1 {
		return fmt.Sprintf("%.2fs", total.Seconds())
	}
	return fmt.Sprintf("%.2fms", total.Seconds()*1000)
}

func formatThroughput(throughput float64) string {
	if throughput >= 1e6 {
		return fmt.Sprintf("%.0fm", throughput/1e6)
	} else if throughput >= 1e3 {
		return fmt.Sprintf("%.0fk", throughput/1e3)
	}
	return fmt.Sprintf("%.0f", throughput)
}

func printStatistic(label string, stats Statistics) {
	var throughput float64
	meanOperationTime := stats.mean.Seconds()
	if meanOperationTime > 0 {
		throughput = 1 / meanOperationTime
	} else {
		throughput = 0
	}

	sort.Slice(stats.durations, func(i, j int) bool { return stats.durations[i] < stats.durations[j] })
	p95Index := int(float64(len(stats.durations)) * 0.95)
	percentile95 := stats.durations[p95Index]

	fmt.Printf("%-30s %-10s %-10s %-10s %-10s %-10s %-10s %-10s %-15s\n",
		label,
		formatDuration(stats.mean),
		formatDuration(stats.median),
		formatDuration(stats.min),
		formatDuration(stats.max),
		formatDuration(percentile95),
		formatThroughput(throughput),
		fmt.Sprintf("%.2f", stats.standardDeviation),
		formatTotalDuration(stats.durations),
	)

	if strings.Contains(label, "Combined") {
		fmt.Println(strings.Repeat("-", 115))
	}
}

func printStatistics(label string, createStats Statistics, removeStats Statistics, combinedStats Statistics) {
	printStatistic(fmt.Sprintf("%s Create", label), createStats)
	printStatistic(fmt.Sprintf("%s Remove", label), removeStats)
	printStatistic(fmt.Sprintf("%s Combined", label), combinedStats)
}
