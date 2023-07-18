package goSecretSharing

import (
	"encoding/hex"
	"errors"

	secretsharing "github.com/pakt-world/goSecretSharing/shamir"
)

var KError = "Cannot do this, as requiredShares greater then noOfShares"

func SplitSecret(plainText string, noOfShares, requiredShares int) (map[byte]string, error) {
	if requiredShares > noOfShares {
		return nil, errors.New(KError)
	}
	var shareData = make(map[byte]string, noOfShares)
	message := []byte(plainText)
	n := byte(noOfShares)
	k := byte(requiredShares)
	shares, _ := secretsharing.Split(n, k, []byte(message))
	for x, y := range shares {
		shareData[x] = hex.EncodeToString(y)
	}
	return shareData, nil
}

func RecoverSecret(shares []string) (string, error) {
	// get length of shares
	lenShares := len(shares)
	subset := make(map[byte][]byte, lenShares)
	for x, y := range shares {
		xh := uint8(x + 1)
		dY, _ := hex.DecodeString(y)
		subset[xh] = dY
		if len(subset) == int(lenShares) {
			break
		}
	}
	reconstructed := string(secretsharing.Combine(subset))
	return reconstructed, nil
}
