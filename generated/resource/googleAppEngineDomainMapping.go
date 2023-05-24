package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAppEngineDomainMapping = `{
  "block": {
    "attributes": {
      "domain_name": {
        "description": "Relative name of the domain serving the application. Example: example.com.",
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
        "description": "Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.",
        "description_kind": "plain",
        "type": "string"
      },
      "override_strategy": {
        "description": "Whether the domain creation should override any existing mappings for this domain.\nBy default, overrides are rejected. Default value: \"STRICT\" Possible values: [\"STRICT\", \"OVERRIDE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_records": {
        "computed": true,
        "description": "The resource records required to configure this domain mapping. These records must be added to the domain's DNS\nconfiguration in order to serve the application via this domain mapping.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "rrdata": "string",
              "type": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "ssl_settings": {
        "block": {
          "attributes": {
            "certificate_id": {
              "computed": true,
              "description": "ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will\nremove SSL support.\nBy default, a managed certificate is automatically created for every domain mapping. To omit SSL support\nor to configure SSL manually, specify 'SslManagementType.MANUAL' on a 'CREATE' or 'UPDATE' request. You must be\nauthorized to administer the 'AuthorizedCertificate' resource to manually map it to a DomainMapping resource.\nExample: 12345.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pending_managed_certificate_id": {
              "computed": true,
              "description": "ID of the managed 'AuthorizedCertificate' resource currently being provisioned, if applicable. Until the new\nmanaged certificate has been successfully provisioned, the previous SSL state will be preserved. Once the\nprovisioning process completes, the 'certificateId' field will reflect the new managed certificate and this\nfield will be left empty. To remove SSL support while there is still a pending managed certificate, clear the\n'certificateId' field with an update request.",
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_management_type": {
              "description": "SSL management type for this domain. If 'AUTOMATIC', a managed certificate is automatically provisioned.\nIf 'MANUAL', 'certificateId' must be manually specified in order to configure SSL for this domain. Possible values: [\"AUTOMATIC\", \"MANUAL\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.",
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

func GoogleAppEngineDomainMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAppEngineDomainMapping), &result)
	return &result
}
