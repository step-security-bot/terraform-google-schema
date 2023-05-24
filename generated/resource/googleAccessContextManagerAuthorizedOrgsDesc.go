package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerAuthorizedOrgsDesc = `{
  "block": {
    "attributes": {
      "asset_type": {
        "description": "The type of entities that need to use the authorization relationship during\nevaluation, such as a device. Valid values are \"ASSET_TYPE_DEVICE\" and\n\"ASSET_TYPE_CREDENTIAL_STRENGTH\". Possible values: [\"ASSET_TYPE_DEVICE\", \"ASSET_TYPE_CREDENTIAL_STRENGTH\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorization_direction": {
        "description": "The direction of the authorization relationship between this organization\nand the organizations listed in the \"orgs\" field. The valid values for this\nfield include the following:\n\nAUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic\nin the organizations listed in the 'orgs' field.\n\nAUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the 'orgs'\nfield to evaluate the traffic in this organization.\n\nFor the authorization relationship to take effect, all of the organizations\nmust authorize and specify the appropriate relationship direction. For\nexample, if organization A authorized organization B and C to evaluate its\ntraffic, by specifying \"AUTHORIZATION_DIRECTION_TO\" as the authorization\ndirection, organizations B and C must specify\n\"AUTHORIZATION_DIRECTION_FROM\" as the authorization direction in their\n\"AuthorizedOrgsDesc\" resource. Possible values: [\"AUTHORIZATION_DIRECTION_TO\", \"AUTHORIZATION_DIRECTION_FROM\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorization_type": {
        "description": "A granular control type for authorization levels. Valid value is \"AUTHORIZATION_TYPE_TRUST\". Possible values: [\"AUTHORIZATION_TYPE_TRUST\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Time the AuthorizedOrgsDesc was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Resource name for the 'AuthorizedOrgsDesc'. Format:\n'accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}'.\nThe 'authorized_orgs_desc' component must begin with a letter, followed by\nalphanumeric characters or '_'.\nAfter you create an 'AuthorizedOrgsDesc', you cannot change its 'name'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "orgs": {
        "description": "The list of organization ids in this AuthorizedOrgsDesc.\nFormat: 'organizations/\u003corg_number\u003e'\nExample: 'organizations/123456'",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "parent": {
        "description": "Required. Resource name for the access policy which owns this 'AuthorizedOrgsDesc'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the AuthorizedOrgsDesc was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAccessContextManagerAuthorizedOrgsDescSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerAuthorizedOrgsDesc), &result)
	return &result
}
