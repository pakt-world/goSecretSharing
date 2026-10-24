package secretsharing

import "testing"

func TestSharing(t *testing.T) {
	secret := "hello world!"
	n := byte(5) // create 5 shares
	k := byte(2) // require 2 of them to combine

	shares, err := Split(n, k, []byte(secret))
	if err != nil {
		t.Fatal(err)
	}

	// any k shares, keyed by their own index, recover the secret
	for _, ids := range [][]byte{{1, 2}, {4, 2}, {5, 3}} {
		subset := map[byte][]byte{ids[0]: shares[ids[0]], ids[1]: shares[ids[1]]}
		if got := string(Combine(subset)); got != secret {
			t.Fatalf("shares %v: got %q", ids, got)
		}
	}
}
