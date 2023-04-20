package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const templateFile = `{
  "block": {
    "attributes": {
      "filename": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rendered": {
        "computed": true,
        "description": "rendered template",
        "description_kind": "plain",
        "type": "string"
      },
      "template": {
        "description": "Contents of the template",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vars": {
        "description": "variables to substitute",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func TemplateFileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(templateFile), &result)
	return &result
}
