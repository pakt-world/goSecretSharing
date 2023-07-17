package goSecretSharing

func TestSharingAndRecover() {
	secretText := "Hello World"
	noOfShares := 5
	noRequiredShares := 2
	splitedStrings := SplitSecret(secretText, noOfShares, noRequiredShares)
	// recover splitedStrings
	recString := []string{splitedStrings[1], splitedStrings[2]}
	recoveredSecret := RecoverSecret(recString)
	return
}
