package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {

	initialLink := "https://www.google.com/"
	shortUrl := "UAQV6Wq3"

	SaveUrlMapping(shortUrl, initialLink)

	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)

}
