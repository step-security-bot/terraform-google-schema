package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMonitoringUptimeCheckIps = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uptime_check_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": "string",
              "location": "string",
              "region": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMonitoringUptimeCheckIpsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMonitoringUptimeCheckIps), &result)
	return &result
}
