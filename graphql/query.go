package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

/**
====================================
Root query
====================================
**/

// GetRootQuery function
func GetRootQuery() *graphql.Object {
	return rootQuery
}

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"validate": validate,
		// "policies": policies,
	},
})

var validate = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Validate a request",
	Args: graphql.FieldConfigArgument{
		"requests": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.NewList(RequestInputType)),
			Description: "List of Policy requests",
		},
	},
	Resolve: validateResolver,
}

var validateResolver = func(params graphql.ResolveParams) (interface{}, error) {

	c := params.Args["requests"]
	// fmt.Printf("Type: %T\n", c)

	for _, r := range c.([]interface{}) {
		fmt.Printf("Type: %T\n", r)
		for k, f := range r.(map[string]interface{}) {
			fmt.Printf("Name: %v\n", k)
			fmt.Printf("Type: %v\n", f)
			println()
		}

		// if d, err := json.MarshalIndent(f.Value.GetValue(), "", " "); err != nil {
		// 	fmt.Printf("%s", err)
		// } else {
		// 	fmt.Printf("%s", string(d))
		// }
	}

	// ctx := params.Context
	// userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

	// names := []interface{}{}
	// ids := []interface{}{}
	// first := 100
	// after := 1
	// hasNext := true
	// requestDuration := time.Now()

	// if _, ok := params.Args["attribute_ids"]; ok {
	// 	ids = params.Args["attribute_ids"].([]interface{})
	// }

	// if _, ok := params.Args["names"]; ok {
	// 	names = params.Args["names"].([]interface{})
	// }

	// if _, ok := params.Args["first"]; ok {
	// 	if params.Args["first"].(int) > 100 {
	// 		first = 100
	// 	} else {
	// 		first = params.Args["first"].(int)
	// 	}
	// }

	// if _, ok := params.Args["after"]; ok {
	// 	after = params.Args["after"].(int) + 1
	// }

	// graphqlDuration := time.Now()
	// verts, err := attribute.GetList(userID, ids, names, first, after, "user", "userId", "has")
	// graphqlDurationTotal := time.Since(graphqlDuration)
	// vertices := verts[0]

	// // d, _ := json.MarshalIndent(vertices, "", " ")
	// // logger.Info("Verices: ", string(d))

	// if err != nil {
	// 	logger.Errorf("[attribute.GetList] %v", InternalError)
	// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "attributes", InternalError)
	// }

	// nodes := make([]Node, len(vertices.Attributes))
	// for i, v := range vertices.Attributes {
	// 	a := Node{
	// 		Node: attribute.Attribute{
	// 			AttributeID: v.Value.ID.AttributeID.Value,
	// 			Name:        v.Value.Properties.Name[0].Value.Value,
	// 			Value:       v.Value.Properties.Value[0].Value.Value,
	// 			CreatedOn:   v.Value.Properties.CreatedOn[0].Value.Value.Value,
	// 			ModifiedOn:  v.Value.Properties.ModifiedOn[0].Value.Value.Value,
	// 		},
	// 		Cursor: after,
	// 	}
	// 	nodes[i] = a
	// 	after++
	// }

	// after--
	// if after == vertices.TotalCount.Value {
	// 	hasNext = false
	// }
	// requestDurationTotal := time.Since(requestDuration)

	// res := Results{
	// 	TotalCount: vertices.TotalCount.Value,
	// 	Edges:      nodes,
	// 	PageInfo: PageInfo{
	// 		EndCursor:   after,
	// 		HasNextPage: hasNext,
	// 	},
	// 	Metrics: Metrics{
	// 		QueryTime:   graphqlDurationTotal.String(),
	// 		RequestTime: requestDurationTotal.String(),
	// 	},
	// }

	// // res.Connection = Connection{
	// // 	TotalCount: vertices.TotalCount.Value,
	// // 	Edges:      nodes,
	// // 	PageInfo: PageInfo{
	// // 		EndCursor:   after,
	// // 		HasNextPage: hasNext,
	// // 	},
	// // 	Metrics: Metrics{
	// // 		QueryTime:   graphqlDurationTotal.String(),
	// // 		RequestTime: requestDurationTotal.String(),
	// // 	},
	// // }

	// // d, _ := json.MarshalIndent(result, "", " ")
	// // logger.Info("Results: ", string(d))

	return nil, nil
}

// var policies = &graphql.Field{
// 	Type:        PolicyResultsType,
// 	Description: "Get all policies created by a publisher",
// 	Args: graphql.FieldConfigArgument{
// 		"publisherId": &graphql.ArgumentConfig{
// 			Type:        graphql.NewNonNull(ValidUUID),
// 			Description: "Publisher id the creates the policies",
// 		},
// 		"policyIds": &graphql.ArgumentConfig{
// 			Type:        graphql.NewList(ValidUUID),
// 			Description: "List of policy ids",
// 		},
// 		"names": &graphql.ArgumentConfig{
// 			Type:        graphql.NewList(String),
// 			Description: "List of policy names",
// 		},
// 		"first": &graphql.ArgumentConfig{
// 			Type:        NaturalNumber,
// 			Description: "Number of records to retreive",
// 		},
// 		"after": &graphql.ArgumentConfig{
// 			Type:        NaturalNumber,
// 			Description: "Cursor from where to start the list",
// 		},
// 	},
// 	Resolve: policiesResolve,
// }

// var policiesResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	// ctx := params.Context
// 	// userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

// 	names := []interface{}{}
// 	pid := ""
// 	ids := []interface{}{}
// 	first := 100
// 	after := 1
// 	hasNext := true
// 	requestDuration := time.Now()

// 	if _, ok := params.Args["publisherId"]; ok {
// 		pid = params.Args["publisherId"].(string)
// 	}

// 	if _, ok := params.Args["policyIds"]; ok {
// 		ids = params.Args["policyIds"].([]interface{})
// 	}

// 	if _, ok := params.Args["names"]; ok {
// 		names = params.Args["names"].([]interface{})
// 	}

// 	if _, ok := params.Args["first"]; ok {
// 		if params.Args["first"].(int) > 100 {
// 			first = 100
// 		} else {
// 			first = params.Args["first"].(int)
// 		}
// 	}

// 	if _, ok := params.Args["after"]; ok {
// 		after = params.Args["after"].(int) + 1
// 	}

// 	graphqlDuration := time.Now()
// 	verts, err := policy.GetList(pid, ids, names, first, after, publisher.PUBLISHER, publisher.PUBLISHERID, policy.CREATED)
// 	graphqlDurationTotal := time.Since(graphqlDuration)
// 	vertices := verts[0]

// 	if err != nil {
// 		logger.Errorf("[policies.GetList] %v", InternalError)
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "policies", InternalError)
// 	}

// 	nodes := make([]Node, len(vertices.Policies))
// 	for i, v := range vertices.Policies {

// 		resources := []string{}
// 		for _, r := range v.Value.Properties.Resources {
// 			resources = append(resources, r.Value.Value)
// 		}

// 		actions := []string{}
// 		for _, a := range v.Value.Properties.Actions {
// 			resources = append(resources, a.Value.Value)
// 		}

// 		// conditions := map[string]interface{}{}
// 		// for k, c := range v.Value.Properties.Conditions {
// 		// 	conditions[k] = c.Value.Value
// 		// }

// 		e := Node{
// 			Node: policy.Policy{
// 				Name:        v.Value.Properties.Name[0].Value.Value,
// 				PolicyID:    v.Value.ID.PolicyID.Value,
// 				Description: v.Value.Properties.Description[0].Value.Value,
// 				Resources:   resources,
// 				Actions:     actions,
// 				// Conditions:  conditions,
// 				CreatedOn:  v.Value.Properties.CreatedOn[0].Value.Value.Value,
// 				ModifiedOn: v.Value.Properties.ModifiedOn[0].Value.Value.Value,
// 			},
// 			Cursor: after,
// 		}
// 		nodes[i] = e
// 		after++
// 	}

// 	// res := Results{}
// 	after--
// 	if after == vertices.TotalCount.Value {
// 		hasNext = false
// 	}
// 	requestDurationTotal := time.Since(requestDuration)

// 	res := Results{
// 		TotalCount: vertices.TotalCount.Value,
// 		Edges:      nodes,
// 		PageInfo: PageInfo{
// 			EndCursor:   after,
// 			HasNextPage: hasNext,
// 		},
// 		Metrics: Metrics{
// 			QueryTime:   graphqlDurationTotal.String(),
// 			RequestTime: requestDurationTotal.String(),
// 		},
// 	}

// 	// res.Connection = Connection{
// 	// 	TotalCount: vertices.TotalCount.Value,
// 	// 	Edges:      nodes,
// 	// 	PageInfo: PageInfo{
// 	// 		EndCursor:   after,
// 	// 		HasNextPage: hasNext,
// 	// 	},
// 	// 	Metrics: Metrics{
// 	// 		QueryTime:   graphqlDurationTotal.String(),
// 	// 		RequestTime: requestDurationTotal.String(),
// 	// 	},
// 	// }

// 	return res, nil
// }

// // Policy queries
// var getAccessPolicies = &graphql.Field{
// 	Type:        graphql.NewList(PolicyType),
// 	Description: "Ger user access policies associated to marketers",
// 	Args: graphql.FieldConfigArgument{
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		// log.Printf("Context: %+v\n", ctx)
// 		if userID, ok := ProccessCredentials(ctx); ok {
// 			table := policy.TableAccessPolicyByUserID
// 			partitionKey := "user_id"
// 			partitionValue := userID.String()
// 			offsetKey := "policy_id"
// 			offsetValue := ""
// 			limit := 0

// 			if o, ok := params.Args["offset"]; ok {
// 				offsetValue = o.(string)
// 			}
// 			if l, ok := params.Args["limit"]; ok {
// 				limit = l.(int)
// 			}
// 			return policy.GetPaginatedPolicies(partitionKey, partitionValue, offsetKey, offsetValue, table, limit), nil
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 	},
// }

// var getAccessPolicy = &graphql.Field{
// 	Type:        PolicyType,
// 	Description: "Ger user access policy",
// 	Args: graphql.FieldConfigArgument{
// 		"policy_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		if userID, ok := ProccessCredentials(ctx); ok {
// 			policyID, _ := gocql.ParseUUID(params.Args["policy_id"].(string))
// 			if policy.UserHasAccessPolicy(userID, policyID) {
// 				if m, ok := policy.GetMap(policyID); ok {
// 					return m, nil
// 				}
// 			}
// 			return nil, SimpleJSONFormattedError(params.Info.FieldName, policyID.String(), PolicyNotFound)
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 	},
// }

// var getMarketerPolicies = &graphql.Field{
// 	Type:        graphql.NewList(PolicyType),
// 	Description: "Ger all policies associated with marketer",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		// log.Printf("Context: %+v\n", ctx)
// 		if userID, ok := ProccessCredentials(ctx); ok {
// 			marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 			if ProcessAccessPolicy(params, userID, marketerID) {
// 				if m := policy.GetPolicies(marketerID); m != nil {
// 					return m, nil
// 				}
// 				return nil, SimpleJSONFormattedError(params.Info.FieldName, "marketer_id", UserNotFound)
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 	},
// }

// var getPolicy = &graphql.Field{
// 	Type:        PolicyType,
// 	Description: "Get a policy associated with marketer",
// 	Args: graphql.FieldConfigArgument{
// 		"policy_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		// log.Printf("Context: %+v\n", ctx)
// 		if userID, ok := ProccessCredentials(ctx); ok {
// 			policyID, _ := gocql.ParseUUID(params.Args["policy_id"].(string))
// 			if policy.UserHasAccessPolicy(userID, policyID) {
// 				if p, ok := policy.GetMap(policyID); ok {
// 					if marketer.IsDelegate(userID, p["marketer_id"].(gocql.UUID)) {
// 						return p, nil
// 					}
// 					return nil, SimpleJSONFormattedError(params.Info.FieldName, "policy_id", Unauthorized)
// 				}
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 	},
// }

// // Marketer queries
// var getMarketer = &graphql.Field{
// 	Type:        MarketerType,
// 	Description: "Get marketer information",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerResolve,
// }

// var getMarketerResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx := params.Context
// 	if _, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))

// 		if m := marketer.Get(marketerID); m != nil {
// 			return m, nil
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "marketer_id", MarketerNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getMarketers = &graphql.Field{
// 	Type:        graphql.NewList(MarketerType),
// 	Description: "Get all marketers for a user",
// 	Args: graphql.FieldConfigArgument{
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		if userID, ok := ProccessCredentials(ctx); ok {
// 			partitionKey := "user_id"
// 			partitionValue := userID.String()
// 			offsetKey := "attribute_id"
// 			offsetValue := ""

// 			table := marketer.TableMarketerByUser

// 			limit := 0
// 			if o, ok := params.Args["offset"]; ok {
// 				offsetValue = o.(string)
// 			}
// 			if l, ok := params.Args["limit"]; ok {
// 				limit = l.(int)
// 			}
// 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 	},
// }

// // Address queries
// var getMarketerAddress = &graphql.Field{
// 	Type:        MarketerAddressType,
// 	Description: "Get marketer address",
// 	Args: graphql.FieldConfigArgument{
// 		"address_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerAddressResolve,
// }

// var getMarketerAddressResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		addressID, _ := gocql.ParseUUID(params.Args["address_id"].(string))
// 		if a, ok := maddress.GetMap(addressID); ok {
// 			marketerID, _ := a["marketer_id"].(gocql.UUID)
// 			if marketer.IsDelegate(userID, marketerID) {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, addressID.String(), AddressNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getMarketerAddresses = &graphql.Field{
// 	Type:        graphql.NewList(MarketerAddressType),
// 	Description: "Get marketer addresses",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerAddressesResolve,
// }

// var getMarketerAddressesResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		// log.Println("MarketerID: ", marketerID)
// 		// log.Println("UserID: ", userID)
// 		if marketer.IsDelegate(userID, marketerID) {
// 			if a, ok := maddress.GetMapByMarketer(marketerID); ok {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, marketerID.String(), MarketerNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // Emails queries
// var getMarketerEmail = &graphql.Field{
// 	Type:        MarketerEmailType,
// 	Description: "Get marketer email",
// 	Args: graphql.FieldConfigArgument{
// 		"email_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerEmailResolve,
// }

// var getMarketerEmailResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		emailID, _ := gocql.ParseUUID(params.Args["email_id"].(string))
// 		if a, ok := memail.GetMap(emailID); ok {
// 			marketerID, _ := a["marketer_id"].(gocql.UUID)
// 			if marketer.IsDelegate(userID, marketerID) {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, emailID.String(), EmailNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getMarketerEmails = &graphql.Field{
// 	Type:        graphql.NewList(MarketerEmailType),
// 	Description: "Get marketer emails",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerEmailsResolve,
// }

// var getMarketerEmailsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		if marketer.IsDelegate(userID, marketerID) {
// 			if a, ok := memail.GetMapByMarketer(marketerID); ok {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, marketerID.String(), MarketerNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // Phone queries
// var getMarketerPhone = &graphql.Field{
// 	Type:        MarketerPhoneType,
// 	Description: "Get marketer phone",
// 	Args: graphql.FieldConfigArgument{
// 		"phone_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerPhoneResolve,
// }

// var getMarketerPhoneResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		phoneID, _ := gocql.ParseUUID(params.Args["phone_id"].(string))
// 		if a, ok := mphone.GetMap(phoneID); ok {
// 			marketerID, _ := a["marketer_id"].(gocql.UUID)
// 			if marketer.IsDelegate(userID, marketerID) {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, phoneID.String(), PhoneNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getMarketerPhones = &graphql.Field{
// 	Type:        graphql.NewList(MarketerPhoneType),
// 	Description: "Get marketer phone records",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerPhonesResolve,
// }

// var getMarketerPhonesResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		if marketer.IsDelegate(userID, marketerID) {
// 			if a, ok := mphone.GetMapByMarketer(marketerID); ok {
// 				return a, nil
// 			}
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, marketerID.String(), MarketerNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // // Product
// // var getMarketerProduct = &graphql.Field{
// // 	Type:        ProductType,
// // 	Description: "Get marketer product",
// // 	Args: graphql.FieldConfigArgument{
// // 		"product_id": &graphql.ArgumentConfig{
// // 			Type: graphql.NewNonNull(ValidUUID),
// // 		},
// // 	},
// // 	Resolve: getMarketerProductResolve,
// // }

// // var getMarketerProductResolve = func(params graphql.ResolveParams) (interface{}, error) {
// // 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// // 	if !level {
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// // 	}
// // 	if userID, ok := ProccessCredentials(ctx); ok {
// // 		productID, _ := gocql.ParseUUID(params.Args["product_id"].(string))
// // 		if a, ok := product.GetMap(productID); ok {
// // 			marketerID, _ := a["marketer_id"].(gocql.UUID)
// // 			if marketer.IsDelegate(userID, marketerID) {
// // 				return a, nil
// // 			}
// // 		}
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, productID.String(), PhoneNotFound)
// // 	}
// // 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// // }

// // var getMarketerProducts = &graphql.Field{
// // 	Type:        graphql.NewList(ProductType),
// // 	Description: "Get marketer product records",
// // 	Args: graphql.FieldConfigArgument{
// // 		"marketer_id": &graphql.ArgumentConfig{
// // 			Type: graphql.NewNonNull(ValidUUID),
// // 		},
// // 		"offset": &graphql.ArgumentConfig{
// // 			Type: graphql.String,
// // 		},
// // 		"limit": &graphql.ArgumentConfig{
// // 			Type: graphql.Int,
// // 		},
// // 	},
// // 	Resolve: getMarketerProductsResolve,
// // }

// // var getMarketerProductsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// // 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// // 	if !level {
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// // 	}
// // 	if _, ok := ProccessCredentials(ctx); ok {
// // 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// // 		if marketer.Exists(marketerID) {
// // 			partitionKey := "marketer_id"
// // 			partitionValue := marketerID.String()
// // 			offsetKey := "product_id"
// // 			offsetValue := ""

// // 			table := product.TableProductByID

// // 			limit := 0
// // 			if o, ok := params.Args["offset"]; ok {
// // 				offsetValue = o.(string)
// // 			}
// // 			if l, ok := params.Args["limit"]; ok {
// // 				limit = l.(int)
// // 			}
// // 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// // 		}
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, marketerID.String(), MarketerNotFound)
// // 	}
// // 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// // }

// // // Product Attributes
// // var getProductAttribute = &graphql.Field{
// // 	Type:        ProductAttributeType,
// // 	Description: "Get product attribute",
// // 	Args: graphql.FieldConfigArgument{
// // 		"attribute_id": &graphql.ArgumentConfig{
// // 			Type: graphql.NewNonNull(ValidUUID),
// // 		},
// // 	},
// // 	Resolve: getProductAttributeResolve,
// // }

// // var getProductAttributeResolve = func(params graphql.ResolveParams) (interface{}, error) {
// // 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// // 	if !level {
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// // 	}
// // 	if _, ok := ProccessCredentials(ctx); ok {
// // 		attributeID, _ := gocql.ParseUUID(params.Args["attribute_id"].(string))
// // 		if a, ok := product.GetAttributeMap(attributeID); ok {
// // 			return a, nil
// // 		}
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, attributeID.String(), ProductAttributeNotFound)
// // 	}
// // 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// // }

// // var getProductAttributes = &graphql.Field{
// // 	Type:        graphql.NewList(ProductAttributeType),
// // 	Description: "Get marketer product attributes",
// // 	Args: graphql.FieldConfigArgument{
// // 		"product_id": &graphql.ArgumentConfig{
// // 			Type: graphql.NewNonNull(ValidUUID),
// // 		},
// // 		"offset": &graphql.ArgumentConfig{
// // 			Type: graphql.String,
// // 		},
// // 		"limit": &graphql.ArgumentConfig{
// // 			Type: graphql.Int,
// // 		},
// // 	},
// // 	Resolve: getProductAttributesResolve,
// // }

// // var getProductAttributesResolve = func(params graphql.ResolveParams) (interface{}, error) {
// // 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// // 	if !level {
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// // 	}
// // 	if _, ok := ProccessCredentials(ctx); ok {
// // 		productID, _ := gocql.ParseUUID(params.Args["product_id"].(string))
// // 		if product.Exists(productID) {
// // 			partitionKey := "product_id"
// // 			partitionValue := productID.String()
// // 			offsetKey := "attribute_id"
// // 			offsetValue := ""

// // 			table := product.TableProductAttributeByID

// // 			limit := 0
// // 			if o, ok := params.Args["offset"]; ok {
// // 				offsetValue = o.(string)
// // 			}
// // 			if l, ok := params.Args["limit"]; ok {
// // 				limit = l.(int)
// // 			}
// // 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// // 		}
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, productID.String(), ProductNotFound)
// // 	}
// // 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// // }

// // Ad
// var getUserAd = &graphql.Field{
// 	Type:        AdType,
// 	Description: "Get a published ad",
// 	Args: graphql.FieldConfigArgument{
// 		"ad_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getUserAdResolve,
// }

// var getUserAdResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if _, ok := ProccessCredentials(ctx); ok {
// 		adID, _ := gocql.ParseUUID(params.Args["ad_id"].(string))
// 		if a, ok := ad.GetMap(adID); ok {
// 			return a, nil
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, adID.String(), AdNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getUserAds = &graphql.Field{
// 	Type:        graphql.NewList(AdType),
// 	Description: "Get marketer ads",
// 	Args: graphql.FieldConfigArgument{
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: getUserAdsResolve,
// }

// var getUserAdsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if _, ok := ProccessCredentials(ctx); ok {
// 		partitionKey := ""
// 		partitionValue := ""
// 		offsetKey := ""
// 		offsetValue := ""

// 		table := ad.TableAdByID

// 		limit := 0
// 		if o, ok := params.Args["offset"]; ok {
// 			offsetValue = o.(string)
// 		}
// 		if l, ok := params.Args["limit"]; ok {
// 			limit = l.(int)
// 		}
// 		return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // var getAdsByMarketer = &graphql.Field{
// // 	Type:        graphql.NewList(AdType),
// // 	Description: "Get marketer ads",
// // 	Args: graphql.FieldConfigArgument{
// // 		"marketer_id": &graphql.ArgumentConfig{
// // 			Type: graphql.NewNonNull(ValidUUID),
// // 		},
// // 		"offset": &graphql.ArgumentConfig{
// // 			Type: graphql.String,
// // 		},
// // 		"limit": &graphql.ArgumentConfig{
// // 			Type: graphql.Int,
// // 		},
// // 	},
// // 	Resolve: getAdsByMarketerResolve,
// // }

// // var getAdsByMarketerResolve = func(params graphql.ResolveParams) (interface{}, error) {
// // 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// // 	if !level {
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// // 	}
// // 	if _, ok := ProccessCredentials(ctx); ok {
// // 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// // 		if marketer.Exists(marketerID) {
// // 			partitionKey := "marketer_id"
// // 			partitionValue := marketerID.String()
// // 			offsetKey := "ad_id"
// // 			offsetValue := ""

// // 			table := ad.TableAdByID

// // 			limit := 0
// // 			if o, ok := params.Args["offset"]; ok {
// // 				offsetValue = o.(string)
// // 			}
// // 			if l, ok := params.Args["limit"]; ok {
// // 				limit = l.(int)
// // 			}
// // 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// // 		}
// // 		return nil, SimpleJSONFormattedError(params.Info.FieldName, marketerID.String(), MarketerNotFound)
// // 	}
// // 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// // }

// // Subscription

// var getSubscription = &graphql.Field{
// 	Type:        SubscriptionType,
// 	Description: "Get a user subscription",
// 	Args: graphql.FieldConfigArgument{
// 		"subscription_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getSubscriptionResolve,
// }

// var getSubscriptionResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx := params.Context
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		subscriptionID, _ := gocql.ParseUUID(params.Args["subscription_id"].(string))
// 		if a, ok := subscription.GetMap(userID, subscriptionID); ok {
// 			return a, nil
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, params.Args["subscription"].(string), SubscriptionNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getSubscriptions = &graphql.Field{
// 	Type:        graphql.NewList(SubscriptionType),
// 	Description: "Get user subscriptions",
// 	Args: graphql.FieldConfigArgument{
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: getSubscriptionsAdsResolve,
// }

// var getSubscriptionsAdsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		partitionKey := "user_id"
// 		partitionValue := userID.String()
// 		offsetKey := ""
// 		offsetValue := ""

// 		table := subscription.TableSubscriptionByUserID

// 		limit := 0
// 		if o, ok := params.Args["offset"]; ok {
// 			offsetValue = o.(string)
// 		}
// 		if l, ok := params.Args["limit"]; ok {
// 			limit = l.(int)
// 		}
// 		return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // Marketer Ad

// var getMarketerAd = &graphql.Field{
// 	Type:        AdType,
// 	Description: "Get marketer ad",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 		"ad_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getMarketerAdResolve,
// }

// var getMarketerAdResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		if ProcessAccessPolicy(params, userID, marketerID) {
// 			adID, _ := gocql.ParseUUID(params.Args["ad_id"].(string))
// 			if a, ok := ad.GetMap(adID); ok {
// 				return a, nil
// 			}
// 			return nil, SimpleJSONFormattedError(params.Info.FieldName, adID.String(), AdNotFound)
// 		}
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getMarketerAds = &graphql.Field{
// 	Type:        graphql.NewList(AdType),
// 	Description: "Get marketer ads",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: getMarketerAdsResolve,
// }

// var getMarketerAdsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		if ProcessAccessPolicy(params, userID, marketerID) {
// 			partitionKey := "marketer_id"
// 			partitionValue := marketerID.String()
// 			offsetKey := "ad_id"
// 			offsetValue := ""

// 			table := ad.TableAdByID

// 			limit := 0
// 			if o, ok := params.Args["offset"]; ok {
// 				offsetValue = o.(string)
// 			}
// 			if l, ok := params.Args["limit"]; ok {
// 				limit = l.(int)
// 			}
// 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 		}
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getAdsByMarketer = &graphql.Field{
// 	Type:        graphql.NewList(AdType),
// 	Description: "Get marketer ads",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: getAdsByMarketerResolve,
// }

// var getAdsByMarketerResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, AdType.String(), TypeRecursionViolation)
// 	}
// 	if _, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		partitionKey := "marketer_id"
// 		partitionValue := marketerID.String()
// 		offsetKey := "ad_id"
// 		offsetValue := ""

// 		table := ad.TableAdByID

// 		limit := 0
// 		if o, ok := params.Args["offset"]; ok {
// 			offsetValue = o.(string)
// 		}
// 		if l, ok := params.Args["limit"]; ok {
// 			limit = l.(int)
// 		}
// 		return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// // Campaign

// var getCampaign = &graphql.Field{
// 	Type:        CampaignType,
// 	Description: "Get marketer campaign",
// 	Args: graphql.FieldConfigArgument{
// 		"campaign_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 	},
// 	Resolve: getCampaignResolve,
// }

// var getCampaignResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		campaignID, _ := gocql.ParseUUID(params.Args["campaign_id"].(string))
// 		if a, ok := campaign.Get(campaignID); ok {
// 			if ProcessAccessPolicy(params, userID, a.MarketerID) {
// 				campaignID, _ := gocql.ParseUUID(params.Args["campaign_id"].(string))
// 				if a, ok := campaign.GetMap(campaignID); ok {
// 					return a, nil
// 				}
// 			}
// 			return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, campaignID.String(), CampaignAdNotFound)
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }

// var getCampaigns = &graphql.Field{
// 	Type:        graphql.NewList(CampaignType),
// 	Description: "Get marketer campaigns",
// 	Args: graphql.FieldConfigArgument{
// 		"marketer_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 		"offset": &graphql.ArgumentConfig{
// 			Type: graphql.String,
// 		},
// 		"limit": &graphql.ArgumentConfig{
// 			Type: graphql.Int,
// 		},
// 	},
// 	Resolve: getCampaignsResolve,
// }

// var getCampaignsResolve = func(params graphql.ResolveParams) (interface{}, error) {
// 	ctx, level := ProcessGQLLevel(params.Context, params.Info.ReturnType.Name())
// 	if !level {
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "query", TypeRecursionViolation)
// 	}
// 	if userID, ok := ProccessCredentials(ctx); ok {
// 		marketerID, _ := gocql.ParseUUID(params.Args["marketer_id"].(string))
// 		if ProcessAccessPolicy(params, userID, marketerID) {
// 			partitionKey := "marketer_id"
// 			partitionValue := marketerID.String()
// 			offsetKey := "campaign_id"
// 			offsetValue := ""

// 			table := campaign.TableCampaignByID

// 			limit := 0
// 			if o, ok := params.Args["offset"]; ok {
// 				offsetValue = o.(string)
// 			}
// 			if l, ok := params.Args["limit"]; ok {
// 				limit = l.(int)
// 			}
// 			return cassandra.Client.GetPaginatedRecords(partitionKey, partitionValue, offsetKey, offsetValue, table, "", limit), nil
// 		}
// 	}
// 	return nil, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
// }
