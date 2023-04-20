package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-template-schema/v2/generated/data"
	"github.com/lonegunmanb/terraform-template-schema/v2/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["template_cloudinit_config"] = resource.TemplateCloudinitConfigSchema()  
	resources["template_dir"] = resource.TemplateDirSchema()  
	resources["template_file"] = resource.TemplateFileSchema()  
	dataSources["template_cloudinit_config"] = data.TemplateCloudinitConfigSchema()  
	dataSources["template_file"] = data.TemplateFileSchema()  
	Resources = resources
	DataSources = dataSources
}