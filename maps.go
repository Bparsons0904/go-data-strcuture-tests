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
	var addTime time.Duration
	testMap := make(map[int]int, testArrayLength)

	for i, value := range testCreateOrders[i] {
		startTime := time.Now()
		testMap[i] = value
		addTime += time.Since(startTime)
	}
	return addTime, testMap
}

func testMapRemove(testMap map[int]int, i int) time.Duration {
	var removeTime time.Duration
	for _, value := range testRemoveOrders[i] {
		startTime := time.Now()
		delete(testMap, value)
		removeTime += time.Since(startTime)
	}
	return removeTime
}
