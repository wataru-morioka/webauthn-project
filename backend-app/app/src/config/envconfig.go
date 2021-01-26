package config

import "sync"

type Env struct {
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
			CognitoJwksEndpoint: "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_QiIPUCCfP/.well-known/jwks.json",
			Issure: "https://cognito-idp.ap-northeast-1.amazonaws.com/ap-northeast-1_QiIPUCCfP",
			PostgresEndpoint: "postgres-service:30432",
			PostrgresUser: "wataru",
			PostgresPassword: "wataru",
			PostgresDatabase: "webauthn",
		}
    })
    return env
}