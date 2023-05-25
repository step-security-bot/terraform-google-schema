package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleStorageObjectSignedUrlSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleStorageObjectSignedUrlSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
