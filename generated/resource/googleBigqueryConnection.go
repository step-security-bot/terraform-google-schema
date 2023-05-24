package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigqueryConnection = `{
  "block": {
    "attributes": {
      "connection_id": {
        "computed": true,
        "description": "Optional connection id that should be assigned to the created connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "A descriptive description for the connection",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "friendly_name": {
        "description": "A descriptive name for the connection",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "has_credential": {
        "computed": true,
        "description": "True if the connection has credential assigned.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The geographic location where the connection should reside.\nCloud SQL instance must be in the same location as the connection\nwith following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.\nExamples: US, EU, asia-northeast1, us-central1, europe-west1.\nSpanner Connections same as spanner region\nAWS allowed regions are aws-us-east-1\nAzure allowed regions are azure-eastus2",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the connection in the form of:\n\"projects/{project_id}/locations/{location_id}/connections/{connectionId}\"",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "aws": {
        "block": {
          "block_types": {
            "access_role": {
              "block": {
                "attributes": {
                  "iam_role_id": {
                    "description": "The user’s AWS IAM Role that trusts the Google-owned AWS IAM user Connection.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "identity": {
                    "computed": true,
                    "description": "A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's AWS IAM Role.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Authentication using Google owned service account to assume into customer's AWS IAM Role.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Connection properties specific to Amazon Web Services.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "azure": {
        "block": {
          "attributes": {
            "application": {
              "computed": true,
              "description": "The name of the Azure Active Directory Application.",
              "description_kind": "plain",
              "type": "string"
            },
            "client_id": {
              "computed": true,
              "description": "The client id of the Azure Active Directory Application.",
              "description_kind": "plain",
              "type": "string"
            },
            "customer_tenant_id": {
              "description": "The id of customer's directory that host the data.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "federated_application_client_id": {
              "description": "The Azure Application (client) ID where the federated credentials will be hosted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity": {
              "computed": true,
              "description": "A unique Google-owned and Google-generated identity for the Connection. This identity will be used to access the user's Azure Active Directory Application.",
              "description_kind": "plain",
              "type": "string"
            },
            "object_id": {
              "computed": true,
              "description": "The object id of the Azure Active Directory Application.",
              "description_kind": "plain",
              "type": "string"
            },
            "redirect_uri": {
              "computed": true,
              "description": "The URL user will be redirected to after granting consent during connection setup.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Container for connection properties specific to Azure.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloud_resource": {
        "block": {
          "attributes": {
            "service_account_id": {
              "computed": true,
              "description": "The account ID of the service created for the purpose of this connection.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Container for connection properties for delegation of access to GCP resources.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloud_spanner": {
        "block": {
          "attributes": {
            "database": {
              "description": "Cloud Spanner database in the form 'project/instance/database'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_parallelism": {
              "description": "If parallelism should be used when reading from Cloud Spanner",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_serverless_analytics": {
              "description": "If the serverless analytics service should be used to read data from Cloud Spanner. useParallelism must be set when using serverless analytics",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Connection properties specific to Cloud Spanner",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloud_sql": {
        "block": {
          "attributes": {
            "database": {
              "description": "Database name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "instance_id": {
              "description": "Cloud SQL instance ID in the form project:location:instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_account_id": {
              "computed": true,
              "description": "When the connection is used in the context of an operation in BigQuery, this service account will serve as the identity being used for connecting to the CloudSQL instance specified in this connection.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "description": "Type of the Cloud SQL database. Possible values: [\"DATABASE_TYPE_UNSPECIFIED\", \"POSTGRES\", \"MYSQL\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "credential": {
              "block": {
                "attributes": {
                  "password": {
                    "description": "Password for database.",
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
                    "description": "Username for database.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Cloud SQL properties.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Connection properties specific to the Cloud SQL.",
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

func GoogleBigqueryConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigqueryConnection), &result)
	return &result
}
