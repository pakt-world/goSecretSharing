package secretsharing

import (
	"crypto/rand"
	"errors"
)

var ErrInvalidThreshold = errors.New("Invalid Threshold")
var ErrInvalidCount = errors.New("Invalid Count")

func Split(n, k byte, secret []byte) (map[byte][]byte, error) {
	if k <= 1 {
		return nil, ErrInvalidThreshold
	}

	if n < k {
		return nil, ErrInvalidCount
	}

	shares := make(map[byte][]byte, n)

	for _, b := range secret {
		p, err := generate(k-1, b, rand.Reader)
		if err != nil {
			return nil, err
		}

		for x := byte(1); x <= n; x++ {
			shares[x] = append(shares[x], eval(p, x))
		}
	}

	return shares, nil
}

func Combine(shares map[byte][]byte) []byte {
	var secret []byte
	for _, v := range shares {
		secret = make([]byte, len(v))
		break
	}

	points := make([]pair, len(shares))
	for i := range secret {
		p := 0
		for k, v := range shares {
			points[p] = pair{x: k, y: v[i]}
			p++
		}
		secret[i] = interpolate(points, 0)
	}

	return secret
}
