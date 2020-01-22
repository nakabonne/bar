package main

import (
	"errors"
	"fmt"
)

func main() {
	_ = bar()
	_ = baz("baz")
	_ = foo()
	_ = added()
	_ = isOne(2)
}

type Bar struct {
}

func bar() error {
	return nil
}

func baz(s string) string {
	return ""
}

func foo() error {
	msg := "foo"
	return errors.New(fmt.Sprintf("error: %s", msg))
}

func added() error {
	msg := "added"
	return errors.New(fmt.Sprintf("error: %s", msg))
}

func isOne(i int) bool {
	return i == 1
}
