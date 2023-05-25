package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleComputeTargetGrpcProxySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleComputeTargetGrpcProxySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
