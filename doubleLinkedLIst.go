package main

import "time"

type DoubleListNode struct {
	Value int
	Next  *DoubleListNode
	Prev  *DoubleListNode
}

type DoubleLinkedList struct {
	Head *DoubleListNode
	Tail *DoubleListNode
}

func (doubleLinkedList *DoubleLinkedList) Append(value int) {
	newNode := &DoubleListNode{Value: value}
	if doubleLinkedList.Head == nil {
		doubleLinkedList.Head = newNode
		doubleLinkedList.Tail = newNode
		return
	}
	doubleLinkedList.Tail.Next = newNode
	newNode.Prev = doubleLinkedList.Tail
	doubleLinkedList.Tail = newNode
}

func (doubleLinkedList *DoubleLinkedList) Remove(value int) {
	current := doubleLinkedList.Head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				doubleLinkedList.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				doubleLinkedList.Tail = current.Prev
			}
			return
		}
		current = current.Next
	}
}

func testDoubleLinkedList() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, doubleLinkedList := testDoubleLinkedListCreate(i)
		createTimes[i] = createTime
		removeTime := testDoubleLinkedListRemove(doubleLinkedList, i)
		removeTimes[i] = removeTime
	}

	doubleLinkedListCreateStats := getStatistics(createTimes[:])
	doubleLinkedListRemoveStats := getStatistics(removeTimes[:])
	doubleLinkedListCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Double Linked List", doubleLinkedListCreateStats, doubleLinkedListRemoveStats, doubleLinkedListCombinedStats)
}

func testDoubleLinkedListCreate(i int) (time.Duration, *DoubleLinkedList) {
	startTime := time.Now()
	doubleLinkedList := &DoubleLinkedList{}
	for _, value := range testCreateOrders[i] {
		doubleLinkedList.Append(value)
	}
	return time.Since(startTime), doubleLinkedList
}

func testDoubleLinkedListRemove(doubleLinkedList *DoubleLinkedList, i int) time.Duration {
	startTime := time.Now()
	for _, value := range testRemoveOrders[i] {
		doubleLinkedList.Remove(value)
	}
	return time.Since(startTime)
}
