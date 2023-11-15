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
	startTime := time.Now()
	testSlice := make([]int, testArrayLength)
	for i := 0; i < testArrayLength; i++ {
		testSlice[i] = testCreateOrders[j][i]
	}
	return time.Since(startTime), testSlice
}

func testUnsortedSliceRemove(slice []int, j int) time.Duration {
	startTime := time.Now()

	for _, value := range testRemoveOrders[j] {
		for i, v := range slice {
			if v == value {
				slice = append(slice[:i], slice[i+1:]...)
				break
			}
		}
	}

	return time.Since(startTime)
}
