package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleContainerAttachedInstallManifestSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleContainerAttachedInstallManifestSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
