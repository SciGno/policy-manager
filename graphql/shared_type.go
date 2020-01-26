package graphql

import "github.com/graphql-go/graphql"

// PageInfoType for graphql
var PageInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PageInfoType",
	Fields: graphql.Fields{
		"endCursor": &graphql.Field{
			Type: graphql.Int,
		},
		"hasNextPage": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

// ConnectionMetrics for graphql
var ConnectionMetrics = graphql.NewObject(graphql.ObjectConfig{
	Name: "ConnectionMetrics",
	Fields: graphql.Fields{
		"queryTime": &graphql.Field{
			Type: graphql.String,
		},
		"requestTime": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func getNodeType(nodeName string, edgeType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: nodeName,
		Fields: graphql.Fields{
			"node": &graphql.Field{
				Type: edgeType,
			},
			"cursor": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})
}

func getTypeFields(nodeName string, edgeType *graphql.Object) graphql.Fields {
	return graphql.Fields{
		"edges": &graphql.Field{
			Type: graphql.NewList(getNodeType(nodeName, edgeType)),
		},
		"totalCount": &graphql.Field{
			Type: graphql.Int,
		},
		"pageInfo": &graphql.Field{
			Type: PageInfoType,
		},
		"metrics": &graphql.Field{
			Type: ConnectionMetrics,
		},
	}
}

// // getConnection("UserEmailConnection", "UserEmailNodeType", UserEmailType)
// func getConnection(connectionName string, nodeName string, edgeType *graphql.Object) *graphql.Object {
// 	return graphql.NewObject(graphql.ObjectConfig{
// 		Name: connectionName,
// 		Fields: graphql.Fields{
// 			"edges": &graphql.Field{
// 				Type: graphql.NewList(getNodeType(nodeName, edgeType)),
// 			},
// 			"totalCount": &graphql.Field{
// 				Type: graphql.Int,
// 			},
// 			"pageInfo": &graphql.Field{
// 				Type: PageInfoType,
// 			},
// 			"metrics": &graphql.Field{
// 				Type: ConnectionMetrics,
// 			},
// 		},
// 	})
// }
