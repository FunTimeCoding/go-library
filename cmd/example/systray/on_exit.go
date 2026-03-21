//go:build local

package main

import "fmt"

func onExit() {
	fmt.Println("onExit")
}
