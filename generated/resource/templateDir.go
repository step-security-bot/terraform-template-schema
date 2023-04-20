package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const templateDir = `{
  "block": {
    "attributes": {
      "destination_dir": {
        "description": "Path to the directory where the templated files will be written",
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
      "source_dir": {
        "description": "Path to the directory where the files to template reside",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vars": {
        "description": "Variables to substitute",
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

func TemplateDirSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(templateDir), &result)
	return &result
}
