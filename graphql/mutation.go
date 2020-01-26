package graphql

import (
	"fmt"
	"time"

	"stagezero.com/leandro/marketbin/auth"
	"stagezero.com/leandro/marketbin/logger"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"stagezero.com/leandro/marketbin/enduser"
	"stagezero.com/leandro/marketbin/session"
)

const (
	// AccessTimeout is used for access token duration
	AccessTimeout = 3600 // 1 hour
	// RefreshTimeout is used for refresh token duration
	RefreshTimeout = 2592000 // 30 days
	// MaxLoginAttempts is the maximum login failures a user can attempt before the account is locked
	MaxLoginAttempts = 2
	// LockedGracePeriod is number of minutes an account can re-login after being locked
	LockedGracePeriod = 5
	// MarketerAccessTimeout is used for access token duration
	MarketerAccessTimeout = 300 // 5 minutes
	// MarketerRefreshTimeout is used for refresh token duration
	MarketerRefreshTimeout = 604800 // 7 days
	// MarketerMaxLoginAttempts is the maximum login failures a user can attempt before the account is locked
	MarketerMaxLoginAttempts = 5
)

// GetRootMutation function
func GetRootMutation() *graphql.Object {
	return rootMutation
}

/**
	Root mutation
**/
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"register":    registerUser,
		"login":       login,
		"logout":      logout,
		"refresh":     refresh,
		"credentials": credentials,
		// Attributes
		// "validate":        createAttribute,
		// "createAttribute": createAttribute,
		// "updateAttribute": updateAttribute,
		// "deleteAttribute": deleteAttribute,
	},
})

/**
	User mutations
**/
// SignIn
var registerUser = &graphql.Field{
	Type:        TokenPairType,
	Description: "Register a new user",
	Args: graphql.FieldConfigArgument{
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(String),
		},
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(String),
		},
		"avatar": &graphql.ArgumentConfig{
			Type: String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		password, _ := params.Args["password"].(string)
		username, _ := params.Args["username"].(string)

		// logger.Info("Checking for user...")
		if enduser.UsernameExists(username) {
			return nil, SimpleJSONFormattedError(params.Info.FieldName, "username", UsernameExists)
		}

		userID, _ := uuid.NewRandom()

		pass, err := GenerateUserPasswordHash(password)
		if err != nil {
			return nil, SimpleJSONFormattedError(params.Info.FieldName, "password", WeakPasswordError)
		}

		accessID, _ := uuid.NewRandom()
		refreshID, _ := uuid.NewRandom()
		access, refresh := GenerateTokenPairs(accessID.String(), refreshID.String())
		accessTokenString := string(access)
		refreshTokenString := string(refresh)

		// logger.Info(accessTokenString)
		// logger.Info(refreshTokenString)

		if err := enduser.Create(username, pass, userID); err != nil {
			return nil, SimpleJSONFormattedError(params.Info.FieldName, "user", UnableToCreateUser)
		}

		if err := session.CreateAccessToken(session.Token{UserID: userID.String(), TokenID: accessID.String(), TokenString: accessTokenString}, AccessTimeout); err != nil {
			logger.Errorf("[session.CreateAccessToken] %v", InternalError)
		}
		if err := session.CreateRefreshToken(session.Token{UserID: userID.String(), TokenID: refreshID.String(), TokenString: refreshTokenString}, RefreshTimeout); err != nil {
			logger.Errorf("[session.CreateRefreshToken] %v", InternalError)
		}

		return map[string]interface{}{
			"access_token":  fmt.Sprintf("%s", access),
			"refresh_token": fmt.Sprintf("%s", refresh),
		}, nil
	},
}

var credentials = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Reset user credentials",
	Args: graphql.FieldConfigArgument{
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		ctx := params.Context
		userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

		// if userID, ok := ProccessCredentials(ctx); ok {
		password, _ := params.Args["password"].(string)
		pass, err := GenerateUserPasswordHash(password)
		if err != nil {
			return nil, SimpleJSONFormattedError(params.Info.FieldName, "password", WeakPasswordError)
		}

		enduser.Update(userID, map[string]interface{}{"password": pass})
		return true, nil

		// }
		// return false, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
	},
}

var login = &graphql.Field{
	Type:        TokenPairType,
	Description: "Login a user",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// ctx := params.Context
		// contextErrors := ContextErrors{}

		username, _ := params.Args["username"].(string)
		password, _ := params.Args["password"].(string)

		if users, ok := enduser.GetByUsername(username); ok {
			userID := users.Value.Properties.UserID[0].Value.Value.Value
			status := users.Value.Properties.Status[0].Value.Value
			updated := users.Value.Properties.StatusUpdated[0].Value.Value.Value
			hashedPassword := users.Value.Properties.Password[0].Value.Value
			// log.Println("Username: ", users.Value.Properties.Username[0].Value.Value)
			// log.Println("UserID: ", userID)
			// log.Println("Password: ", users.Value.Properties.Password[0].Value.Value)
			// log.Println("Status: ", status)
			// log.Println("StatusUpdated: ", updated)

			if status == enduser.ACTIVE {
				if !VerifyHashedPassword(hashedPassword, password) { // Password is not valid
					failures, err := enduser.IncreaseLoginFailure(userID)
					if err != nil {
						logger.Errorf(err.Error())
						return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
					}

					if failures >= MaxLoginAttempts {
						if ok, err := enduser.UpdateStatus(userID, enduser.LOCKED); !ok {
							logger.Errorf(err.Error())
							return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
						}
						return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", AccountLocked)
					}
					return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InvalidCredentials)
				}
			} else if status == enduser.LOCKED {
				now := time.Now().UTC()
				if now.Sub(updated).Minutes() >= LockedGracePeriod { // user can login now
					if ok, err := enduser.UpdateStatus(userID, enduser.ACTIVE); !ok {
						logger.Errorf(err.Error())
						return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
					}
				}
				return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", AccountLocked)
			}

			if err := enduser.ResetLoginFailures(userID); err != nil {
				logger.Errorf("[enduser.ResetLoginFailures] %v", err.Error())
				return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
			}

			acessToken, err := session.GetAccessTokenString(userID)
			if err != nil {
				logger.Errorf(err.Error())
				return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
			}

			u, _ := uuid.Parse(userID)

			if acessToken == "" {
				accessID, _ := uuid.NewRandom()
				accessTokenID := accessID.String()
				acc := GenerateAccessToken(accessTokenID)
				accessTokenString := string(acc)

				if err := session.CreateAccessToken(session.Token{UserID: u.String(), TokenID: accessTokenID, TokenString: accessTokenString}, AccessTimeout); err != nil {
					logger.Errorf("[session.CreateAccessToken] %v", InternalError)
					return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
				}
				acessToken = accessTokenString
			}

			refreshToken, err := session.GetRefreshTokenString(userID)
			if err != nil {
				logger.Errorf(err.Error())
				return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
			}

			if refreshToken == "" {
				refreshID, _ := uuid.NewRandom()
				refreshTokenID := refreshID.String()
				refresh := GenerateRefreshToken(refreshTokenID)
				refreshTokenString := string(refresh)
				if err := session.CreateRefreshToken(session.Token{UserID: u.String(), TokenID: refreshTokenID, TokenString: refreshTokenString}, RefreshTimeout); err != nil {
					logger.Errorf("[session.CreateRefreshToken] %v", InternalError)
					return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", InternalError)
				}
				refreshToken = refreshTokenString
			}

			return map[string]interface{}{
				"access_token":  acessToken,
				"refresh_token": refreshToken,
			}, nil

		}
		return nil, SimpleJSONFormattedError(params.Info.FieldName, "login", UserNotFound)
	},
}

var logout = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Logout a particular user.  All tokens belonging to requester will be deleted.",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := params.Context
		userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

		if err := session.DeleteRefreshToken(userID); err != nil {
			logger.Errorf("[mutation.logout] session.DeleteRefreshToken %v", err)
		}
		if err := session.DeleteAccessToken(userID); err != nil {
			logger.Errorf("[mutation.logout] session.DeleteAccessToken %v", err)
		}

		return true, nil
	},
}

var refresh = &graphql.Field{
	Type:        AccessTokenType,
	Description: "Return a new access token. A refresh token is needed for this operation",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ctx := params.Context

		if ok, claims := VerifyJWT(ctx); ok {
			if jti, ok := claims.JWTID(); ok {
				if subject, ok2 := claims.Subject(); ok2 && subject == "refresh" {
					tokenID, _ := uuid.Parse(jti)
					if userID, e := session.GetRefreshTokenUserID(tokenID.String()); e == nil {
						if len(userID) <= 0 {
							return false, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
						}

						acessToken, err := session.GetAccessTokenString(userID)
						if err != nil {
							logger.Errorf(err.Error())
							return nil, SimpleJSONFormattedError(params.Info.FieldName, "refresh", InternalError)
						}

						if len(acessToken) > 0 {
							if err := session.DeleteAccessToken(userID); err != nil {
								return nil, SimpleJSONFormattedError(params.Info.FieldName, "refresh", InternalError)
							}
						}

						accessID, _ := uuid.NewRandom()
						acc := GenerateAccessToken(accessID.String())
						accessTokenString := string(acc)

						if err := session.CreateAccessToken(session.Token{UserID: userID, TokenID: accessID.String(), TokenString: accessTokenString}, AccessTimeout); err != nil {
							logger.Errorf("[session.CreateAccessToken] %v", InternalError)
							return nil, SimpleJSONFormattedError(params.Info.FieldName, "refresh", InternalError)
						}

						return map[string]interface{}{
							"access_token": accessTokenString,
						}, nil
					}
				}
			}
		}
		return false, SimpleJSONFormattedError(params.Info.FieldName, "access", Unauthorized)
	},
}

// Attribute
// var createAttribute = &graphql.Field{
// 	Type:        AttributeType,
// 	Description: "Create a new attribute",
// 	Args: graphql.FieldConfigArgument{
// 		"name": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(enumTypeUserAttributes),
// 		},
// 		"value": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(String),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

// 		name := params.Args["name"].(string)
// 		value := params.Args["value"].(string)
// 		attributeID, _ := uuid.NewRandom()
// 		a := attribute.Attribute{
// 			AttributeID: attributeID.String(),
// 			Name:        name,
// 			Value:       value,
// 		}

// 		if a, err := attribute.Create(a, userID, userID, enduser.USER, enduser.USERID, "has"); err == nil {
// 			return a, nil
// 		}

// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, "internal", InternalError)
// 	},
// }

// var updateAttribute = &graphql.Field{
// 	Type:        AttributeType,
// 	Description: "Update an attribute by name and value",
// 	Args: graphql.FieldConfigArgument{
// 		"attributeId": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(ValidUUID),
// 		},
// 		"value": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(String),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

// 		attributeID := params.Args["attributeId"].(string)
// 		value := params.Args["value"].(string)

// 		if a, ok := attribute.OwnedBy(userID, attributeID, enduser.USER, enduser.USERID, "has"); ok {
// 			if _, err := attribute.Update(userID, attributeID, value, userID, enduser.USER, enduser.USERID, "has"); err == nil {
// 				return a, nil
// 			}
// 			return nil, SimpleJSONFormattedError(params.Info.FieldName, "internal", InternalError)
// 		}
// 		return nil, SimpleJSONFormattedError(params.Info.FieldName, attributeID, AttributeNotFound)
// 		// logger.Error(err)
// 	},
// }

// // Attribute
// var deleteAttribute = &graphql.Field{
// 	Type:        graphql.Boolean,
// 	Description: "Delete an attribute",
// 	Args: graphql.FieldConfigArgument{
// 		"attribute_id": &graphql.ArgumentConfig{
// 			Type: graphql.NewNonNull(String),
// 		},
// 	},
// 	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 		ctx := params.Context
// 		userID := ctx.Value(auth.ContextKey(auth.UserID)).(string)

// 		attributeID, _ := params.Args["attribute_id"].(string)
// 		b, err := attribute.Delete(userID, attributeID, enduser.USER, enduser.USERID, "has")
// 		if b {
// 			logger.Auditf("Actor: %s | Action: %s | Subject: %s", userID, params.Info.FieldName, attributeID)
// 			return true, nil
// 		}
// 		logger.Error(err)
// 		return false, SimpleJSONFormattedError(params.Info.FieldName, "internal", InternalError)
// 	},
// }
