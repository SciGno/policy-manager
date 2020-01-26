package graphql

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SermoDigital/jose/jwt"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
	"stagezero.com/leandro/marketbin/auth"
	"stagezero.com/leandro/marketbin/logger"
)

const (
	// GQLLEVEL constant
	GQLLEVEL = "GraphQLLevel"
	//GQLMAXLEVELS constant
	GQLMAXLEVELS = 2
)

// GenerateUserPasswordHash function
func GenerateUserPasswordHash(pass string) (string, error) {
	var password string
	score := zxcvbn.PasswordStrength(pass, nil)
	if score.Score >= 0 {
		// Generate "hash" to store from user password
		hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			logger.Errorf(err.Error())
			return "", errors.New("unable to generate password hash")
		}
		password = fmt.Sprintf("%s", hash)
	} else {
		return "", errors.New("password is weak")
	}
	return password, nil
}

// func isAuthorized(params graphql.ResolveParams) (gocql.UUID, bool) {

// 	// logger.Info("Token Found: %+v", params.Context.Value(auth.ContextKey(auth.JWTTokenFound)))
// 	// logger.Info("Token Valid: %+v", params.Context.Value(auth.ContextKey(auth.JWTTokenValid)))
// 	// logger.Info("JWTClaim: %+v", params.Context.Value(auth.ContextKey(auth.JWTClaim)))

// 	token := params.Context.Value(auth.ContextKey(auth.JWTTokenValid))
// 	claims := params.Context.Value(auth.ContextKey(auth.JWTClaim)).(jwt.Claims)
// 	if token != nil {
// 		if token.(bool) {
// 			if jwtID, ok := claims.JWTID(); ok {
// 				if id, err := gocql.ParseUUID(jwtID); err == nil {
// 					logger.Info("JWTID: %+v", id)
// 					if uid, ok := session.GetAccessTokenUserID(id); ok {
// 						logger.Info("UserID: %+v", uid)
// 						return uid, true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return gocql.UUID{}, false
// }

// VerifyUserPassword function
// func VerifyUserPassword(uuid gocql.UUID, password string) bool {
// 	userHashPassword := enduser.GetUserPassword(uuid)["password"].(string)
// 	// Comparing the password with the hash
// 	err := bcrypt.CompareHashAndPassword([]byte(userHashPassword), []byte(password))
// 	if err != nil {
// 		logger.Error("[VerifyUserPassword] %v", err)
// 		return false
// 	}
// 	return true
// }

// // VerifyPasswordByUsername function
// func VerifyPasswordByUsername(username string, password string) bool {
// 	userHashPassword := enduser.GetPasswordByUsername(username)
// 	// Comparing the password with the hash
// 	err := bcrypt.CompareHashAndPassword([]byte(userHashPassword), []byte(password))
// 	if err != nil {
// 		// logger.Error("[VerifyPasswordByUsername] %v", err)
// 		return false
// 	}
// 	return true
// }

// VerifyHashedPassword function
func VerifyHashedPassword(userHashPassword string, password string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(userHashPassword), []byte(password))
	if err != nil {
		// logger.Error("[VerifyPasswordByUsername] %v", err)
		return false
	}
	return true
}

// GenerateTokenPairs returns an access and refresh token
func GenerateTokenPairs(a string, r string) ([]byte, []byte) {
	return GenerateAccessToken(a), GenerateRefreshToken(r)
}

// GenerateRefreshToken returns an access and refresh token
func GenerateRefreshToken(r string) []byte {
	expiration := time.Now().UTC().Add(time.Duration(RefreshTimeout) * time.Second)
	// log.Printf("[GenerateRefreshToken] Expiration: %v\n", expiration)
	// Refresh token
	refresh, _ := auth.NewJWSToken(
		r,
		expiration,
		"api.mydomain.com",
		"mydomain.com",
		time.Now().UTC(),
		"refresh",
		map[string]interface{}{
			"policy": nil,
		}, // map of custom claims
	)
	return refresh
}

// GenerateAccessToken returns an access and refresh token
func GenerateAccessToken(a string) []byte {
	expiration := time.Now().UTC().Add(time.Duration(AccessTimeout) * time.Second)
	// log.Printf("[GenerateAccessToken] Expiration: %v\n", expiration)

	// Access token
	access, _ := auth.NewJWSToken(
		a,
		expiration,
		"api.mydomain.com",
		"mydomain.com",
		time.Now().UTC(),
		"access",
		nil,
		// p, // map of custom claims
	)
	return access
}

// VerifyJWT function
func VerifyJWT(ctx context.Context) (bool, jwt.Claims) {

	// logger.Info("Context: %+v", ctx)

	found := ctx.Value(auth.ContextKey(auth.JWTTokenFound)).(bool)
	valid := ctx.Value(auth.ContextKey(auth.JWTTokenValid)).(bool)

	claims := ctx.Value(auth.ContextKey(auth.JWTClaim))
	var subject, tokenID string
	if claims != nil {
		subject, _ = (claims.(jwt.Claims)).Subject()
		tokenID, _ = (claims.(jwt.Claims)).JWTID()

	}
	if found && valid {
		// logger.Infof("[VerifyJWT] %s token %s is valid", subject, tokenID)
		if claims != nil {
			return true, claims.(jwt.Claims)
		}
		return false, nil
	}

	if claims != nil {
		logger.Errorf("[VerifyJWT] %s token %s is invalid", subject, tokenID)
	} else {
		logger.Errorf("[VerifyJWT] no token found")
	}
	return false, nil
}

// // ProcessAccessPolicy function
// func ProcessAccessPolicy(params graphql.ResolveParams, userID gocql.UUID, marketerID gocql.UUID) bool {
// 	allowAccess := []string{}
// 	denyAccess := []string{}

// 	if a, ok := policy.GetAllowAccessValuesByUser(userID, marketerID); ok {
// 		allowAccess = a
// 	}

// 	if d, ok := policy.GetDenyAccessValuesByUser(userID, marketerID); ok {
// 		denyAccess = d
// 	}
// 	// log.Printf("Allow access: %+v\n", allowAccess)
// 	// log.Printf("Deny access: %+v\n", denyAccess)

// 	for _, v := range denyAccess {
// 		if v == params.Info.FieldName || v == "*" {
// 			log.Printf("Matched deny rule: %+v\n", v)
// 			return false
// 		}
// 	}

// 	for _, v := range allowAccess {
// 		if v == params.Info.FieldName || v == "*" {
// 			log.Printf("Matched allow rule: %+v\n", v)
// 			return true
// 		}
// 	}
// 	return false
// }

// // ProcessPolicy function
// func ProcessPolicy(ctx context.Context, params graphql.ResolveParams) bool {
// 	claims := ctx.Value(auth.ContextKey(auth.JWTClaim)).(jwt.Claims)
// 	if claims.Get("policy") != nil {
// 		p := claims.Get("policy").(map[string]interface{})
// 		if p["deny_access"] != nil {
// 			denyAccess := p["deny_access"].([]interface{})
// 			log.Printf("Processing deny rules: %+v\n", p["deny_access"])
// 			for _, v := range denyAccess {
// 				if v.(string) == params.Info.FieldName || v.(string) == "*" {
// 					log.Printf("Matched deny rule: %+v\n", v)
// 					return false
// 				}
// 			}
// 		}
// 		if p["allow_access"] != nil {
// 			log.Printf("Processing allow rules: %+v\n", p["allow_access"])
// 			allowAccess := p["allow_access"].([]interface{})
// 			for _, v := range allowAccess {
// 				if v.(string) == params.Info.FieldName || v.(string) == "*" {
// 					log.Printf("Matched allow rule: %+v\n", v)
// 					return true
// 				}
// 			}
// 		}
// 	}

// 	// log.Printf("Root: %+v\n", p)
// 	// log.Printf("AllowAccess: %+v\n", p["allow_access"])
// 	// log.Printf("Access Path: %s:%s", params.Info.ParentType.Name(), params.Info.FieldName)
// 	return false
// }

// // ProccessCredentials function validated JWT and jit.
// func ProccessCredentials(ctx context.Context) (string, bool) {
// 	// log.Println("Processing credentials...")
// 	if ok, claims := VerifyJWT(ctx); ok {
// 		if jti, ok := claims.JWTID(); ok {
// 			tokenID, _ := uuid.Parse(jti)
// 			if userID, e := session.GetAccessTokenUserID(tokenID.String()); e == nil {
// 				// id, _ := uuid.Parse(userID)
// 				if len(userID) > 0 {
// 					return userID, true
// 				}
// 				return userID, false
// 			}
// 		}
// 	}
// 	return "", false
// }

// func PrintJSON(i interface{}) []byte {
// 	json.MarshalIndent(data, "", "    ")
// }

// ProcessGQLLevel func
func ProcessGQLLevel(ctx context.Context, name string) (context.Context, bool) {
	log.Println("Processing level: ", name)
	levels := ctx.Value(ContextKey(GQLLEVEL)).(map[string]int)
	log.Println("Levels Map 1: ", levels)
	if l, ok := levels[name]; ok {
		if levels[name] > GQLMAXLEVELS {
			return ctx, false
		}
		l++
		levels[name] = l
		log.Println("Levels Map 1: ", levels)
		return ctx, false
	}
	levels[name] = 1
	log.Println("Levels Map 2: ", levels)
	return ctx, true
}
