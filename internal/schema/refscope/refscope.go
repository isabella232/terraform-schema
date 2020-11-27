package refscope

import (
	"github.com/hashicorp/hcl-lang/schema"
)

const (
	DatasourceBlock = schema.ScopeId("datasource")
	LocalAttr       = schema.ScopeId("local")
	ModuleBlock     = schema.ScopeId("module")
	OutputBlock     = schema.ScopeId("output")
	ProviderBlock   = schema.ScopeId("provider")
	ResourceBlock   = schema.ScopeId("resource")
	VariableBlock   = schema.ScopeId("variable")
)
