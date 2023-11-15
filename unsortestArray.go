package main

import "time"

func testUnsortedArray() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, array := testUnsortedArrayCreate(i)
		createTimes[i] = createTime
		removeTime := testUnSortedArrayRemove(array, i)
		removeTimes[i] = removeTime
	}

	unsortedCreateStats := getStatistics(createTimes[:])
	unsortedRemoveStats := getStatistics(removeTimes[:])
	unsortedCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Unsorted Array", unsortedCreateStats, unsortedRemoveStats, unsortedCombinedStats)
}

func testUnsortedArrayCreate(j int) (time.Duration, *[testArrayLength]int) {
	var addTime time.Duration
	var testArray [testArrayLength]int
	for i := 0; i < testArrayLength; i++ {
		startTime := time.Now()
		testArray[i] = testCreateOrders[j][i]
		addTime += time.Since(startTime)
	}
	return addTime, &testArray
}

func testUnSortedArrayRemove(array *[testArrayLength]int, j int) time.Duration {
	var removeTime time.Duration
	for _, value := range testRemoveOrders[j] {
		removeIndex := -1
		for i, v := range array {
			if v == value {
				removeIndex = i
				break
			}
		}

		if removeIndex != -1 {
			startTime := time.Now()
			for i := removeIndex; i < testArrayLength-1; i++ {
				array[i] = array[i+1]
			}
			array[testArrayLength-1] = 0
			removeTime += time.Since(startTime)
		}
	}

	return removeTime
}
