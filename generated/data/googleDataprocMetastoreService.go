package data

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
        "computed": true,
        "description": "The database type that the Metastore service stores its data. Default value: \"MYSQL\" Possible values: [\"MYSQL\", \"SPANNER\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_config": {
        "computed": true,
        "description": "Information used to configure the Dataproc Metastore service to encrypt\ncustomer data at rest.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key": "string"
            }
          ]
        ]
      },
      "endpoint_uri": {
        "computed": true,
        "description": "The URI of the endpoint used to access the metastore service.",
        "description_kind": "plain",
        "type": "string"
      },
      "hive_metastore_config": {
        "computed": true,
        "description": "Configuration information specific to running Hive metastore software as the metastore service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "config_overrides": [
                "map",
                "string"
              ],
              "kerberos_config": [
                "list",
                [
                  "object",
                  {
                    "keytab": [
                      "list",
                      [
                        "object",
                        {
                          "cloud_secret": "string"
                        }
                      ]
                    ],
                    "krb5_config_gcs_uri": "string",
                    "principal": "string"
                  }
                ]
              ],
              "version": "string"
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
      "labels": {
        "computed": true,
        "description": "User-defined labels for the metastore service.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the metastore service should reside.\nThe default value is 'global'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_window": {
        "computed": true,
        "description": "The one hour maintenance window of the metastore service.\nThis specifies when the service can be restarted for maintenance purposes in UTC time.\nMaintenance window is not needed for services with the 'SPANNER' database type.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "day_of_week": "string",
              "hour_of_day": "number"
            }
          ]
        ]
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
        "type": "string"
      },
      "network_config": {
        "computed": true,
        "description": "The configuration specifying the network settings for the Dataproc Metastore service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "consumers": [
                "list",
                [
                  "object",
                  {
                    "endpoint_uri": "string",
                    "subnetwork": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "port": {
        "computed": true,
        "description": "The TCP port at which the metastore service is reached. Default: 9083.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "release_channel": {
        "computed": true,
        "description": "The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: \"STABLE\" Possible values: [\"CANARY\", \"STABLE\"]",
        "description_kind": "plain",
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
      "telemetry_config": {
        "computed": true,
        "description": "The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "log_format": "string"
            }
          ]
        ]
      },
      "tier": {
        "computed": true,
        "description": "The tier of the service. Possible values: [\"DEVELOPER\", \"ENTERPRISE\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "The globally unique resource identifier of the metastore service.",
        "description_kind": "plain",
        "type": "string"
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
