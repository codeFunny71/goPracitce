package main

import "fmt"

func main() {
	test := []int{1, 7, 3, 2, 5, 4}
	test2 := []int{1, 1, 3, 3, 1, 4}

	n := 3
	fmt.Println("original")
	fmt.Println(test)
	fmt.Println("shuffled")
	test = shuffle2(test, n)

	fmt.Println(test)
	fmt.Println("summed")
	fmt.Println(runningSum(test))
	fmt.Println(numIdenticalPairs(test2))

}

func shuffle2(nums []int, n int) []int {
	intsResult := make([]int, 0)
	i := 0
	for {
		intsResult = append(intsResult, nums[i])
		intsResult = append(intsResult, nums[i+n])
		if (i+1)*2 >= len(nums) {
			break
		} else {
			i++
		}
	}
	return intsResult
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count += 1
			}
		}
	}
	return count
}
