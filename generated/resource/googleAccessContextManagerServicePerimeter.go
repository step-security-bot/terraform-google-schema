package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAccessContextManagerServicePerimeter = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the AccessPolicy was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the ServicePerimeter and its use. Does not affect\nbehavior.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Resource name for the ServicePerimeter. The short_name component must\nbegin with a letter and only include alphanumeric and '_'.\nFormat: accessPolicies/{policy_id}/servicePerimeters/{short_name}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The AccessPolicy this ServicePerimeter lives in.\nFormat: accessPolicies/{policy_id}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "perimeter_type": {
        "description": "Specifies the type of the Perimeter. There are two types: regular and\nbridge. Regular Service Perimeter contains resources, access levels,\nand restricted services. Every resource can be in at most\nONE regular Service Perimeter.\n\nIn addition to being in a regular service perimeter, a resource can also\nbe in zero or more perimeter bridges. A perimeter bridge only contains\nresources. Cross project operations are permitted if all effected\nresources share some perimeter (whether bridge or regular). Perimeter\nBridge does not contain access levels or services: those are governed\nentirely by the regular perimeter that resource is in.\n\nPerimeter Bridges are typically useful when building more complex\ntopologies with many independent perimeters that need to share some data\nwith a common perimeter, but should not be able to share data among\nthemselves.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "title": {
        "description": "Human readable title. Must be unique within the Policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Time the AccessPolicy was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "status": {
        "block": {
          "attributes": {
            "access_levels": {
              "description": "A list of AccessLevel resource names that allow resources within\nthe ServicePerimeter to be accessed from the internet.\nAccessLevels listed must be in the same policy as this\nServicePerimeter. Referencing a nonexistent AccessLevel is a\nsyntax error. If no AccessLevel names are listed, resources within\nthe perimeter can only be accessed via GCP calls with request\norigins within the perimeter. For Service Perimeter Bridge, must\nbe empty.\n\nFormat: accessPolicies/{policy_id}/accessLevels/{access_level_name}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resources": {
              "description": "A list of GCP resources that are inside of the service perimeter.\nCurrently only projects are allowed.\nFormat: projects/{project_number}",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "restricted_services": {
              "description": "GCP services that are subject to the Service Perimeter\nrestrictions. Must contain a list of services. For example, if\n'storage.googleapis.com' is specified, access to the storage\nbuckets inside the perimeter must meet the perimeter's access\nrestrictions.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleAccessContextManagerServicePerimeterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAccessContextManagerServicePerimeter), &result)
	return &result
}
