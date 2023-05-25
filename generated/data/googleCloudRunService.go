package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudRunService = `{
  "block": {
    "attributes": {
      "autogenerate_revision_name": {
        "computed": true,
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
        "description": "The location of the cloud run instance. eg us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "Metadata associated with this Service, including name, namespace, labels,\nand annotations.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "annotations": [
                "map",
                "string"
              ],
              "generation": "number",
              "labels": [
                "map",
                "string"
              ],
              "namespace": "string",
              "resource_version": "string",
              "self_link": "string",
              "uid": "string"
            }
          ]
        ]
      },
      "name": {
        "description": "Name must be unique within a namespace, within a Cloud Run region.\nIs required when creating resources. Name is primarily intended\nfor creation idempotence and configuration definition. Cannot be updated.\nMore info: http://kubernetes.io/docs/user-guide/identifiers#names",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the Service.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "message": "string",
                    "reason": "string",
                    "status": "string",
                    "type": "string"
                  }
                ]
              ],
              "latest_created_revision_name": "string",
              "latest_ready_revision_name": "string",
              "observed_generation": "number",
              "url": "string"
            }
          ]
        ]
      },
      "template": {
        "computed": true,
        "description": "template holds the latest specification for the Revision to\nbe stamped out. The template references the container image, and may also\ninclude labels and annotations that should be attached to the Revision.\nTo correlate a Revision, and/or to force a Revision to be created when the\nspec doesn't otherwise change, a nonce label may be provided in the\ntemplate metadata. For more details, see:\nhttps://github.com/knative/serving/blob/master/docs/client-conventions.md#associate-modifications-with-revisions\n\nCloud Run does not currently support referencing a build that is\nresponsible for materializing the container image from source.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "metadata": [
                "list",
                [
                  "object",
                  {
                    "annotations": [
                      "map",
                      "string"
                    ],
                    "generation": "number",
                    "labels": [
                      "map",
                      "string"
                    ],
                    "name": "string",
                    "namespace": "string",
                    "resource_version": "string",
                    "self_link": "string",
                    "uid": "string"
                  }
                ]
              ],
              "spec": [
                "list",
                [
                  "object",
                  {
                    "container_concurrency": "number",
                    "containers": [
                      "list",
                      [
                        "object",
                        {
                          "args": [
                            "list",
                            "string"
                          ],
                          "command": [
                            "list",
                            "string"
                          ],
                          "env": [
                            "set",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "env_from": [
                            "list",
                            [
                              "object",
                              {
                                "config_map_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "local_object_reference": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string"
                                          }
                                        ]
                                      ],
                                      "optional": "bool"
                                    }
                                  ]
                                ],
                                "prefix": "string",
                                "secret_ref": [
                                  "list",
                                  [
                                    "object",
                                    {
                                      "local_object_reference": [
                                        "list",
                                        [
                                          "object",
                                          {
                                            "name": "string"
                                          }
                                        ]
                                      ],
                                      "optional": "bool"
                                    }
                                  ]
                                ]
                              }
                            ]
                          ],
                          "image": "string",
                          "ports": [
                            "list",
                            [
                              "object",
                              {
                                "container_port": "number",
                                "name": "string",
                                "protocol": "string"
                              }
                            ]
                          ],
                          "resources": [
                            "list",
                            [
                              "object",
                              {
                                "limits": [
                                  "map",
                                  "string"
                                ],
                                "requests": [
                                  "map",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "working_dir": "string"
                        }
                      ]
                    ],
                    "service_account_name": "string",
                    "serving_state": "string",
                    "timeout_seconds": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "traffic": {
        "computed": true,
        "description": "Traffic specifies how to distribute traffic over a collection of Knative Revisions\nand Configurations",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "latest_revision": "bool",
              "percent": "number",
              "revision_name": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func GoogleCloudRunServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudRunService), &result)
	return &result
}
