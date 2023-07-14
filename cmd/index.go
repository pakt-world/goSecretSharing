package cmd

import (
	"fmt"
	"os"
	methods "pakt-secret-sharing/methods"
	utils "pakt-secret-sharing/utils"
)

func Execute(method string) {
	argsWithoutProg := os.Args[1:]

	switch method {
	// Split Secret into shares and return
	case "split":
		// return utils.GetShares()
		if len(argsWithoutProg) < 3 {
			fmt.Println("Error::: Insufficent arguments to split secret")
			return
		}
		plainText := argsWithoutProg[1]
		noOfShares := utils.StringToInt(argsWithoutProg[2])
		minNoShares := utils.StringToInt(argsWithoutProg[3])
		shares := methods.SplitSecret(plainText, noOfShares, minNoShares)
		// fmt.Println("all shares===", shares)
		for x, y := range shares {
			fmt.Printf("Share:\t%d\t%s\n", x, y)
		}
		return
	// Recover Secret from Shares
	case "recover":
		shares := argsWithoutProg[1:]
		methods.RecoverSecret(shares)
		return
	default:
		fmt.Println("Invalid Method Passed...")
		return
	}
}
