# Data Structure Performance Tests

## Overview

This project is dedicated to assessing and comparing the performance of various data structures in Go. The primary focus is on understanding how different data structures behave under operations such as adding and removing records. These tests are crucial for identifying the most efficient data structures for specific use cases, especially in scenarios where performance is critical.

## Data Structures Tested

- **Unsorted Array**: Tests how arrays handle insertion and deletion.
- **Unsorted Slice**: Evaluates the performance of slices in Go for similar operations.
- **Linked List**: Assesses the efficiency of singly linked lists.
- **Double Linked List**: Compares the performance of doubly linked lists against singly linked lists.
- **Map (Hash Table)**: Analyzes Go's implementation of hash tables.
- **Binary Tree**: Tests the basic binary tree implementation for add/remove operations.

## Test Methodology

Each data structure undergoes a series of tests to measure the time taken for:

- **Adding Records**: Timed tests for inserting data into the data structure.
- **Removing Records**: Timed tests for deleting data from the data structure.

These tests are designed to isolate the add/remove operations, ensuring that setup or preparatory steps do not influence the timing results.

## Running the Tests

To run the tests, ensure you have Go installed and execute the main program. This will automatically perform all the tests and output the results to the console.

```bash
go run main.go
```

## Results Interpretation

The test results include the following metrics for each data structure:

- **Mean**: The average time taken for the operation.
- **Median**: The middle value in the set of times.
- **Min/Max**: The minimum and maximum times observed.
- **95th Percentile**: The value below which 95% of the operation times fall.
- **Throughput**: The number of operations that can be completed per second.
- **Standard Deviation**: The amount of variance or dispersion in the operation times.
- **Total Duration**: The total time taken for all test iterations.

These metrics provide a comprehensive view of each data structure's performance, allowing for informed decisions when choosing the appropriate structure for a particular application.
