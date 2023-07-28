package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("bv93q4htugr=9wpg895q0sa!4fi4=")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
