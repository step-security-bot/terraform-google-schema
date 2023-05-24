package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestGoogleComputeRouterStatusSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.GoogleComputeRouterStatusSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
