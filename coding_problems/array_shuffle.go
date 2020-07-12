package main

import "fmt"

func main() {
	test := []int{1, 7, 3, 2, 5, 4}
	n := 3
	shuffle(test, n)
}
func shuffle(nums []int, n int) []int {
	part1 := nums[:n]
	part2 := nums[n:]

	fmt.Println(part1)
	fmt.Println(part2)

	var result []int

	var int1 int
	var int2 int

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			result = append(result, part1[int1])
			int1 += 1
		} else {
			result = append(result, part2[int2])
			int2 += 1
		}
	}
	return result
}
