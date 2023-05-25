package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleActiveDirectoryDomain = `{
  "block": {
    "attributes": {
      "admin": {
        "description": "The name of delegated administrator account used to perform Active Directory operations. \nIf not specified, setupadmin will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "authorized_networks": {
        "description": "The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.\nIf CIDR subnets overlap between networks, domain creation will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "domain_name": {
        "description": "The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions, \nhttps://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fqdn": {
        "computed": true,
        "description": "The fully-qualified domain name of the exposed domain used by clients to connect to the service. \nSimilar to what would be chosen for an Active Directory set up on an internal network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Resource labels that can contain user-provided metadata",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "locations": {
        "description": "Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/] \ne.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The unique name of the domain using the format: 'projects/{project}/locations/global/domains/{domainName}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reserved_ip_range": {
        "description": "The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. \nRanges must be unique and non-overlapping with existing subnets in authorizedNetworks",
        "description_kind": "plain",
        "required": true,
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

func GoogleActiveDirectoryDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleActiveDirectoryDomain), &result)
	return &result
}
