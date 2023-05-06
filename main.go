package main

import (
	"docman/cmd"
	"docman/pkg/log"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Error(fmt.Sprintf("Application panic: %v", err))
		}
	}()
	cmd.Execute()
}
