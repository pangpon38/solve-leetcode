package main

import "fmt"

/**
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]

Output: false

Explanation:

All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]

Output: true
*/

func main() {

	testCase := []int{1, 2, 3, 1}
	result := containsDuplicate(testCase)
	fmt.Println("result", result)
}

func containsDuplicate(nums []int) bool {

	hashMap := make(map[int]int)

	for _, val := range nums {
		if _, isFound := hashMap[val]; isFound {
			return true
		}
		hashMap[val] = 1
	}

	return false
}
