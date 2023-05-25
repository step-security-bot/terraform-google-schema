package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudAssetOrganizationFeed = `{
  "block": {
    "attributes": {
      "asset_names": {
        "description": "A list of the full names of the assets to receive updates. You must specify either or both of \nassetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are\nexported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.\nSee https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "asset_types": {
        "description": "A list of types of the assets to receive updates. You must specify either or both of assetNames\nand assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to\nthe feed. For example: \"compute.googleapis.com/Disk\"\nSee https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all\nsupported asset types.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "billing_project": {
        "description": "The project whose identity will be used when sending messages to the\ndestination pubsub topic. It also specifies the project for API \nenablement check, quota, and billing.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_type": {
        "description": "Asset content type. If not specified, no content but the asset name and type will be returned. Possible values: [\"CONTENT_TYPE_UNSPECIFIED\", \"RESOURCE\", \"IAM_POLICY\", \"ORG_POLICY\", \"ACCESS_POLICY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "feed_id": {
        "description": "This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The format will be organizations/{organization_number}/feeds/{client-assigned_feed_identifier}.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The organization this feed should be created in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "condition": {
        "block": {
          "attributes": {
            "description": {
              "description": "Description of the expression. This is a longer text which describes the expression,\ne.g. when hovered over it in a UI.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description": "Textual representation of an expression in Common Expression Language syntax.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "location": {
              "description": "String indicating the location of the expression for error reporting, e.g. a file \nname and a position in the file.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "title": {
              "description": "Title for the expression, i.e. a short string describing its purpose.\nThis can be used e.g. in UIs which allow to enter the expression.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "A condition which determines whether an asset update should be published. If specified, an asset\nwill be returned only when the expression evaluates to true. When set, expression field\nmust be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with\nexpression \"temporal_asset.deleted == true\" will only publish Asset deletions. Other fields of\ncondition are optional.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "feed_output_config": {
        "block": {
          "block_types": {
            "pubsub_destination": {
              "block": {
                "attributes": {
                  "topic": {
                    "description": "Destination on Cloud Pubsub topic.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Destination on Cloud Pubsub.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Output configuration for asset feed destination.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
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

func GoogleCloudAssetOrganizationFeedSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudAssetOrganizationFeed), &result)
	return &result
}
