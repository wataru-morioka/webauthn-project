package auth

import (
	"log"
	"net/http"
	"context"
	"strings"
	"errors"
	"fmt"
	"time"
	"encoding/base64"
	"crypto/rsa"
	"encoding/binary"
	"math/big"

	"github.com/dgrijalva/jwt-go"

	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/config"
	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/data/repository"
	. "github.com/wataru-morioka/webauthn-project/backend-app/app/src/data/interface"
)

type CongnitoUserInfo struct {
	Sub string `json:"sub"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	PreferredUsername string `json:"preferred_username"`
	Email string `json:"email"`
}

type JWK struct {
    Keys []JWKKey
}

// JWKKey is json data struct for cognito jwk key
type JWKKey struct {
    Alg string
    E   string
    Kid string
    Kty string
    N   string
    Use string
}

type contextKey struct {
	name string
}

type Middleware struct{
	userCtxKey contextKey
}

func NewMiddleware() *Middleware {
	middleware := &Middleware {
		userCtxKey: contextKey{name: ContextKey,},
	}
	return middleware
}

func (m Middleware) VerifyAccessToken() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			bearerHeader := req.Header.Get("Authorization")
			accessToken := extractAccessToken(bearerHeader)

			uid, err := isAccessTokenValid(&accessToken)
			if err != nil {
				log.Printf("トークン検証エラー: %+v", err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			log.Printf("get sub!: %s", uid)

			var dbRepo DbRepositoryInterface = NewDbRepository()
			err = dbRepo.CreateAccount(&uid)
			if err != nil {
				log.Printf("データ登録エラー: %+v", err)
			}

			user, err := getUserInfo(bearerHeader)
			if err != nil {
				log.Print("ユーザ情報取得エラー")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(req.Context(), m.userCtxKey, user)
			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}
}

func getUserInfo(bearerHeader string) (*CongnitoUserInfo, error) {
	env := NewEnv()
	header := createHeader("Authorization", bearerHeader)
	userInfo := &CongnitoUserInfo{}
	var apiRepo ApiRepositoryInterface = NewApiRepository()

	err := apiRepo.ApiRequest(GET.String(), env.CognitoUserInfoEndpoint, userInfo, header)
	if err != nil {
		return nil, err
	}
	log.Printf("ユーザ情報: %s", userInfo.Email)

	return userInfo, nil
}

func createHeader(key string, value string) map[string]string {
	return map[string]string{key: value}
}

func isAccessTokenValid(token *string) (string, error) {
	log.Printf("access token: %s", *token)
	env := NewEnv()

	jwks, err := getJWK(env.CognitoJwksEndpoint)
	if err != nil {
		return "", fmt.Errorf("jwk取得エラー: %+v", err)
	}

	decodedToken, err := jwt.Parse(*token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return "", fmt.Errorf("Unexpected signing method: %+v", token.Header["alg"])
		}
		if kid, ok := token.Header["kid"]; ok {
			if kidStr, ok := kid.(string); ok {
				key := jwks[kidStr]
				rsaPublicKey := convertKey(key.E, key.N)
				return rsaPublicKey, nil
			}
		}
        return "", fmt.Errorf("invalid kid")
	})
	if err != nil {
		return "", fmt.Errorf("Hash処理エラー: %+v", err)
	}

	if !decodedToken.Valid {
		return "", fmt.Errorf("jwt署名エラー: %+v", decodedToken)
	}

	claims, ok := decodedToken.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("not found claims")
    }

	ok = claims.VerifyIssuer(env.Issure, false)
	if !ok {
		return "", errors.New("invalid issuer")
	} 

	tokenUseErr := func() error {
        if tokenUse, ok := claims["token_use"].(string); ok {
			if tokenUse == "access" {
				return nil
			}
		}
        return errors.New("token_use should be access")
    }
	if tokenUseErr() != nil {
		return "", tokenUseErr()
	}

	ok = claims.VerifyExpiresAt(time.Now().Unix(), false)
	if !ok {
		return "", errors.New("token is expired")
	} 

	// TODO scope, aud 検証

    uid, ok := claims["sub"].(string)
    if !ok {
        return "", errors.New("not found sub")
    }

	return uid, nil
}

func extractAccessToken(bearerHeader string) string {
	return strings.Replace(bearerHeader, "Bearer ", "", 1)
}

func getJWK(jwksURL string) (map[string]JWKKey, error) {
    jwk := &JWK{}
	var repo ApiRepositoryInterface = NewApiRepository()
    err := repo.ApiRequest(GET.String(), jwksURL, jwk, nil)
	if err != nil {
		return nil, err
	}

    jwkMap := make(map[string]JWKKey, 0)
    for _, jwk := range jwk.Keys {
        jwkMap[jwk.Kid] = jwk
    }
    return jwkMap, nil
}

func convertKey(rawE string, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	// fmt.Printf("%#v\n", *pubKey)
	return pubKey
}




