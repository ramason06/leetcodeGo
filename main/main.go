package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums1, 2))
	fmt.Println(nums1)
}

func removeElement(nums []int, val int) int {
	i, k := 0, 0

	for i < len(nums) {
		if nums[i] != val {
			i++
			k++
		} else if i < len(nums)-1 {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return k
}
