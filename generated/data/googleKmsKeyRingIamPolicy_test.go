package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleKmsKeyRingIamPolicySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleKmsKeyRingIamPolicySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
