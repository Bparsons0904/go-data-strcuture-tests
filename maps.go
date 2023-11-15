package main

import "time"

func testMap() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, testMap := testMapCreate(i)
		createTimes[i] = createTime
		removeTime := testMapRemove(testMap, i)
		removeTimes[i] = removeTime
	}

	mapCreateStats := getStatistics(createTimes[:])
	mapRemoveStats := getStatistics(removeTimes[:])
	mapCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Map", mapCreateStats, mapRemoveStats, mapCombinedStats)
}

func testMapCreate(i int) (time.Duration, map[int]int) {
	startTime := time.Now()
	testMap := make(map[int]int)
	for i, value := range testCreateOrders[i] {
		testMap[i] = value
	}
	return time.Since(startTime), testMap
}

func testMapRemove(testMap map[int]int, i int) time.Duration {
	startTime := time.Now()
	for _, value := range testRemoveOrders[i] {
		delete(testMap, value)
	}
	return time.Since(startTime)
}
