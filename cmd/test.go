package goSecretSharing

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("CMD GO-Secret-Sharing....")
	if len(os.Args) < 2 {
		fmt.Println("Invalid Argument Passed")
		os.Exit(1)
	}
	method := os.Args[1]
	Execute(method)
}
