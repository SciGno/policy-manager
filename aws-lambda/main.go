// This is the main package
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/scigno/policy-manager/gremlin"

	"github.com/scigno/policy-manager/auth"
	"github.com/scigno/policy-manager/config"
	"github.com/scigno/policy-manager/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
	gql "github.com/scigno/policy-manager/graphql"
)

// RequestQuery struct
type requestQuery struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

// Schema for rootQuery and rootMutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    gql.GetRootQuery(),
	Mutation: gql.GetRootMutation(),
})

// Handler func
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx = context.WithValue(ctx, gql.ContextKey(gql.GQLLEVEL), map[string]int{})
	ctx = context.WithValue(ctx, gql.ContextKey("sourceIP"), request.RequestContext.Identity.SourceIP)
	ctx = context.WithValue(ctx, gql.ContextKey("resource"), config.ResourceName)

	// logger.Infof("Processing Lambda request %s\n", request.RequestContext.RequestID)
	// s, _ := json.MarshalIndent(request, "", " ")
	// logger.Infof("APIGatewayProxyRequest %+s\n", s)

	ctx = auth.ParseLambdaJWT(ctx, request.Headers)

	if err := gremlin.DialerConfig(config.Keyspace, config.Database); err != nil {
		logger.Error("[gremlin.DialerConfig] Error: ", err)
		// return events.APIGatewayProxyResponse{
		// 	Body:       gql.SimpleJSONFormattedError(request.Path, "internal", gql.InternalError).Error(),
		// 	StatusCode: 500,
		// }, nil
	}
	// defer gremlin.Client.Close()

	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{
			Body:       gql.SimpleJSONFormattedError(request.Path, "access", gql.QueryParameterError).Error(),
			StatusCode: 401,
		}, nil
	}

	req := requestQuery{}

	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		logger.Error("Could not decode body", err)
	}

	// Use this block only in development
	// This is only used to bypass an invalid or missing JWT token
	// TODO: remove block for production
	if userID, ok := auth.ProccessCredentials(ctx); !ok {
		// we don't have a valid JWT.
		// we need to check if any of the fields that don't require JTW are present
		authorized := false
		fields := map[string]struct{}{
			"login":               struct{}{},
			"refresh":             struct{}{},
			"register":            struct{}{},
			"validate":            struct{}{},
			"__schema":            struct{}{},
			"__type":              struct{}{},
			"__Directive":         struct{}{},
			"__DirectiveLocation": struct{}{},
			"__Field":             struct{}{},
			"__InputValue":        struct{}{},
			"__EnumValue":         struct{}{},
			"__TypeKind":          struct{}{},
		}
		source := source.NewSource(&source.Source{
			Body: []byte(req.Query),
			Name: "GraphQL request",
		})

		doc, _ := parser.Parse(parser.ParseParams{Source: source})
		sel := doc.Definitions[0].(*ast.OperationDefinition).GetSelectionSet()
		for _, v := range sel.Selections {
			// logger.Info(v.(*ast.Field).Name.Value)
			if _, authorized = fields[v.(*ast.Field).Name.Value]; authorized {
				break
			}
		}
		if !authorized {
			return events.APIGatewayProxyResponse{
				Body: gql.SimpleJSONFormattedError(request.Path, "access", gql.Unauthorized).Error(),
				// Body:       "{ \"data\": null, \"errors\": [ {\"message\":\"not authorized\"}]} ",
				StatusCode: 401,
			}, nil
		}
	} else {
		ctx = context.WithValue(ctx, auth.ContextKey(auth.UserID), userID)
	}

	// logger.Info("Context: \n", ctx)

	params := graphql.Params{
		Schema:         Schema,
		RequestString:  req.Query,
		OperationName:  req.OperationName,
		VariableValues: req.Variables,
		Context:        ctx,
	}

	result := graphql.Do(params)

	// d, _ := json.MarshalIndent(result, "", "   ")
	// logger.Info("Result: ", string(d))

	responseJSON, err := json.Marshal(result)
	if err != nil {
		logger.Error("Could not decode body")
	}
	// logger.Infof("==> Response: %+s\n", responseJSON)
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", responseJSON),
		StatusCode: 200,
	}, nil

}

func main() {
	// os.Exit(0)
	lambda.Start(Handler)
}
