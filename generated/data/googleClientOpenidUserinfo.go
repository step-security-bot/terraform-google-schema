package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleClientOpenidUserinfo = `{
  "block": {
    "attributes": {
      "email": {
        "computed": true,
        "description": "The email of the account used by the provider to authenticate with GCP.",
        "description_kind": "markdown",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The ID of this data source in Terraform state. Its value is the same as the ` + "`" + `email` + "`" + ` attribute. Do not use this field, use the ` + "`" + `email` + "`" + ` attribute instead.",
        "description_kind": "markdown",
        "type": "string"
      }
    },
    "description": "Get OpenID userinfo about the credentials used with the Google provider, specifically the email.\nThis datasource enables you to export the email of the account you've authenticated the provider with; this can be used alongside data.google_client_config's access_token to perform OpenID Connect authentication with GKE and configure an RBAC role for the email used.\n\n~\u003e This resource will only work as expected if the provider is configured to use the https://www.googleapis.com/auth/userinfo.email scope! You will receive an error otherwise. The provider uses this scope by default.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func GoogleClientOpenidUserinfoSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleClientOpenidUserinfo), &result)
	return &result
}
