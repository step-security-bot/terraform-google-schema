package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataprocMetastoreService = `{
  "block": {
    "attributes": {
      "artifact_gcs_uri": {
        "computed": true,
        "description": "A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_type": {
        "description": "The database type that the Metastore service stores its data. Default value: \"MYSQL\" Possible values: [\"MYSQL\", \"SPANNER\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_uri": {
        "computed": true,
        "description": "The URI of the endpoint used to access the metastore service.",
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
        "description": "User-defined labels for the metastore service.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the metastore service should reside.\nThe default value is 'global'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The relative resource name of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:\n\n\"projects/{projectNumber}/global/networks/{network_id}\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The TCP port at which the metastore service is reached. Default: 9083.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_channel": {
        "description": "The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: \"STABLE\" Possible values: [\"CANARY\", \"STABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_id": {
        "description": "The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),\nand hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between\n3 and 63 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_message": {
        "computed": true,
        "description": "Additional information about the current state of the metastore service, if available.",
        "description_kind": "plain",
        "type": "string"
      },
      "tier": {
        "computed": true,
        "description": "The tier of the service. Possible values: [\"DEVELOPER\", \"ENTERPRISE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "The globally unique resource identifier of the metastore service.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "The fully qualified customer provided Cloud KMS key name to use for customer data encryption.\nUse the following format: 'projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Information used to configure the Dataproc Metastore service to encrypt\ncustomer data at rest.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hive_metastore_config": {
        "block": {
          "attributes": {
            "config_overrides": {
              "computed": true,
              "description": "A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).\nThe mappings override system defaults (some keys cannot be overridden)",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "version": {
              "description": "The Hive metastore schema version.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "kerberos_config": {
              "block": {
                "attributes": {
                  "krb5_config_gcs_uri": {
                    "description": "A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "principal": {
                    "description": "A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form \"primary/instance@REALM\", but there is no exact format.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "keytab": {
                    "block": {
                      "attributes": {
                        "cloud_secret": {
                          "description": "The relative resource name of a Secret Manager secret version, in the following form:\n\n\"projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}\".",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Information used to configure the Hive metastore service as a service principal in a Kerberos realm.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration information specific to running Hive metastore software as the metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_window": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description": "The day of week, when the window starts. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hour_of_day": {
              "description": "The hour of day (0-23) when the window starts.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "The one hour maintenance window of the metastore service.\nThis specifies when the service can be restarted for maintenance purposes in UTC time.\nMaintenance window is not needed for services with the 'SPANNER' database type.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "block_types": {
            "consumers": {
              "block": {
                "attributes": {
                  "endpoint_uri": {
                    "computed": true,
                    "description": "The URI of the endpoint used to access the metastore service.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnetwork": {
                    "description": "The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.\nIt is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.\nThere must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:\n'projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The consumer-side network configuration for the Dataproc Metastore instance.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The configuration specifying the network settings for the Dataproc Metastore service.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "telemetry_config": {
        "block": {
          "attributes": {
            "log_format": {
              "description": "The output format of the Dataproc Metastore service's logs. Default value: \"JSON\" Possible values: [\"LEGACY\", \"JSON\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.",
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

func GoogleDataprocMetastoreServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataprocMetastoreService), &result)
	return &result
}
