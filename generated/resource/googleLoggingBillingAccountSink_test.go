package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleLoggingBillingAccountSinkSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleLoggingBillingAccountSinkSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
