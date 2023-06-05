package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDnsRecordSet = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_zone": {
        "description": "The name of the zone in which this record set will reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The DNS name this record set will apply to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rrdatas": {
        "description": "The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding \\\" if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add \\\"\\\" inside the Terraform configuration string (e.g. \"first255characters\\\"\\\"morecharacters\").",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ttl": {
        "description": "The time-to-live of this record set (seconds).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "type": {
        "description": "The DNS record set type.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "routing_policy": {
        "block": {
          "attributes": {
            "enable_geo_fencing": {
              "description": "Specifies whether to enable fencing for geo queries.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "geo": {
              "block": {
                "attributes": {
                  "location": {
                    "description": "The location name defined in Google Cloud.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rrdatas": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "health_checked_targets": {
                    "block": {
                      "block_types": {
                        "internal_load_balancers": {
                          "block": {
                            "attributes": {
                              "ip_address": {
                                "description": "The frontend IP address of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ip_protocol": {
                                "description": "The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [\"tcp\", \"udp\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "load_balancer_type": {
                                "description": "The type of load balancer. This value is case-sensitive. Possible values: [\"regionalL4ilb\", \"regionalL7ilb]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "network_url": {
                                "description": "The fully qualified url of the network in which the load balancer belongs. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "The configured port of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project": {
                                "description": "The ID of the project in which the load balancer belongs.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "region": {
                                "description": "The region of the load balancer. Only needed for regional load balancers.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The list of internal load balancers to health check.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "For A and AAAA types only. The list of targets to be health checked. These can be specified along with ` + "`" + `rrdatas` + "`" + ` within this item.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration for Geo location based routing policy.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "primary_backup": {
              "block": {
                "attributes": {
                  "enable_geo_fencing_for_backups": {
                    "description": "Specifies whether to enable fencing for backup geo queries.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "trickle_ratio": {
                    "description": "Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "backup_geo": {
                    "block": {
                      "attributes": {
                        "location": {
                          "description": "The location name defined in Google Cloud.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "rrdatas": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "block_types": {
                        "health_checked_targets": {
                          "block": {
                            "block_types": {
                              "internal_load_balancers": {
                                "block": {
                                  "attributes": {
                                    "ip_address": {
                                      "description": "The frontend IP address of the load balancer.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "ip_protocol": {
                                      "description": "The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [\"tcp\", \"udp\"]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "load_balancer_type": {
                                      "description": "The type of load balancer. This value is case-sensitive. Possible values: [\"regionalL4ilb\", \"regionalL7ilb]",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "network_url": {
                                      "description": "The fully qualified url of the network in which the load balancer belongs. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + `.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "port": {
                                      "description": "The configured port of the load balancer.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "project": {
                                      "description": "The ID of the project in which the load balancer belongs.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "region": {
                                      "description": "The region of the load balancer. Only needed for regional load balancers.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "description": "The list of internal load balancers to health check.",
                                  "description_kind": "plain"
                                },
                                "min_items": 1,
                                "nesting_mode": "list"
                              }
                            },
                            "description": "For A and AAAA types only. The list of targets to be health checked. These can be specified along with ` + "`" + `rrdatas` + "`" + ` within this item.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "primary": {
                    "block": {
                      "block_types": {
                        "internal_load_balancers": {
                          "block": {
                            "attributes": {
                              "ip_address": {
                                "description": "The frontend IP address of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ip_protocol": {
                                "description": "The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [\"tcp\", \"udp\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "load_balancer_type": {
                                "description": "The type of load balancer. This value is case-sensitive. Possible values: [\"regionalL4ilb\", \"regionalL7ilb]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "network_url": {
                                "description": "The fully qualified url of the network in which the load balancer belongs. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "The configured port of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project": {
                                "description": "The ID of the project in which the load balancer belongs.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "region": {
                                "description": "The region of the load balancer. Only needed for regional load balancers.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The list of internal load balancers to health check.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of global primary targets to be health checked.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration for a primary-backup policy with global to regional failover. Queries are responded to with the global primary targets, but if none of the primary targets are healthy, then we fallback to a regional failover policy.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "wrr": {
              "block": {
                "attributes": {
                  "rrdatas": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "weight": {
                    "description": "The ratio of traffic routed to the target.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "health_checked_targets": {
                    "block": {
                      "block_types": {
                        "internal_load_balancers": {
                          "block": {
                            "attributes": {
                              "ip_address": {
                                "description": "The frontend IP address of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ip_protocol": {
                                "description": "The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [\"tcp\", \"udp\"]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "load_balancer_type": {
                                "description": "The type of load balancer. This value is case-sensitive. Possible values: [\"regionalL4ilb\", \"regionalL7ilb]",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "network_url": {
                                "description": "The fully qualified url of the network in which the load balancer belongs. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + `.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "port": {
                                "description": "The configured port of the load balancer.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "project": {
                                "description": "The ID of the project in which the load balancer belongs.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "region": {
                                "description": "The region of the load balancer. Only needed for regional load balancers.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "description": "The list of internal load balancers to health check.",
                            "description_kind": "plain"
                          },
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of ` + "`" + `rrdatas` + "`" + ` or ` + "`" + `health_checked_targets` + "`" + ` can be set.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration for Weighted Round Robin based routing policy.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The configuration for steering traffic based on query. You can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleDnsRecordSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDnsRecordSet), &result)
	return &result
}
