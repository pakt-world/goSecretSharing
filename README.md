# goSecretSharing

go-secret-sharing is a Golang package that implements Shamir's Secret Sharing (SSS) algorithm, specifically designed to work with mnemonics and string secrets. It allows you to split a secret into n number of shares such that a specified number (t) of those shares are required to recover the original secret.

The design of this tool was inspired by Ava Labs' [mnemonic-shamir-secret-sharing-cli](https://github.com/ava-labs/mnemonic-shamir-secret-sharing-cli).

# Features

- Implements Shamir's Secret Sharing algorithm for string secrets and mnemonics.
- Customize the number of shares and threshold required to reconstruct the secret.
- Provides cryptographic security for the shares generated.
- Optimized for performance and scalability.

## Installation

```go
  go get github.com/pakt/goSecretSharing
```

## Usage

```go
  import (
    "fmt"

    goSecretSharing "github.com/pakt/goSecretSharing"
  )

  func main(){
    secretText := "Hello World"
    noOfShares := 5
    noRequiredShares := 2
    splitedStrings, sErr := goSecretSharing.SplitSecret(secretText, noOfShares, noRequiredShares)
    if sErr != nil {
      fmt.Println("Errorpopp===", sErr)
    }
    // recover splitedStrings
    recString := []string{splitedStrings[1], splitedStrings[2]}
    recoveredSecret, rErr := goSecretSharing.RecoverSecret(recString)
    if rErr != nil {
      fmt.Println("Errorpopp===", rErr)
    }
    fmt.Println("recovered==", recoveredSecret)
  }
```

## Testing

## Contributing

We welcome contributions from the open-source community. Please fork the repository and submit a pull request to contribute to this project.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

- Thanks to Adi Shamir for the SSS algorithm.
- Thanks to Ava Labs for inspiring this project.

## Contact

If you have any questions or issues, please open an issue on this repository, or contact the maintainer directly.
