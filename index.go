package goSecretSharing

import (
	"encoding/hex"
	"errors"
	"fmt"

	secretsharing "github.com/pakt-world/goSecretSharing/shamir"
)

var KError = "Cannot do this, as requiredShares greater then noOfShares"

// SplitSecret splits plainText into noOfShares shares, any requiredShares of which recover it.
// The map key is the share's index (x coordinate, 1..noOfShares). Recovery needs that index:
// keep it with the share, and pass both to RecoverSecretIndexed.
func SplitSecret(plainText string, noOfShares, requiredShares int) (map[byte]string, error) {
	if requiredShares > noOfShares {
		return nil, errors.New(KError)
	}
	if requiredShares < 2 || noOfShares > 255 {
		return nil, errors.New("requiredShares must be at least 2 and noOfShares at most 255")
	}
	var shareData = make(map[byte]string, noOfShares)
	shares, err := secretsharing.Split(byte(noOfShares), byte(requiredShares), []byte(plainText))
	if err != nil {
		return nil, err
	}
	for x, y := range shares {
		shareData[x] = hex.EncodeToString(y)
	}
	return shareData, nil
}

// RecoverSecretIndexed recovers the secret from shares keyed by their index, as SplitSecret
// returned them. Any requiredShares of the shares, in any order, give the secret.
func RecoverSecretIndexed(shares map[byte]string) (string, error) {
	if len(shares) < 2 {
		return "", errors.New("need at least 2 shares")
	}
	subset := make(map[byte][]byte, len(shares))
	length := -1
	for x, y := range shares {
		if x == 0 {
			return "", errors.New("share index 0 is invalid")
		}
		dY, err := hex.DecodeString(y)
		if err != nil {
			return "", fmt.Errorf("share %d is not valid hex: %w", x, err)
		}
		if length >= 0 && len(dY) != length {
			return "", errors.New("shares have different lengths")
		}
		length = len(dY)
		subset[x] = dY
	}
	return string(secretsharing.Combine(subset)), nil
}

// RecoverSecret recovers the secret from shares given as a list.
//
// Deprecated: the list position is used as the share index (first share = index 1, and so
// on), so this is correct only when the list holds share 1, share 2, ... share n in that order.
// It happens to work for every ordering when exactly the three shares of a 3-of-3 split are
// passed (in GF(256) every recovery weight for the indices {1, 2, 3} is 1), but any threshold
// below the share count, or any subset, silently returns the wrong secret. Use
// RecoverSecretIndexed.
func RecoverSecret(shares []string) (string, error) {
	indexed := make(map[byte]string, len(shares))
	for i, y := range shares {
		indexed[byte(i+1)] = y
	}
	return RecoverSecretIndexed(indexed)
}
