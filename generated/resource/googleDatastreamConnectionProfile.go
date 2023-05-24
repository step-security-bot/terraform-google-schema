package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatastreamConnectionProfile = `{
  "block": {
    "attributes": {
      "connection_profile_id": {
        "description": "The connection profile identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name.",
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
      "labels": {
        "description": "Labels.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The name of the location this connection profile is located in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource's name.",
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
      "bigquery_profile": {
        "block": {
          "description": "BigQuery warehouse profile.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "forward_ssh_connectivity": {
        "block": {
          "attributes": {
            "hostname": {
              "description": "Hostname for the SSH tunnel.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "SSH password.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description": "Port for the SSH tunnel.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "private_key": {
              "description": "SSH private key.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "username": {
              "description": "Username for the SSH tunnel.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Forward SSH tunnel connectivity.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gcs_profile": {
        "block": {
          "attributes": {
            "bucket": {
              "description": "The Cloud Storage bucket name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "root_path": {
              "description": "The root path inside the Cloud Storage bucket.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Cloud Storage bucket profile.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mysql_profile": {
        "block": {
          "attributes": {
            "hostname": {
              "description": "Hostname for the MySQL connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "Password for the MySQL connection.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description": "Port for the MySQL connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "username": {
              "description": "Username for the MySQL connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ssl_config": {
              "block": {
                "attributes": {
                  "ca_certificate": {
                    "description": "PEM-encoded certificate of the CA that signed the source database\nserver's certificate.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "ca_certificate_set": {
                    "computed": true,
                    "description": "Indicates whether the clientKey field is set.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "client_certificate": {
                    "description": "PEM-encoded certificate that will be used by the replica to\nauthenticate against the source database server. If this field\nis used then the 'clientKey' and the 'caCertificate' fields are\nmandatory.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_certificate_set": {
                    "computed": true,
                    "description": "Indicates whether the clientCertificate field is set.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "client_key": {
                    "description": "PEM-encoded private key associated with the Client Certificate.\nIf this field is used then the 'client_certificate' and the\n'ca_certificate' fields are mandatory.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_key_set": {
                    "computed": true,
                    "description": "Indicates whether the clientKey field is set.",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "description": "SSL configuration for the MySQL connection.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "MySQL database profile.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "oracle_profile": {
        "block": {
          "attributes": {
            "connection_attributes": {
              "description": "Connection string attributes",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "database_service": {
              "description": "Database for the Oracle connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hostname": {
              "description": "Hostname for the Oracle connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "Password for the Oracle connection.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description": "Port for the Oracle connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "username": {
              "description": "Username for the Oracle connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Oracle database profile.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "postgresql_profile": {
        "block": {
          "attributes": {
            "database": {
              "description": "Database for the PostgreSQL connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "hostname": {
              "description": "Hostname for the PostgreSQL connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "Password for the PostgreSQL connection.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "port": {
              "description": "Port for the PostgreSQL connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "username": {
              "description": "Username for the PostgreSQL connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "PostgreSQL database profile.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "private_connectivity": {
        "block": {
          "attributes": {
            "private_connection": {
              "description": "A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Private connectivity.",
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

func GoogleDatastreamConnectionProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatastreamConnectionProfile), &result)
	return &result
}
