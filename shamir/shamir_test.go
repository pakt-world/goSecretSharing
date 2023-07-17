package secretsharing

import (
	"fmt"
)

func TestSharing() {
	secret := "hello world!" // our secret
	n := byte(5)             // create 5 shares
	k := byte(2)             // require 2 of them to combine

	shares, err := Split(n, k, []byte(secret)) // split into 30 shares
	if err != nil {
		fmt.Println(err)
		return
	}

	// select a random subset of the total shares
	subset := make(map[byte][]byte, k)
	for x, y := range shares { // just iterate since maps are randomized
		subset[x] = y
		if len(subset) == int(k) {
			break
		}
	}

	// combine two shares and recover the secret
	recovered := string(Combine(subset))
	fmt.Println(recovered)

	// Output: hello world!
}
