package goSecretSharing

import "testing"

const secret = "correct horse battery staple"

func TestIndexedRecoveryWorksForAnySubsetAndOrder(t *testing.T) {
	shares, err := SplitSecret(secret, 5, 3)
	if err != nil {
		t.Fatal(err)
	}
	subsets := [][]byte{{1, 2, 3}, {3, 1, 5}, {2, 4, 5}, {5, 4, 3, 2, 1}}
	for _, ids := range subsets {
		sub := map[byte]string{}
		for _, id := range ids {
			sub[id] = shares[id]
		}
		got, err := RecoverSecretIndexed(sub)
		if err != nil || got != secret {
			t.Fatalf("subset %v: got %q, err %v", ids, got, err)
		}
	}
}

func TestThreeOfThreeRecoversInAnyOrder(t *testing.T) {
	// The configuration the KMS uses today: positional recovery works for 3-of-3 in any order.
	shares, _ := SplitSecret(secret, 3, 3)
	for _, order := range [][]byte{{1, 2, 3}, {2, 3, 1}, {3, 1, 2}} {
		got, err := RecoverSecret([]string{shares[order[0]], shares[order[1]], shares[order[2]]})
		if err != nil || got != secret {
			t.Fatalf("order %v: got %q, err %v", order, got, err)
		}
	}
}

func TestPositionalRecoveryIsWrongForASubset(t *testing.T) {
	// Documents the deprecated behaviour: shares 2 and 3 of a 2-of-3 split, passed as a list,
	// are treated as indices 1 and 2 and give the wrong secret.
	shares, _ := SplitSecret(secret, 3, 2)
	got, _ := RecoverSecret([]string{shares[2], shares[3]})
	if got == secret {
		t.Fatal("positional recovery of a subset unexpectedly returned the secret")
	}
	ok, err := RecoverSecretIndexed(map[byte]string{2: shares[2], 3: shares[3]})
	if err != nil || ok != secret {
		t.Fatalf("indexed recovery of the same subset failed: %q, %v", ok, err)
	}
}

func TestBadInputIsRejected(t *testing.T) {
	if _, err := SplitSecret(secret, 3, 4); err == nil {
		t.Fatal("requiredShares > noOfShares accepted")
	}
	if _, err := SplitSecret(secret, 3, 1); err == nil {
		t.Fatal("requiredShares 1 accepted")
	}
	if _, err := RecoverSecretIndexed(map[byte]string{1: "zz", 2: "00"}); err == nil {
		t.Fatal("invalid hex accepted")
	}
	if _, err := RecoverSecretIndexed(map[byte]string{1: "0011"}); err == nil {
		t.Fatal("a single share accepted")
	}
	if _, err := RecoverSecretIndexed(map[byte]string{0: "00", 1: "01"}); err == nil {
		t.Fatal("index 0 accepted")
	}
}
