package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleKmsKeyRingSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleKmsKeyRingSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
