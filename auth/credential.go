package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/scigno/policy-manager/session"
)

const (
	// UserID is used for Request context
	UserID = "UserID"
)

// ProccessCredentials function validated JWT and jit.
func ProccessCredentials(ctx context.Context) (string, bool) {
	// log.Println("Processing credentials...")
	if ok, claims := VerifyJWT(ctx); ok {
		if jti, ok := claims.JWTID(); ok {
			tokenID, _ := uuid.Parse(jti)
			if userID, e := session.GetAccessTokenUserID(tokenID.String()); e == nil {
				// id, _ := uuid.Parse(userID)
				if len(userID) > 0 {
					return userID, true
				}
				return userID, false
			}
		}
	}
	return "", false
}
