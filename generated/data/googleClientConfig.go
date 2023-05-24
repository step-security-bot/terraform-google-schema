package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClientConfig = `{
  "block": {
    "attributes": {
      "access_token": {
        "computed": true,
        "description": "The OAuth2 access token used by the client to authenticate against the Google Cloud API.",
        "description_kind": "markdown",
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The ID of this data source in Terraform state. It is created in a projects/{{project}}/regions/{{region}}/zones/{{zone}} format and is NOT used by the data source in requests to Google APIs.",
        "description_kind": "markdown",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project to apply any resources to.",
        "description_kind": "markdown",
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region to operate under.",
        "description_kind": "markdown",
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "The zone to operate under.",
        "description_kind": "markdown",
        "type": "string"
      }
    },
    "description": "Use this data source to access the configuration of the Google Cloud provider.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func GoogleClientConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClientConfig), &result)
	return &result
}
