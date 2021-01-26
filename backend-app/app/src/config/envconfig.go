package config

import "sync"

type Env struct {
	CognitoUserInfoEndpoint string
	CognitoJwksEndpoint string
	Issure string
	PostgresEndpoint string
	PostrgresUser string
	PostgresPassword string
	PostgresDatabase string
}

func NewEnv() *Env {
	var once sync.Once
	var env *Env
	once.Do(func() {
        env = &Env{
			CognitoUserInfoEndpoint: "https://webauthn-app.auth.ap-northeast-1.amazoncognito.com/oauth2/userInfo",
			CognitoJwksEndpoint: "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_QiIPUCCfP/.well-known/jwks.json",
			Issure: "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_QiIPUCCfP",
			PostgresEndpoint: "postgres-service",
			PostrgresUser: "wataru",
			PostgresPassword: "wataru",
			PostgresDatabase: "webauthn",
		}
    })
    return env
}