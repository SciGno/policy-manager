package graphql

// import (
// 	"context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/graphql-go/graphql"
// 	"stagezero.com/leandro/marketbin/api"
// )

// // RequestQuery struct
// type requestQuery struct {
// 	Query         string                 `json:"query"`
// 	OperationName string                 `json:"operationName"`
// 	Variables     map[string]interface{} `json:"variables"`
// }

// // Schema for rootQuery and rootMutation
// var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
// 	Query:    rootQuery,
// 	Mutation: rootMutation,
// })

// func executeQuery(ctx context.Context, req *requestQuery, schema graphql.Schema) *graphql.Result {
// 	result := graphql.Do(graphql.Params{
// 		Schema:         schema,
// 		RequestString:  req.Query,
// 		OperationName:  req.OperationName,
// 		VariableValues: variables,
// 		Context:        ctx,
// 	})

// 	return result
// }

// // Query function over HTTP GET
// // =============================
// // example: curl -s -g 'http://localhost:8080/graphql?query={user(user_id:"16255f11-7a75-4d72-88f6-3eba4abea16f"){user_id,first_name,last_name}}' | python -m json.tool
// // func Query(w http.ResponseWriter, r *http.Request) {
// // 	// result := executeQuery(r.URL.Query().Get("query"), schema)
// // 	result := graphql.Do(graphql.Params{
// // 		Schema:        schema,
// // 		RequestString: r.URL.Query().Get("query"),
// // 	})
// // 	json.NewEncoder(w).Encode(result)
// // }

// // GraphQL function over HTTP POST
// // ================================
// func GraphQL(w http.ResponseWriter, r *http.Request) {

// 	req := &requestQuery{}
// 	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 		api.ResponseError(w, api.JSONValidationError)
// 		return
// 	}

// 	result := executeQuery(r.Context(), req, Schema)
// 	// if result.HasErrors() {
// 	// 	logger.Error("Result errors: %v", result.Errors)
// 	// }
// 	json.NewEncoder(w).Encode(result)
// }
