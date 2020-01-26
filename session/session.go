package session

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scigno/gizmo"
	"stagezero.com/leandro/marketbin/db/cassandra"
	"stagezero.com/leandro/marketbin/enduser"
	"stagezero.com/leandro/marketbin/gremlin"
	"stagezero.com/leandro/marketbin/logger"
)

const (
	// TableAccessToken map
	TableAccessToken = "access_token"
	// TableAccessTokenByID map
	TableAccessTokenByID = "access_token_by_id"
	// TableRefreshToken map
	TableRefreshToken = "refresh_token"
	// TableRefreshTokenByID map
	TableRefreshTokenByID = "refresh_token_by_id"
	// TableLoginAttempts map
	TableLoginAttempts = "login_attempts"
	// LOCKED map
	LOCKED = "locked"
	// STATUS map
	STATUS = "status"
	// PASSWORD map
	PASSWORD = "password"
)

// Token data structure
type Token struct {
	UserID      string `json:"user_id"`
	TokenID     string `json:"token_id"`
	TokenString string `json:"token_string"`
}

// LoginAttempt data structure
type LoginAttempt struct {
	UserID string `json:"user_id"`
	Date   string `json:"date"`
}

// CreateAccessToken a new session with duration in seconds
// If duration is zero, session will have no expiration
func CreateAccessToken(data Token, duration int) error {

	type response []struct {
		Type  string `json:"@type"`
		Value struct {
			ID struct {
				AccessTokenID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"accessTokenId"`
				Label string `json:"~label"`
			} `json:"id"`
			Label string `json:"label"`
		} `json:"@value"`
	}

	g := gizmo.NewGraph()

	g.AddV("accessToken").
		Property("accessTokenId", data.TokenID).
		Property("accessTokenString", data.TokenString).
		Property("createdBy", data.UserID).
		Property("modifiedBy", data.UserID).
		Property("createdOn", time.Now().UnixNano()/1000000).
		Property("modifiedOn", time.Now().UnixNano()/1000000).
		As("a").V().Has("user", "userId", data.UserID).As("u").AddE("has").From("u").To("a")

	// logger.Info("[session.CreateAccessToken] Gremlin Query: ", g)
	_, err := gremlin.Client.Execute(g.String(), "session.CreateAccessToken")

	if err != nil {
		logger.Error("%+v", err)
		return err
	}

	// d, _ := json.MarshalIndent(res.([]interface{})[0], "", " ")
	// // logger.Info("[session.CreateAccessToken] Gremlin Response: ", string(d))

	// result := response{}
	// json.Unmarshal(d, &result)
	// logger.Info("[session.CreateAccessToken] Result: ", result)

	return nil
}

// CreateRefreshToken a new session with duration in seconds
// If duration is zero, session will have no expiration
func CreateRefreshToken(data Token, duration int) error {

	type response []struct {
		Type  string `json:"@type"`
		Value struct {
			ID struct {
				RefreshTokenID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"refreshTokenId"`
				Label string `json:"~label"`
			} `json:"id"`
			Label string `json:"label"`
		} `json:"@value"`
	}

	g := gizmo.NewGraph()
	g.AddV("refreshToken").
		Property("refreshTokenId", data.TokenID).
		Property("refreshTokenString", data.TokenString).
		Property("createdBy", data.UserID).
		Property("modifiedBy", data.UserID).
		Property("createdOn", time.Now().UnixNano()/1000000).
		Property("modifiedOn", time.Now().UnixNano()/1000000).
		As("r").V().Has("user", "userId", data.UserID).As("u").AddE("has").From("u").To("r")

	// logger.Info("[session.CreateRefreshToken] Gremlin Query: ", query)
	_, err := gremlin.Client.Execute(g.String(), "session.CreateRefreshToken")

	if err != nil {
		logger.Error("%+v", err)
		return err
	}

	return nil
}

// func addTokenEdge(vertex string, key string, data Token) error {

// 	type response []struct {
// 		Type  string `json:"@type"`
// 		Value struct {
// 			ID struct {
// 				InVertex struct {
// 					RefreshTokenID struct {
// 						Type  string `json:"@type"`
// 						Value string `json:"@value"`
// 					} `json:"refreshTokenId"`
// 					Label string `json:"~label"`
// 				} `json:"~in_vertex"`
// 				Label   string `json:"~label"`
// 				LocalID struct {
// 					Type  string `json:"@type"`
// 					Value string `json:"@value"`
// 				} `json:"~local_id"`
// 				OutVertex struct {
// 					UserID struct {
// 						Type  string `json:"@type"`
// 						Value string `json:"@value"`
// 					} `json:"userId"`
// 					Label string `json:"~label"`
// 				} `json:"~out_vertex"`
// 			} `json:"id"`
// 			InV struct {
// 				RefreshTokenID struct {
// 					Type  string `json:"@type"`
// 					Value string `json:"@value"`
// 				} `json:"refreshTokenId"`
// 				Label string `json:"~label"`
// 			} `json:"inV"`
// 			InVLabel string `json:"inVLabel"`
// 			Label    string `json:"label"`
// 			OutV     struct {
// 				UserID struct {
// 					Type  string `json:"@type"`
// 					Value string `json:"@value"`
// 				} `json:"userId"`
// 				Label string `json:"~label"`
// 			} `json:"outV"`
// 			OutVLabel string `json:"outVLabel"`
// 		} `json:"@value"`
// 	}

//
// 	g := gizmo.NewGraph()
// 	query := g.New()
// 	getUser := g.New().Raw("u=").Append(g.New("g").V().Has("user", "userId", data.UserID).Next())
// 	getToken := g.New().Raw("t=").Append(g.New("g").V().Has(vertex, key, data.TokenID).Next())
// 	addEdge := g.New("g").V(g.Var("u")).AddE("has").To(g.Var("t")).Next()
// 	query.Append(getUser).AddLine(getToken).AddLine(addEdge)

// 	logger.Info("[session.addTokenEdge] Gremlin Query: ", query)
// 	res, err := gremlin.Client.Execute(query.String())

// 	if err != nil {
// 		logger.Error("%+v", err)
// 		return err
// 	}

// 	d, err := json.MarshalIndent(res.([]interface{})[0], "", " ")
// 	if err != nil {
// 		logger.Error("%+v", err)
// 		return err
// 	}

// 	logger.Info("[session.addTokenEdge] Gremlin Response: ", string(d))
// 	result := response{}
// 	json.Unmarshal(d, &result)

// 	logger.Info("[session.addTokenEdge] Result: ", result)

// 	return nil
// }

// ===
// Refresh token
// ===

// GetAccessToken the specified user based on a token id
func GetAccessToken(userID uuid.UUID) map[string]interface{} {
	q := fmt.Sprintf("SELECT * FROM %s WHERE user_id=%v", TableAccessToken, userID)
	if m := cassandra.Client.Read(q); m != nil {
		return m
	}
	return nil
}

// GetAccessTokenID returns the user id based on a token id
func GetAccessTokenID(userID uuid.UUID) (uuid.UUID, bool) {
	q := fmt.Sprintf("SELECT token_id FROM %s WHERE user_id=%v", TableAccessToken, userID)
	if m := cassandra.Client.Read(q); m != nil {
		return m["token_id"].(uuid.UUID), true
	}
	return uuid.UUID{}, false
}

// GetAccessTokenString returns the access token string
// It returns an error if request fails
func GetAccessTokenString(id string) (string, error) {

	type Response []struct {
		Type  string `json:"@type"`
		Value struct {
			ID struct {
				AccessTokenID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"accessTokenId"`
				Label string `json:"~label"`
			} `json:"id"`
			Label      string `json:"label"`
			Properties struct {
				AccessTokenID []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"accessTokenId"`
				AccessTokenString []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value string `json:"value"`
					} `json:"@value"`
				} `json:"accessTokenString"`
				CreatedBy []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"createdBy"`
				CreatedOn []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string    `json:"@type"`
							Value time.Time `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"createdOn"`
				ModifiedBy []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"modifiedBy"`
				ModifiedOn []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								AccessTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"accessTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string    `json:"@type"`
							Value time.Time `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"modifiedOn"`
			} `json:"properties"`
		} `json:"@value"`
	}

	g := gizmo.NewGraph()
	g.V().Has("user", "userId", id).OutE("has").InV().HasLabel("accessToken")

	// logger.Info("[session.GetAccessTokenString] Gremling: ", accessToken.String())
	res, err := gremlin.Client.Execute(g.String(), "session.GetAccessTokenString")
	// logger.Info("[session.GetAccessTokenString] Response: ", res)
	if err != nil {
		return "", err
	}

	response := res.([]interface{})[0]
	if response == nil {
		return "", nil
	}

	d, _ := json.MarshalIndent(response, "", " ")
	// logger.Info("JSON Response: ", string(d))
	r := Response{}

	if err := json.Unmarshal(d, &r); err != nil {
		return "", err
	}

	token := r[0].Value.Properties.AccessTokenString[0].Value.Value

	return token, nil
}

// GetAccessTokenUserID returns the user id if found or empty string.
// Returns an error if gremlin request fails
func GetAccessTokenUserID(tokenID string) (string, error) {

	g := gizmo.NewGraph()
	g.V().Has("accessToken", "accessTokenId", tokenID).In().HasLabel("user")

	// logger.Info("[session.GetAccessTokenString] Gremling: ", token.String())
	res, err := gremlin.Client.Execute(g.String(), "session.GetAccessTokenUserID")
	// logger.Info("[session.GetAccessTokenUserID] Response: ", res)
	if err != nil {
		return "", err
	}

	response := res.([]interface{})[0]
	if response == nil {
		return "", nil
	}

	d, _ := json.MarshalIndent(response, "", " ")
	// logger.Info("[session.GetAccessTokenUserID] JSON Response: ", string(d))
	r := enduser.UserVertex{}

	if err := json.Unmarshal(d, &r); err != nil {
		return "", err
	}

	userID := r[0].Value.ID.UserID.Value

	// logger.Info("[session.GetAccessTokenUserID] UserID: ", userID)
	return userID, nil
}

// DeleteAccessToken user in the database
func DeleteAccessToken(userID string) error {

	g := gizmo.NewGraph()
	g.V().Has("user", "userId", userID).OutE("has").InV().HasLabel("accessToken").Drop()
	// logger.Info("[session.DeleteAccessToken] Gremlin: ", g.String())

	_, err := gremlin.Client.Execute(g.String(), "session.DeleteAccessToken")
	if err != nil {
		return err
	}

	return nil
}

// ===
// Refresh token
// ===

// GetRefreshToken the specified user based on a token id
func GetRefreshToken(userID uuid.UUID) map[string]interface{} {
	q := fmt.Sprintf("SELECT * FROM %s WHERE user_id=%v", TableRefreshToken, userID)
	if m := cassandra.Client.Read(q); m != nil {
		return m
	}
	return nil
}

// GetRefreshTokenID returns the user id based on a token id
func GetRefreshTokenID(userID uuid.UUID) (uuid.UUID, bool) {
	q := fmt.Sprintf("SELECT token_id FROM %s WHERE user_id=%v", TableRefreshToken, userID)
	if m := cassandra.Client.Read(q); m != nil {
		return m["token_id"].(uuid.UUID), true
	}
	return uuid.UUID{}, false
}

// GetRefreshTokenString the specified user based on a token id
func GetRefreshTokenString(id string) (string, error) {
	type Response []struct {
		Type  string `json:"@type"`
		Value struct {
			ID struct {
				RefreshTokenID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"refreshTokenId"`
				Label string `json:"~label"`
			} `json:"id"`
			Label      string `json:"label"`
			Properties struct {
				CreatedBy []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"createdBy"`
				CreatedOn []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string    `json:"@type"`
							Value time.Time `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"createdOn"`
				ModifiedBy []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"modifiedBy"`
				ModifiedOn []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string    `json:"@type"`
							Value time.Time `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"modifiedOn"`
				RefreshTokenID []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value struct {
							Type  string `json:"@type"`
							Value string `json:"@value"`
						} `json:"value"`
					} `json:"@value"`
				} `json:"refreshTokenId"`
				RefreshTokenString []struct {
					Type  string `json:"@type"`
					Value struct {
						ID struct {
							Label   string `json:"~label"`
							LocalID struct {
								Type  string `json:"@type"`
								Value string `json:"@value"`
							} `json:"~local_id"`
							OutVertex struct {
								RefreshTokenID struct {
									Type  string `json:"@type"`
									Value string `json:"@value"`
								} `json:"refreshTokenId"`
								Label string `json:"~label"`
							} `json:"~out_vertex"`
						} `json:"id"`
						Label string `json:"label"`
						Value string `json:"value"`
					} `json:"@value"`
				} `json:"refreshTokenString"`
			} `json:"properties"`
		} `json:"@value"`
	}

	g := gizmo.NewGraph()
	g.V().Has("user", "userId", id).OutE("has").InV().HasLabel("refreshToken")

	// logger.Info("[session.GetRefreshTokenString] Gremling: ", accessToken.String())
	res, err := gremlin.Client.Execute(g.String(), "session.GetRefreshTokenString")
	// logger.Info("[session.GetRefreshTokenString] Response: ", res)
	if err != nil {
		return "", err
	}

	response := res.([]interface{})[0]
	if response == nil {
		return "", err
	}

	d, _ := json.MarshalIndent(response, "", " ")
	// logger.Info("JSON Response: ", string(d))
	r := Response{}

	if err := json.Unmarshal(d, &r); err != nil {
		return "", err
	}

	token := r[0].Value.Properties.RefreshTokenString[0].Value.Value

	return token, nil
}

// GetRefreshTokenUserID returns the user id based on a token id
func GetRefreshTokenUserID(tokenID string) (string, error) {

	type Response []struct {
		Type  string `json:"@type"`
		Value struct {
			ID struct {
				Label   string `json:"~label"`
				LocalID struct {
					Type  string `json:"@type"`
					Value string `json:"@value"`
				} `json:"~local_id"`
				OutVertex struct {
					UserID struct {
						Type  string `json:"@type"`
						Value string `json:"@value"`
					} `json:"userId"`
					Label string `json:"~label"`
				} `json:"~out_vertex"`
			} `json:"id"`
			Label string `json:"label"`
			Value struct {
				Type  string `json:"@type"`
				Value string `json:"@value"`
			} `json:"value"`
		} `json:"@value"`
	}

	g := gizmo.NewGraph()
	g.V().Has("refreshToken", "refreshTokenId", tokenID).In().HasLabel("user").Properties("userId")

	// logger.Info("[session.GetRefreshTokenUserID] Gremling: ", g.String())
	res, err := gremlin.Client.Execute(g.String(), "session.GetRefreshTokenUserID")
	// logger.Info("[session.GetRefreshTokenUserID] Response: ", res)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	response := res.([]interface{})[0]
	if response == nil {
		return "", nil
	}

	d, _ := json.MarshalIndent(response, "", " ")
	// logger.Info("JSON Response: ", string(d))
	r := Response{}

	if err := json.Unmarshal(d, &r); err != nil {
		logger.Error(err)
		return "", err
	}

	id := r[0].Value.Value.Value

	return id, nil
}

// DeleteRefreshToken user in the database
func DeleteRefreshToken(userID string) error {

	g := gizmo.NewGraph()
	g.V().Has("user", "userId", userID).OutE("has").InV().HasLabel("refreshToken").Drop()

	_, err := gremlin.Client.Execute(g.String(), "session.DeleteRefreshToken")
	if err != nil {
		return err
	}

	return nil
}

// // GetTTL the specified user based on an UUID
// func GetTTL(id uuid.UUID) int {
// 	q := fmt.Sprintf("SELECT TTL(user_id) FROM %s WHERE token_id=%v", TableAccessToken, id)
// 	// logger.Info("Query: %+v", q)
// 	if m := cassandra.Client.Read(q); m != nil {
// 		// logger.Info("[Session] Map: %+v", m)
// 		ttl := m["user_id"].(int)
// 		// logger.Info("[Session] GetUserID: %+v", uid)
// 		return ttl
// 	}
// 	return 0
// }
