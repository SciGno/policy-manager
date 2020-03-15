package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/scigno/policy-manager/logger"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jwt"
)

// ContextKey type
type ContextKey string

const (
	// JWTTokenFound is used for Request context
	JWTTokenFound = "JWTTokenFound"
	// JWTTokenValid is used for Request context
	JWTTokenValid = "JWTTokenValid"
	// JWTTokenExpired is used for Request context
	JWTTokenExpired = "JWTTokenExpired"
	// JWTTokenNotYetValid is used for Request context
	JWTTokenNotYetValid = "JWTTokenNotYetValid"
	// JWTClaim is used for Request context
	JWTClaim = "JWTClaim"
	// JWTTokenID represents the JWT jti property
	JWTTokenID = "JWTTokenID"
)

//-----------------------------------------

// ParseLambdaJWT function
func ParseLambdaJWT(ctx context.Context, headers map[string]string) context.Context {

	ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), false)
	ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), false)
	ctx = context.WithValue(ctx, ContextKey(JWTTokenExpired), false)
	ctx = context.WithValue(ctx, ContextKey(JWTTokenNotYetValid), false)

	if ah := headers["Authorization"]; len(ah) > 7 && strings.EqualFold(ah[0:7], "BEARER ") {
		token := []byte(ah[7:])
		// log.Printf("JWT Token: %+s\n", token)
		if headerToken, err := jws.ParseJWT(token); err == nil {
			ctx = context.WithValue(ctx, ContextKey(JWTTokenFound), true)
			// Validate token
			var validationErr error
			for _, v := range publicKeys {
				rsaPublic, _ := crypto.ParseRSAPublicKeyFromPEM([]byte(v))
				// logger.Info("[JWTSecuredHandler] publicKeys: %v", rsaPublic)
				validationErr = headerToken.Validate(rsaPublic, crypto.SigningMethodRS256)
				if validationErr == nil {
					ctx = context.WithValue(ctx, ContextKey(JWTTokenValid), true)
					ctx = context.WithValue(ctx, ContextKey(JWTClaim), headerToken.Claims())
					// logger.Info("CLAIMS: ", headerToken.Claims())
					if jti, ok := headerToken.Claims().JWTID(); ok {
						ctx = context.WithValue(ctx, ContextKey(JWTTokenID), jti)
					}
					break
				}
			}

			s, _ := headerToken.Claims().Subject()
			tokenValid := headerToken.Claims().Validate(time.Now().UTC(), time.Duration(0), time.Duration(0))
			if tokenValid == jwt.ErrTokenIsExpired {
				logger.Infof("[JWTSecuredHandler] %s token expired", s)
				ctx = context.WithValue(ctx, ContextKey(JWTTokenExpired), true)
			}
			// else {
			// 	logger.Infof("[JWTSecuredHandler] %s token NOT expired", s)
			// }

			if tokenValid == jwt.ErrTokenNotYetValid {
				logger.Infof("[JWTSecuredHandler] %s token is not yet valid", s)
				ctx = context.WithValue(ctx, ContextKey(JWTTokenNotYetValid), true)
			}
		}
	}

	// log.Printf("First Context: %+s\n", ctx)
	return ctx
}

//-----------------------------------------

// JWTContextValue func
func JWTContextValue(k string, r *http.Request) interface{} {
	ctx := r.Context()
	// logger.Info("[JWTContextValue] Cookie: %v", ctx.Value(ContextKey(k)))
	return ctx.Value(ContextKey(k))
}

//-----------------------------------------

// NewJWSToken function
func NewJWSToken(jwtid string, expiration time.Time, audience string, issuer string, notbefore time.Time, subject string, custom map[string]interface{}) ([]byte, error) {
	// bytes, _ := ioutil.ReadFile("./app.rsa")

	// logger.Info("[auth.NewJWSToken] Expiration: %v", expiration)
	claims := jws.Claims{}
	if jwtid != "" {
		claims.SetJWTID(jwtid)
	}
	claims.SetExpiration(expiration)
	claims.SetIssuedAt(time.Now().UTC())
	if audience != "" {
		claims.SetAudience(audience)
	}
	if issuer != "" {
		claims.SetIssuer(issuer)
	}
	if notbefore.After(time.Now().UTC()) {
		claims.SetNotBefore(notbefore)
	}
	if subject != "" {
		claims.SetSubject(subject)
	}
	if custom != nil {
		for k, v := range custom {
			claims.Set(k, v)
		}
	}

	// TODO: add a for loop to travers all available private keys
	rsaPrivate, err := crypto.ParseRSAPrivateKeyFromPEM([]byte(privateKeys[0]))
	if err != nil {
		logger.Errorf("Error: %v", err)
		return nil, err
	}

	jwt := jws.NewJWT(claims, crypto.SigningMethodRS256)
	b, _ := jwt.Serialize(rsaPrivate)

	return b, nil
}

// VerifyJWT function
func VerifyJWT(ctx context.Context) (bool, jwt.Claims) {

	// logger.Info("Context: %+v", ctx)

	found := ctx.Value(ContextKey(JWTTokenFound)).(bool)
	valid := ctx.Value(ContextKey(JWTTokenValid)).(bool)

	claims := ctx.Value(ContextKey(JWTClaim))
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

// TODO: implement to ability to periodically acquire and change the private tokens for JWT
