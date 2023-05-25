package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-google-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestGoogleOrganizationIamMemberSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.GoogleOrganizationIamMemberSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
