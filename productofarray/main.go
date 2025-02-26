package main

import "fmt"

/*
*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
*/
func main() {

	testCase := []int{1, 2, 3, 4}
	result := productOfArray(testCase)
	fmt.Println("result", result)
}

func productOfArray(nums []int) []int {
	result := make([]int, len(nums))
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))

	prefix[0] = 1

	for index := 1; index < len(nums); index++ {
		prefix[index] = prefix[index-1] * nums[index-1]
	}

	postfix[len(nums)-1] = 1

	for index := len(nums) - 2; index >= 0; index-- {
		postfix[index] = postfix[index+1] * nums[index+1]
	}

	for index := 0; index < len(nums); index++ {
		result[index] = prefix[index] * postfix[index]
	}

	return result
}
