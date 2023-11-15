package main

import "time"

type ListNode struct {
	Value int
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (linkedList *LinkedList) Append(value int) {
	newNode := &ListNode{Value: value}
	if linkedList.Head == nil {
		linkedList.Head = newNode
		return
	}
	current := linkedList.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (linkedList *LinkedList) Remove(value int) time.Duration {
	startTime := time.Now()

	if linkedList.Head == nil {
		return time.Since(startTime)
	}
	if linkedList.Head.Value == value {
		linkedList.Head = linkedList.Head.Next
		return time.Since(startTime)
	}
	current := linkedList.Head
	for current.Next != nil {
		if current.Next.Value == value {
			startTime = time.Now()
			current.Next = current.Next.Next
			return time.Since(startTime)
		}
		current = current.Next
	}
	return time.Since(startTime)
}

func testLinkedList() {
	var createTimes [numberOfTests]time.Duration
	var removeTimes [numberOfTests]time.Duration

	for i := 0; i < numberOfTests; i++ {
		createTime, linkedList := testLinkedListCreate(i)
		createTimes[i] = createTime
		removeTime := testLinkedListRemove(linkedList, i)
		removeTimes[i] = removeTime
	}

	linkedListCreateStats := getStatistics(createTimes[:])
	linkedListRemoveStats := getStatistics(removeTimes[:])
	linkedListCombinedStats := combineStatistics(createTimes[:], removeTimes[:])
	printStatistics("Linked List", linkedListCreateStats, linkedListRemoveStats, linkedListCombinedStats)
}

func testLinkedListCreate(i int) (time.Duration, *LinkedList) {
	var totalTime time.Duration
	linkedList := &LinkedList{}
	for _, value := range testCreateOrders[i] {
		startTime := time.Now()
		linkedList.Append(value)
		totalTime += time.Since(startTime)
	}
	return totalTime, linkedList
}

func testLinkedListRemove(linkedList *LinkedList, i int) time.Duration {
	for _, value := range testRemoveOrders[i] {
		removeTime := linkedList.Remove(value)
		return removeTime
	}

	return 0
}
