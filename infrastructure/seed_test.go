package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportOK(t *testing.T) {
	sourceURL := "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
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
