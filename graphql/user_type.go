package graphql

import (
	"github.com/graphql-go/graphql"
)

// UserType for graphql
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"user_id": &graphql.Field{
			Type: graphql.String,
		},
		"avatar": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"registered": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func init() {
	// UserType.AddFieldConfig("email", &graphql.Field{
	// 	Type: UserEmailType,
	// 	Args: graphql.FieldConfigArgument{
	// 		"address": &graphql.ArgumentConfig{
	// 			Description: "This is the email address",
	// 			Type:        graphql.NewNonNull(graphql.String),
	// 		},
	// 	},
	// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 		if address, ok := p.Args["address"].(string); ok {
	// 			user := p.Source.(map[string]interface{})
	// 			if userID, ok := user["user_id"]; ok {
	// 				userEmail := email.UserEmailByAddress(userID.(gocql.UUID), address)
	// 				if userEmail != nil {
	// 					return userEmail, nil
	// 				}
	// 			}
	// 		}
	// 		return nil, nil
	// 	},
	// })
	// UserType.AddFieldConfig("attributes", attributes)
	// UserType.AddFieldConfig("emails", emails)
	// UserType.AddFieldConfig("phones", phones)
	// UserType.AddFieldConfig("addresses", addresses)
}
