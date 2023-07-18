package goSecretSharing

import (
	"fmt"
	"testing"
)

func TestSharingAndRecover(t *testing.T) {
	secretText := "Hello World"
	noOfShares := 5
	noRequiredShares := 2
	splitedStrings, tErr := SplitSecret(secretText, noOfShares, noRequiredShares)
	if tErr != nil {
		t.Errorf("splitting Secret Failed... %v", tErr)
	}
	// recover splitedStrings
	recString := []string{splitedStrings[1], splitedStrings[2]}
	recoveredSecret, eErr := RecoverSecret(recString)
	if eErr != nil {
		t.Errorf("splitting Secret Failed... %v", tErr)
	}
	fmt.Println("test successfull..", recoveredSecret)
}
