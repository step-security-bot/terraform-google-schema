package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleComputeTargetPoolSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleComputeTargetPoolSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
