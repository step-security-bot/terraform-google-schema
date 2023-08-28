package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubFeatureMembership = `{
  "block": {
    "attributes": {
      "feature": {
        "description": "The name of the feature",
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
      "location": {
        "description": "The location of the feature",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "membership": {
        "description": "The name of the membership",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The project of the feature",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "configmanagement": {
        "block": {
          "attributes": {
            "version": {
              "computed": true,
              "description": "Optional. Version of ACM to install. Defaults to the latest version.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "binauthz": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description": "Whether binauthz is enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Binauthz configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "config_sync": {
              "block": {
                "attributes": {
                  "prevent_drift": {
                    "computed": true,
                    "description": "Set to true to enable the Config Sync admission webhook to prevent drifts. If set to ` + "`" + `false` + "`" + `, disables the Config Sync admission webhook and does not prevent drifts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_format": {
                    "description": "Specifies whether the Config Sync Repo is in \"hierarchical\" or \"unstructured\" mode.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "git": {
                    "block": {
                      "attributes": {
                        "gcp_service_account_email": {
                          "description": "The GCP Service Account Email used for auth when secretType is gcpServiceAccount.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "https_proxy": {
                          "description": "URL for the HTTPS proxy to be used when communicating with the Git repo.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "policy_dir": {
                          "description": "The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_type": {
                          "description": "Type of secret configured for access to the Git repo. Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_branch": {
                          "description": "The branch of the repository to sync from. Default: master.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_repo": {
                          "description": "The URL of the Git repository to use as the source of truth.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_rev": {
                          "description": "Git revision (tag or hash) to check out. Default HEAD.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_wait_secs": {
                          "description": "Period in seconds between consecutive syncs. Default: 15.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "oci": {
                    "block": {
                      "attributes": {
                        "gcp_service_account_email": {
                          "description": "The GCP Service Account Email used for auth when secret_type is gcpserviceaccount. ",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "policy_dir": {
                          "description": "The absolute path of the directory that contains the local resources. Default: the root directory of the image.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_type": {
                          "description": "Type of secret configured for access to the OCI Image. Must be one of gcenode, gcpserviceaccount or none. The validation of this is case-sensitive.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_repo": {
                          "description": "The OCI image repository URL for the package to sync from. e.g. LOCATION-docker.pkg.dev/PROJECT_ID/REPOSITORY_NAME/PACKAGE_NAME.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sync_wait_secs": {
                          "description": "Period in seconds(int64 format) between consecutive syncs. Default: 15.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Config Sync configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "hierarchy_controller": {
              "block": {
                "attributes": {
                  "enable_hierarchical_resource_quota": {
                    "description": "Whether hierarchical resource quota is enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enable_pod_tree_labels": {
                    "description": "Whether pod tree labels are enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description": "Whether Hierarchy Controller is enabled in this cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description": "Hierarchy Controller configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "policy_controller": {
              "block": {
                "attributes": {
                  "audit_interval_seconds": {
                    "description": "Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "enabled": {
                    "description": "Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "exemptable_namespaces": {
                    "description": "The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "log_denies_enabled": {
                    "description": "Logs all denies and dry run failures.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "mutation_enabled": {
                    "description": "Enable or disable mutation in policy controller. If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "referential_rules_enabled": {
                    "description": "Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "template_library_installed": {
                    "description": "Installs the default template library along with Policy Controller.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "monitoring": {
                    "block": {
                      "attributes": {
                        "backends": {
                          "computed": true,
                          "description": " Specifies the list of backends Policy Controller will export to. Specifying an empty value ` + "`" + `[]` + "`" + ` disables metrics export.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "description": "Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: [\"cloudmonitoring\", \"prometheus\"]. Default: [\"cloudmonitoring\", \"prometheus\"]",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Policy Controller configuration for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Config Management-specific spec.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mesh": {
        "block": {
          "attributes": {
            "control_plane": {
              "deprecated": true,
              "description": "**DEPRECATED** Whether to automatically manage Service Mesh control planes. Possible values: CONTROL_PLANE_MANAGEMENT_UNSPECIFIED, AUTOMATIC, MANUAL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "management": {
              "description": "Whether to automatically manage Service Mesh. Possible values: MANAGEMENT_UNSPECIFIED, MANAGEMENT_AUTOMATIC, MANAGEMENT_MANUAL",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Manage Mesh Features",
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

func GoogleGkeHubFeatureMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFeatureMembership), &result)
	return &result
}
