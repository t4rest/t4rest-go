package api

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/t4rest/t4rest-go/httpserver/response"
)

const (
	authorizationHeader = "Authorization"
)

// Jwt http handler
type Jwt struct {
	publicKey *rsa.PublicKey
}

// NewJwt initialises Jwt structure by base64-encoded pem-key
func NewJwt(pubKey string) (*Jwt, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		return nil, fmt.Errorf("unable to parse RSA public key: %v", err)
	}

	return &Jwt{publicKey: publicKey}, nil
}

// JwtMiddleware validate jwt token
func (md *Jwt) JwtMiddleware(next httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		claims, token, err := md.getTokenClaims(r)
		if err != nil {
			response.ERROR(w, response.NewBadJwtError(err))
			return
		}

		_ = token
		_ = claims

		next(w, r, ps)
	}
}

func (md *Jwt) getTokenClaims(r *http.Request) (jwt.MapClaims, string, error) {
	authHeader := strings.Split(r.Header.Get(authorizationHeader), " ")
	if len(authHeader) < 2 {
		return nil, "", errors.New(fmt.Sprintf("bad %s header", authorizationHeader))
	}

	token, err := jwt.Parse(authHeader[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		return md.publicKey, nil
	})

	if err != nil {
		if er, ok := err.(*jwt.ValidationError); ok {
			return nil, "", er
		}
		return nil, "", errors.Wrap(err, "can not parse jwt token")
	}

	if token == nil {
		return nil, "", errors.New("token is empty")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, "", errors.New("token is invalid")
	}

	return claims, token.Raw, nil
}
