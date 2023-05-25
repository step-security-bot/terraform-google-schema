package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeResourcePolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "An optional description of this resource. Provide this property when you create the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_placement_policy": {
        "computed": true,
        "description": "Resource policy for instances used for placement configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_domain_count": "number",
              "collocation": "string",
              "vm_count": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_schedule_policy": {
        "computed": true,
        "description": "Resource policy for scheduling instance operations.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expiration_time": "string",
              "start_time": "string",
              "time_zone": "string",
              "vm_start_schedule": [
                "list",
                [
                  "object",
                  {
                    "schedule": "string"
                  }
                ]
              ],
              "vm_stop_schedule": [
                "list",
                [
                  "object",
                  {
                    "schedule": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "description": "The name of the resource, provided by the client when initially creating\nthe resource. The resource name must be 1-63 characters long, and comply\nwith RFC1035. Specifically, the name must be 1-63 characters long and\nmatch the regular expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the\nfirst character must be a lowercase letter, and all following characters\nmust be a dash, lowercase letter, or digit, except the last character,\nwhich cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "Region where resource policy resides.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_schedule_policy": {
        "computed": true,
        "description": "Policy for creating snapshots of persistent disks.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "retention_policy": [
                "list",
                [
                  "object",
                  {
                    "max_retention_days": "number",
                    "on_source_disk_delete": "string"
                  }
                ]
              ],
              "schedule": [
                "list",
                [
                  "object",
                  {
                    "daily_schedule": [
                      "list",
                      [
                        "object",
                        {
                          "days_in_cycle": "number",
                          "start_time": "string"
                        }
                      ]
                    ],
                    "hourly_schedule": [
                      "list",
                      [
                        "object",
                        {
                          "hours_in_cycle": "number",
                          "start_time": "string"
                        }
                      ]
                    ],
                    "weekly_schedule": [
                      "list",
                      [
                        "object",
                        {
                          "day_of_weeks": [
                            "set",
                            [
                              "object",
                              {
                                "day": "string",
                                "start_time": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "snapshot_properties": [
                "list",
                [
                  "object",
                  {
                    "guest_flush": "bool",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "storage_locations": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeResourcePolicy), &result)
	return &result
}
