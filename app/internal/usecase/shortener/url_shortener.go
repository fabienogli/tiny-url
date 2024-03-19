package shortener

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
)

type Sha256LinkGenerator struct{}

func NewSha256LinkGenerator() *Sha256LinkGenerator {
	return &Sha256LinkGenerator{}
}

func (s *Sha256LinkGenerator) GenerateShortLink(initialLink string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(initialLink))
	urlHashBytes := algorithm.Sum(nil)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base64.RawURLEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%d", generatedNumber)),
	)
	return finalString[:8]
}
