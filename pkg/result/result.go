package main

import (
	"errors"
	"fmt"
)

type Result struct {
	value interface{}
	err   error
}

func Ok(value interface{}) Result {
	return Result{value, nil}
}

func (r Result) Err() Result {
	return Err(r.err)
}

func Err(err error) Result {
	return Result{nil, err}
}
func (r Result) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func (r Result) Expect(msg string) interface{} {
	if r.err != nil {
		panic(msg + ": " + r.err.Error())
	}
	return r.value
}

func (r Result) IsErr() bool {
	return r.err != nil
}
func divide(a, b float64) Result {
	if b == 0 {
		return Err(errors.New("division by zero"))
	}
	return Ok(a / b)
}
func main() {
	result := divide(10, 0)
	if result.IsErr() {
		fmt.Println("Error:", result.Err())
	} else {
		fmt.Println("Result:", result.Unwrap())
	}
}
