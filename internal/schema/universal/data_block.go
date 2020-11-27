package schema

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
)

const datasourceBlockScope = schema.ScopeId("datasource")

var datasourceBlockSchema = &schema.BlockSchema{
	Labels: []*schema.LabelSchema{
		{
			Name:        "type",
			Description: lang.PlainText("Data Source Type"),
			IsDepKey:    true,
		},
		{
			Name:        "name",
			Description: lang.PlainText("Reference Name"),
		},
	},
	Description: lang.PlainText("A data block requests that Terraform read from a given data source and export the result " +
		"under the given local name. The name is used to refer to this resource from elsewhere in the same " +
		"Terraform module, but has no significance outside of the scope of a module."),
	Body: &schema.BodySchema{
		Attributes: map[string]*schema.AttributeSchema{
			"provider": {
				Expr: schema.ExprSchema{
					schema.ScopeTraversalExpr{ScopeId: providerBlockScope},
				},
				IsOptional:  true,
				Description: lang.Markdown("Reference to a `provider` configuration block, e.g. `mycloud.west` or `mycloud`"),
				IsDepKey:    true,
			},
		},
	},
	Reference: &schema.BlockReference{
		ScopeId: datasourceBlockScope,
		Type:    &schema.InferredRefType{},
		Address: schema.Address{
			schema.StaticStep{Value: "data"},
			schema.LabelValueStep{Index: 0},
			schema.LabelValueStep{Index: 1},
		},
	},
}
