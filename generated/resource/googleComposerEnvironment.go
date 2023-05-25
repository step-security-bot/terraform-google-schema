package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComposerEnvironment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: [a-z]([-a-z0-9]*[a-z0-9])?. Label values must be between 0 and 63 characters long and must conform to the regular expression ([a-z]([-a-z0-9]*[a-z0-9])?)?. No more than 64 labels can be associated with a given environment. Both keys and values must be \u003c= 128 bytes in size.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description": "Name of the environment.",
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
      "region": {
        "description": "The location or Compute Engine region for the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "attributes": {
            "airflow_uri": {
              "computed": true,
              "description": "The URI of the Apache Airflow Web UI hosted within this environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "dag_gcs_prefix": {
              "computed": true,
              "description": "The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.",
              "description_kind": "plain",
              "type": "string"
            },
            "gke_cluster": {
              "computed": true,
              "description": "The Kubernetes Engine cluster used to run this environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "node_count": {
              "computed": true,
              "description": "The number of nodes in the Kubernetes Engine cluster that will be used to run this environment. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "node_config": {
              "block": {
                "attributes": {
                  "disk_size_gb": {
                    "computed": true,
                    "description": "The disk size in GB used for node VMs. Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ip_allocation_policy": {
                    "computed": true,
                    "description": "Configuration for controlling how IPs are allocated in the GKE cluster. Cannot be updated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "cluster_ipv4_cidr_block": "string",
                          "cluster_secondary_range_name": "string",
                          "services_ipv4_cidr_block": "string",
                          "services_secondary_range_name": "string",
                          "use_ip_aliases": "bool"
                        }
                      ]
                    ]
                  },
                  "machine_type": {
                    "computed": true,
                    "description": "The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: \"projects/{project}/zones/{zone}/machineTypes/{machineType}\". Must belong to the enclosing environment's project and region/zone. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network": {
                    "computed": true,
                    "description": "The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: \"projects/{project}/zones/{zone}/machineTypes/{machineType}\". Must belong to the enclosing environment's project and region/zone. The network must belong to the environment's project. If unspecified, the \"default\" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "computed": true,
                    "description": "The set of Google API scopes to be made available on all node VMs. Cannot be updated. If empty, defaults to [\"https://www.googleapis.com/auth/cloud-platform\"]. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "service_account": {
                    "computed": true,
                    "description": "The Google Cloud Platform Service Account to be used by the node VMs. If a service account is not specified, the \"default\" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have roles/composer.worker for any GCP resources created under the Cloud Composer Environment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnetwork": {
                    "description": "The Compute Engine subnetwork to be used for machine communications, , specified as a self-link, relative resource name (e.g. \"projects/{project}/regions/{region}/subnetworks/{subnetwork}\"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tags": {
                    "description": "The list of instance tags applied to all node VMs. Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "zone": {
                    "computed": true,
                    "description": "The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. \"projects/{project}/zones/{zone}\"). Must belong to the enclosing environment's project and region. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The configuration used for the Kubernetes Engine cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "private_environment_config": {
              "block": {
                "attributes": {
                  "cloud_sql_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The CIDR block from which IP range in tenant project will be reserved for Cloud SQL. Needs to be disjoint from web_server_ipv4_cidr_block.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enable_private_endpoint": {
                    "description": "If true, access to the public endpoint of the GKE cluster is denied. If this field is set to true, ip_allocation_policy.use_ip_aliases must be set to true for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "master_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The IP range in CIDR notation to use for the hosted master network. This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network. If left blank, the default value of '172.16.0.0/28' is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "web_server_ipv4_cidr_block": {
                    "computed": true,
                    "description": "The CIDR block from which IP range for web server will be reserved. Needs to be disjoint from master_ipv4_cidr_block and cloud_sql_ipv4_cidr_block. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The configuration used for the Private IP Cloud Composer environment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "software_config": {
              "block": {
                "attributes": {
                  "airflow_config_overrides": {
                    "description": "Apache Airflow configuration properties to override. Property keys contain the section and property names, separated by a hyphen, for example \"core-dags_are_paused_at_creation\". Section names must not contain hyphens (\"-\"), opening square brackets (\"[\"), or closing square brackets (\"]\"). The property name must not be empty and cannot contain \"=\" or \";\". Section and property names cannot contain characters: \".\" Apache Airflow configuration property names must be written in snake_case. Property values can contain any character, and can be written in any lower/upper case format. Certain Apache Airflow configuration property values are blacklisted, and cannot be overridden.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "env_variables": {
                    "description": "Additional environment variables to provide to the Apache Airflow schedulerf, worker, and webserver processes. Environment variable names must match the regular expression [a-zA-Z_][a-zA-Z0-9_]*. They cannot specify Apache Airflow software configuration overrides (they cannot match the regular expression AIRFLOW__[A-Z0-9_]+__[A-Z0-9_]+), and they cannot match any of the following reserved names: AIRFLOW_HOME C_FORCE_ROOT CONTAINER_NAME DAGS_FOLDER GCP_PROJECT GCS_BUCKET GKE_CLUSTER_NAME SQL_DATABASE SQL_INSTANCE SQL_PASSWORD SQL_PROJECT SQL_REGION SQL_USER.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "image_version": {
                    "computed": true,
                    "description": "The version of the software running in the environment. This encapsulates both the version of Cloud Composer functionality and the version of Apache Airflow. It must match the regular expression composer-[0-9]+\\.[0-9]+(\\.[0-9]+)?-airflow-[0-9]+\\.[0-9]+(\\.[0-9]+.*)?. The Cloud Composer portion of the version is a semantic version. The portion of the image version following 'airflow-' is an official Apache Airflow repository release name. See documentation for allowed release names.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pypi_packages": {
                    "description": "Custom Python Package Index (PyPI) packages to be installed in the environment. Keys refer to the lowercase package name (e.g. \"numpy\"). Values are the lowercase extras and version specifier (e.g. \"==1.12.0\", \"[devel,gcp_api]\", \"[devel]\u003e=1.8.2, \u003c1.9.2\"). To specify a package without pinning it to a version specifier, use the empty string as the value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "python_version": {
                    "computed": true,
                    "description": "The major version of Python used to run the Apache Airflow scheduler, worker, and webserver processes. Can be set to '2' or '3'. If not specified, the default is '2'. Cannot be updated. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*. Environments in newer versions always use Python major version 3.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "scheduler_count": {
                    "computed": true,
                    "description": "The number of schedulers for Airflow. This field is supported for Cloud Composer environments in versions composer-1.*.*-airflow-2.*.*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "The configuration settings for software inside the environment.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration parameters for this environment.",
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

func GoogleComposerEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComposerEnvironment), &result)
	return &result
}
