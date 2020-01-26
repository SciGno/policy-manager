package graphql

import "github.com/graphql-go/graphql"

// RequestInputType used for mutation and queries
var RequestInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "RequestInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"principal": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(String),
			},
			"action": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.NewList(String)),
			},
			"resource": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(String),
			},
			"condition": &graphql.InputObjectFieldConfig{
				Type: graphql.NewList(ConditionInputType),
			},
		},
	},
)

// ConditionInputType used for mutation and queries
var ConditionInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ConditionInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"value": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)
