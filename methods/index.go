package methods

import (
	"encoding/hex"
	"fmt"
	"os"

	secretsharing "github.com/pakt/go-secret-sharing/secretSharing"
)

func SplitSecret(plainText string, noOfShares, minNum int) map[byte]string {
	var shareData = make(map[byte]string, noOfShares)
	fmt.Printf("Splitting Secret into %d Shares...\n", noOfShares)
	message := []byte(plainText)
	n := byte(noOfShares)
	k := byte(minNum)
	if minNum > noOfShares {
		fmt.Printf("Cannot do this, as k greater than n")
		os.Exit(0)
	}
	shares, _ := secretsharing.Split(n, k, []byte(message))
	for x, y := range shares {
		shareData[x] = hex.EncodeToString(y)
	}
	return shareData
}

func RecoverSecret(shares []string) string {
	// get length of shares
	lenShares := len(shares)
	subset := make(map[byte][]byte, lenShares)
	for x, y := range shares {
		fmt.Printf("Share:\t%d\t%s\n", x, y)
		xh := uint8(x + 1)
		dY, _ := hex.DecodeString(y)
		subset[xh] = dY
		if len(subset) == int(lenShares) {
			break
		}
	}
	reconstructed := string(secretsharing.Combine(subset))
	fmt.Printf("\nReconstructed: %s\n", reconstructed)
	return reconstructed
}
