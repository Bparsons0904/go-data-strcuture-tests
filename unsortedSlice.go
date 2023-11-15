package main

import "time"

func testUnsortedSlice() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, slice := testUnsortedSliceCreate(i)
		createTimes[i] = createTime
		removeTime := testUnsortedSliceRemove(slice, i)
		removeTimes[i] = removeTime
	}

	unsortedCreateStats := getStatistics(createTimes[:])
	unsortedRemoveStats := getStatistics(removeTimes[:])
	unsortedCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Unsorted Slice", unsortedCreateStats, unsortedRemoveStats, unsortedCombinedStats)
}

func testUnsortedSliceCreate(j int) (time.Duration, []int) {
	var totalTime time.Duration
	testSlice := make([]int, testArrayLength)
	for i := 0; i < testArrayLength; i++ {
		startTime := time.Now()
		testSlice[i] = testCreateOrders[j][i]
		totalTime += time.Since(startTime)
	}
	return totalTime, testSlice
}

func testUnsortedSliceRemove(slice []int, j int) time.Duration {
	var totalTime time.Duration
	for _, value := range testRemoveOrders[j] {
		for i, v := range slice {
			startTime := time.Now()
			if v == value {
				slice = append(slice[:i], slice[i+1:]...)
				totalTime += time.Since(startTime)
				break
			}
		}
	}

	return totalTime
}
