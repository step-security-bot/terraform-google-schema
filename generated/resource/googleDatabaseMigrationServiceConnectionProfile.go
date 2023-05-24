package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDatabaseMigrationServiceConnectionProfile = `{
  "block": {
    "attributes": {
      "connection_profile_id": {
        "description": "The ID of the connection profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. The timestamp when the resource was created. A timestamp in RFC3339 UTC 'Zulu' format, accurate to nanoseconds. Example: '2014-10-02T15:01:23.045123456Z'.",
        "description_kind": "plain",
        "type": "string"
      },
      "dbprovider": {
        "computed": true,
        "description": "The database provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The connection profile display name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "error": {
        "computed": true,
        "description": "Output only. The error details in case of state FAILED.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "number",
              "details": [
                "list",
                [
                  "map",
                  "string"
                ]
              ],
              "message": "string"
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
        "description": "The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the connection profile should reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current connection profile state.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "alloydb": {
        "block": {
          "attributes": {
            "cluster_id": {
              "description": "Required. The AlloyDB cluster ID that this connection profile is associated with.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "settings": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "Labels for the AlloyDB cluster created by DMS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "vpc_network": {
                    "description": "Required. The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.\nIt is specified in the form: 'projects/{project_number}/global/networks/{network_id}'. This is required to create a cluster.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "initial_user": {
                    "block": {
                      "attributes": {
                        "password": {
                          "description": "The initial password for the user.",
                          "description_kind": "plain",
                          "required": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "password_set": {
                          "computed": true,
                          "description": "Output only. Indicates if the initialUser.password field has been set.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "user": {
                          "description": "The database username.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Required. Input only. Initial user to setup during cluster creation.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "primary_instance_settings": {
                    "block": {
                      "attributes": {
                        "database_flags": {
                          "description": "Database flags to pass to AlloyDB when DMS is creating the AlloyDB cluster and instances. See the AlloyDB documentation for how these can be used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "id": {
                          "description": "The database username.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "labels": {
                          "description": "Labels for the AlloyDB primary instance created by DMS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "private_ip": {
                          "computed": true,
                          "description": "Output only. The private IP address for the Instance. This is the connection endpoint for an end-user application.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "machine_config": {
                          "block": {
                            "attributes": {
                              "cpu_count": {
                                "description": "The number of CPU's in the VM instance.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              }
                            },
                            "description": "Configuration for the machines that host the underlying database engine.",
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        }
                      },
                      "description": "Settings for the cluster's primary instance",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Immutable. Metadata used to create the destination AlloyDB cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies required connection parameters, and the parameters required to create an AlloyDB destination cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cloudsql": {
        "block": {
          "attributes": {
            "cloud_sql_id": {
              "computed": true,
              "description": "Output only. The Cloud SQL instance ID that this connection profile is associated with.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_ip": {
              "computed": true,
              "description": "Output only. The Cloud SQL database instance's private IP.",
              "description_kind": "plain",
              "type": "string"
            },
            "public_ip": {
              "computed": true,
              "description": "Output only. The Cloud SQL database instance's public IP.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "settings": {
              "block": {
                "attributes": {
                  "activation_policy": {
                    "description": "The activation policy specifies when the instance is activated; it is applicable only when the instance state is 'RUNNABLE'. Possible values: [\"ALWAYS\", \"NEVER\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "auto_storage_increase": {
                    "description": "If you enable this setting, Cloud SQL checks your available storage every 30 seconds. If the available storage falls below a threshold size, Cloud SQL automatically adds additional storage capacity.\nIf the available storage repeatedly falls below the threshold size, Cloud SQL continues to add storage until it reaches the maximum of 30 TB.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cmek_key_name": {
                    "description": "The KMS key name used for the csql instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "collation": {
                    "description": "The Cloud SQL default instance level collation.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_disk_size_gb": {
                    "description": "The storage capacity available to the database, in GB. The minimum (and default) size is 10GB.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_disk_type": {
                    "description": "The type of storage. Possible values: [\"PD_SSD\", \"PD_HDD\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "database_flags": {
                    "description": "The database flags passed to the Cloud SQL instance at startup.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "database_version": {
                    "description": "The database engine type and version.\nCurrently supported values located at https://cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.connectionProfiles#sqldatabaseversion",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "root_password": {
                    "description": "Input only. Initial root password.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "root_password_set": {
                    "computed": true,
                    "description": "Output only. Indicates If this connection profile root password is stored.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "source_id": {
                    "description": "The Database Migration Service source connection profile ID, in the format: projects/my_project_name/locations/us-central1/connectionProfiles/connection_profile_ID",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "storage_auto_resize_limit": {
                    "description": "The maximum size to which storage capacity can be automatically increased. The default value is 0, which specifies that there is no limit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tier": {
                    "description": "The tier (or machine type) for this instance, for example: db-n1-standard-1 (MySQL instances) or db-custom-1-3840 (PostgreSQL instances).\nFor more information, see https://cloud.google.com/sql/docs/mysql/instance-settings",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "user_labels": {
                    "description": "The resource labels for a Cloud SQL instance to use to annotate any related underlying resources such as Compute Engine VMs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "zone": {
                    "description": "The Google Cloud Platform zone where your Cloud SQL datdabse instance is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ip_config": {
                    "block": {
                      "attributes": {
                        "enable_ipv4": {
                          "description": "Whether the instance should be assigned an IPv4 address or not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "private_network": {
                          "description": "The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP. For example, projects/myProject/global/networks/default.\nThis setting can be updated, but it cannot be removed after it is set.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "require_ssl": {
                          "description": "Whether SSL connections over IP should be enforced or not.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "authorized_networks": {
                          "block": {
                            "attributes": {
                              "expire_time": {
                                "description": "The time when this access control entry expires in RFC 3339 format.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "description": "A label to identify this entry.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "ttl": {
                                "description": "Input only. The time-to-leave of this access control entry.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The allowlisted value for the access control list.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description": "The list of external networks that are allowed to connect to the instance using the IP.",
                            "description_kind": "plain"
                          },
                          "nesting_mode": "list"
                        }
                      },
                      "description": "The settings for IP Management. This allows to enable or disable the instance IP and manage which external networks can connect to the instance. The IPv4 address cannot be disabled.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Immutable. Metadata used to create the destination Cloud SQL database.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mysql": {
        "block": {
          "attributes": {
            "cloud_sql_id": {
              "description": "If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host": {
              "description": "Required. The IP or hostname of the source MySQL database.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description": "Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.\nThis field is not returned on request, and the value is encrypted when stored in Database Migration Service.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "password_set": {
              "computed": true,
              "description": "Output only. Indicates If this connection profile password is stored.",
              "description_kind": "plain",
              "type": "bool"
            },
            "port": {
              "description": "Required. The network port of the source MySQL database.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "username": {
              "description": "Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ssl": {
              "block": {
                "attributes": {
                  "ca_certificate": {
                    "description": "Required. Input only. The x509 PEM-encoded certificate of the CA that signed the source database server's certificate.\nThe replica will use this certificate to verify it's connecting to the right host.",
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_certificate": {
                    "description": "Input only. The x509 PEM-encoded certificate that will be used by the replica to authenticate against the source database server.\nIf this field is used then the 'clientKey' field is mandatory",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_key": {
                    "description": "Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key associated with the Client Certificate.\nIf this field is used then the 'clientCertificate' field is mandatory.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The current connection profile state.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "SSL configuration for the destination to connect to the source database.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies connection parameters required specifically for MySQL databases.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "postgresql": {
        "block": {
          "attributes": {
            "cloud_sql_id": {
              "description": "If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host": {
              "description": "Required. The IP or hostname of the source MySQL database.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_architecture": {
              "computed": true,
              "description": "Output only. If the source is a Cloud SQL database, this field indicates the network architecture it's associated with.",
              "description_kind": "plain",
              "type": "string"
            },
            "password": {
              "description": "Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.\nThis field is not returned on request, and the value is encrypted when stored in Database Migration Service.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "password_set": {
              "computed": true,
              "description": "Output only. Indicates If this connection profile password is stored.",
              "description_kind": "plain",
              "type": "bool"
            },
            "port": {
              "description": "Required. The network port of the source MySQL database.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "username": {
              "description": "Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ssl": {
              "block": {
                "attributes": {
                  "ca_certificate": {
                    "description": "Required. Input only. The x509 PEM-encoded certificate of the CA that signed the source database server's certificate.\nThe replica will use this certificate to verify it's connecting to the right host.",
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_certificate": {
                    "description": "Input only. The x509 PEM-encoded certificate that will be used by the replica to authenticate against the source database server.\nIf this field is used then the 'clientKey' field is mandatory",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_key": {
                    "description": "Input only. The unencrypted PKCS#1 or PKCS#8 PEM-encoded private key associated with the Client Certificate.\nIf this field is used then the 'clientCertificate' field is mandatory.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The current connection profile state.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "SSL configuration for the destination to connect to the source database.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Specifies connection parameters required specifically for PostgreSQL databases.",
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

func GoogleDatabaseMigrationServiceConnectionProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDatabaseMigrationServiceConnectionProfile), &result)
	return &result
}
