package main

import "time"

func testMap() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, testMap := testMapCreate()
		createTimes[i] = createTime
		removeTime := testMapRemove(testMap)
		removeTimes[i] = removeTime
	}

	mapCreateStats := getStatistics(createTimes[:])
	mapRemoveStats := getStatistics(removeTimes[:])
	mapCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Map", mapCreateStats, mapRemoveStats, mapCombinedStats)
}

func testMapCreate() (time.Duration, map[int]int) {
	startTime := time.Now()
	testMap := make(map[int]int)
	for i, value := range testCreateOrder {
		testMap[i] = value
	}
	return time.Since(startTime), testMap
}

func testMapRemove(testMap map[int]int) time.Duration {
	startTime := time.Now()
	for _, value := range testRemoveOrder {
		delete(testMap, value)
	}
	return time.Since(startTime)
}
