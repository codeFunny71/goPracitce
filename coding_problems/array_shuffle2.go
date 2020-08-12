package main

import "fmt"

func main() {
	test := []int{1, 7, 3, 2, 5, 4}

	n := 3
	fmt.Println("original")
	fmt.Println(test)
	fmt.Println("shuffled")
	test = shuffle2(test, n)

	fmt.Println(test)
	fmt.Println("summed")
	fmt.Println(runningSum(test))

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
