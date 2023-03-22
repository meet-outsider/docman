/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"docman/cmd/docman"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("app exit panic cause:", err)
		}
	}()
	docman.Execute()
}
