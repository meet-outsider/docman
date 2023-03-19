/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"docman/cmd/docman"
	"docman/config"
	"fmt"
)

func main() {
	// 启动docman应用
	docman.Execute()
	var server = config.Config.Server
	fmt.Println(server)
}
