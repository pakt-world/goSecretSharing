package main

import (
	"fmt"
	"os"
	cmd "pakt-secret-sharing/cmd"
)

func main() {
	fmt.Println("PAKT-Secret-Sharing....")
	method := os.Args[1]
	cmd.Execute(method)
}
