package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleIapWebTypeAppEngineIamBindingSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleIapWebTypeAppEngineIamBindingSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
