package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func hashUrlSha256(url string) []byte {
	algoritm := sha256.New()
	algoritm.Write([]byte(url))
	return algoritm.Sum(nil)
}

func base58Encoded(bytes []byte) string {

	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode([]byte(bytes))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortUrl(initialUrl string) string {
	hashedUrlBytes := hashUrlSha256(initialUrl)
	generatedNumber := new(big.Int).SetBytes(hashedUrlBytes).Uint64()
	generatedShortUrl := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return generatedShortUrl[:8]
}
