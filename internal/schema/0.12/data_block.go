package schema

import (
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/terraform-schema/internal/schema/refscope"
	"github.com/zclconf/go-cty/cty"
)

func datasourceBlockSchema(v *version.Version) *schema.BlockSchema {
	bs := &schema.BlockSchema{
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
						schema.ScopeTraversalExpr{ScopeId: refscope.ProviderBlock},
					},
					IsOptional:  true,
					Description: lang.Markdown("Reference to a `provider` configuration block, e.g. `mycloud.west` or `mycloud`"),
					IsDepKey:    true,
				},
				"count": {
					Expr: schema.ExprSchema{
						schema.LiteralValueExpr{Type: cty.Number},
						schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock, OfType: cty.Number},
						schema.ScopeTraversalExpr{ScopeId: refscope.LocalAttr, OfType: cty.Number},
						schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock, OfType: cty.Number},
						schema.ScopeTraversalExpr{ScopeId: refscope.VariableBlock, OfType: cty.Number},
					},
					IsOptional:  true,
					Description: lang.Markdown("Number of instances of this data source, e.g. `3`"),
				},
				"depends_on": {
					Expr: schema.ExprSchema{
						schema.TupleExpr{
							Exprs: []schema.Expr{
								schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock},
								schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock},
								schema.ScopeTraversalExpr{ScopeId: refscope.ModuleBlock},
							},
						},
					},
					IsOptional:  true,
					Description: lang.Markdown("Set of references to hidden dependencies, e.g. other resources or data sources"),
				},
			},
		},
	}

	if v.GreaterThanOrEqual(v0_12_6) {
		bs.Body.Attributes["for_each"] = &schema.AttributeSchema{
			Expr: schema.ExprSchema{
				schema.LiteralValueExpr{Type: cty.Set(cty.DynamicPseudoType)},
				schema.LiteralValueExpr{Type: cty.Map(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock, OfType: cty.Set(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.DatasourceBlock, OfType: cty.Map(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock, OfType: cty.Set(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.ResourceBlock, OfType: cty.Map(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.VariableBlock, OfType: cty.Set(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.VariableBlock, OfType: cty.Map(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.LocalAttr, OfType: cty.Set(cty.DynamicPseudoType)},
				schema.ScopeTraversalExpr{ScopeId: refscope.LocalAttr, OfType: cty.Map(cty.DynamicPseudoType)},
			},
			IsOptional:  true,
			Description: lang.Markdown("A set or a map where each item represents an instance of this data source"),
		}
	}

	return bs
}
