package main

import "sort"

func TaskAssignment(k int, tasks []int) [][]int {
	pairedTasks := make([][]int, 0)
	tasksMap := mapTaskstoIdx(tasks)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})

	var task1Index, task2Index int
	for idx := 0; idx < k; idx++ {
		task1Duration := tasks[idx]
		indicesWithTask1Durtaion := tasksMap[task1Duration]
		task1Index, tasksMap[task1Duration] =
			indicesWithTask1Durtaion[len(indicesWithTask1Durtaion)-1], indicesWithTask1Durtaion[:len(indicesWithTask1Durtaion)-1]

		task2SortedIndex := len(tasks) - 1 - idx
		task2Duration := tasks[task2SortedIndex]
		indicesWithTask2Durtaion := tasksMap[task2Duration]
		task2Index, tasksMap[task2Duration] =
			indicesWithTask2Durtaion[len(indicesWithTask2Durtaion)-1], indicesWithTask2Durtaion[:len(indicesWithTask2Durtaion)-1]

		pairedTasks = append(pairedTasks, []int{task1Index, task2Index})
	}

	return pairedTasks

}

func mapTaskstoIdx(tasks []int) map[int][]int {
	tasksMap := map[int][]int{}
	for idx := range tasks {
		taskDuration := tasks[idx]
		tasksMap[taskDuration] = append(tasksMap[taskDuration], idx)
	}
	return tasksMap
}
