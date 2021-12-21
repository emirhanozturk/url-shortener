package shortener

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortUrl(t *testing.T) {
	initialLink1 := "https://www.google.com/"
	shortLink1 := GenerateShortUrl(initialLink1)
	fmt.Println(shortLink1)

	initialLink2 := "https://github.com/"
	shortLink2 := GenerateShortUrl(initialLink2)
	fmt.Println(shortLink2)

	initialLink3 := "https://www.youtube.com/"
	shortLink3 := GenerateShortUrl(initialLink3)
	fmt.Println(shortLink3)

	assert.Equal(t, shortLink1, "UAQV6Wq3")
	assert.Equal(t, shortLink2, "JWmeg4AR")
	assert.Equal(t, shortLink3, "8HG9Zy1V")

}
