package main

import (
	"errors"
	"fmt"
)

func main() {
	_ = bar()
	_ = baz("baz")
	_ = foo()
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
