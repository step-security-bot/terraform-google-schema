package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAppEngineServiceSplitTrafficSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleAppEngineServiceSplitTrafficSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
