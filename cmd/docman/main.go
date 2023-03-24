package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("app exit panic cause:", err)
		}
	}()
	Execute()
}
