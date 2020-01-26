package graphql

import (
	"github.com/graphql-go/graphql"
)

// TokenPairType for graphql
var TokenPairType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Token",
	Fields: graphql.Fields{
		"access_token": &graphql.Field{
			Type:        graphql.String,
			Description: "This token is required for any access to the API.",
		},
		"refresh_token": &graphql.Field{
			Type:        graphql.String,
			Description: "This token is required to refresh the access token provided at login.",
		},
	},
})

// AccessTokenType for graphql
var AccessTokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AccessToken",
	Fields: graphql.Fields{
		"access_token": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// RefreshTokenType for graphql
var RefreshTokenType = graphql.NewObject(graphql.ObjectConfig{
	Name: "RefreshToken",
	Fields: graphql.Fields{
		"refresh_token": &graphql.Field{
			Type: graphql.String,
		},
	},
})
