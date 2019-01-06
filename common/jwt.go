package common

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// SignToken sign jwt token using JWT_SECRET
func SignToken(claims jwt.StandardClaims) (token string, err error) {
	expireDuration, err := time.ParseDuration(os.Getenv("JWT_EXPIRE_DURATION"))

	if err != nil {
		return "", err
	}

	claims.ExpiresAt = time.Now().Add(expireDuration).Unix()
	claims.IssuedAt = time.Now().Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	return t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// VerifyToken verify token
func VerifyToken(token string, issuer string) (claims *jwt.StandardClaims, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*jwt.StandardClaims); ok && t.Valid && claims.VerifyIssuer(issuer, true) {
		return claims, nil
	}

	return nil, jwt.NewValidationError("", jwt.ValidationErrorUnverifiable)
}
