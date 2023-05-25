package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDataflowJob = `{
  "block": {
    "attributes": {
      "additional_experiments": {
        "description": "List of experiments that should be used by the job. An example value is [\"enable_stackdriver_agent_metrics\"].",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "enable_streaming_engine": {
        "description": "Indicates if the job should use the streaming engine feature.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_configuration": {
        "description": "The configuration for VM IPs. Options are \"WORKER_IP_PUBLIC\" or \"WORKER_IP_PRIVATE\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "job_id": {
        "computed": true,
        "description": "The unique ID of this job.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_name": {
        "description": "The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: Google-provided Dataflow templates often provide default labels that begin with goog-dataflow-provided. Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "machine_type": {
        "description": "The machine type to use for the job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_workers": {
        "description": "The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "A unique name for the resource, required by Dataflow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The network to which VMs will be assigned. If it is not provided, \"default\" will be used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "on_delete": {
        "description": "One of \"drain\" or \"cancel\". Specifies behavior of deletion during terraform destroy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "description": "Key/Value pairs to be passed to the Dataflow job (as used in the template).",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "project": {
        "computed": true,
        "description": "The project in which the resource belongs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The region in which the created job should run.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_account_email": {
        "description": "The Service Account email used to create the job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the resource, selected from the JobState enum.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "description": "The subnetwork to which VMs will be assigned. Should be of the form \"regions/REGION/subnetworks/SUBNETWORK\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "temp_gcs_location": {
        "description": "A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "template_gcs_path": {
        "description": "The Google Cloud Storage path to the Dataflow job template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transform_name_mapping": {
        "description": "Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "The type of this job, selected from the JobType enum.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone": {
        "description": "The zone in which the created job should run. If it is not provided, the provider zone is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
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

func GoogleDataflowJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDataflowJob), &result)
	return &result
}
