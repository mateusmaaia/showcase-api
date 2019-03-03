package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	sourceURL := "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"
	seed := &Seed{sourceURL: sourceURL}
	realEstates, _ := seed.Import()
	assert.NotEmpty(t, realEstates)
}
