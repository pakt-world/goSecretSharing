package main

import (
	"fmt"
	"os"
	cmd "pakt-secret-sharing/cmd"
)

func main() {
	fmt.Println("GO-Secret-Sharing....")
	if len(os.Args) < 2 {
		fmt.Println("Invalid Argument Passed")
		os.Exit(1)
	}
	method := os.Args[1]
	cmd.Execute(method)
}
