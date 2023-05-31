package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleComputeDiskIamPolicySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleComputeDiskIamPolicySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
