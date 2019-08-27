package main

import "fmt"

type A struct {
	Name string
}

func main() {
	func() {
		// TODO: Add more option b
		fmt.Println("vim-go")

		TODO := ""
		fmt.Println(TODO)
		fmt.Println("remove")
		fmt.Println("oi")
	}()
	_ = bar()
}

func bar() error {
	return nil
}
