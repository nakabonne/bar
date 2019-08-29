package main

import "fmt"

type A struct {
	Name string
}

func main() {
	func() {
		var a A
		// TODO: Add more option b
		fmt.Println(a)

		TODO := ""
		fmt.Println(TODO)
		fmt.Println("remove")
		fmt.Println("oi")
	}()
	_, _ = buzz()
}

func bar() ([]int, error) {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	return nums[1:], nil
}

func buzz() ([]int, error) {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	return nums[1:], nil
}
