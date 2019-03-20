package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportOK(t *testing.T) {
	sourceURL := "https://showcase-api-beea4.firebaseapp.com/source.json"
	seed := &Seed{SourceURL: sourceURL}
	realEstates, _ := seed.Import()
	assert.NotEmpty(t, realEstates)
}

func TestImportFail(t *testing.T) {
	sourceURL := ""
	seed := &Seed{SourceURL: sourceURL}
	realEstates, _ := seed.Import()
	assert.Empty(t, realEstates)
}
