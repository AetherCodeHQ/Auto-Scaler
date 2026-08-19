package main

import (
	"fmt"
	"os"
)

// auto_scaler - Auto scaling based on load
func auto_scaler(path string) {
	fmt.Println("========================================")
	fmt.Println("  Auto-Scaler")
	fmt.Println("  Auto scaling based on load")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	auto_scaler(path)
}
