package config

type Env struct {
	cognitoClientId string
	cognitoClientSecret string
	postgresEndpoint string
	postrgresUser string
	postgresPassword string
	postgresDatabase string
}

func (e Env) New() *Env {
	env := &Env{
		cognitoClientId: "test",
		cognitoClientSecret: "test",
		postgresEndpoint: "test",
		postrgresUser: "test",
		postgresPassword: "test",
		postgresDatabase: "test",
	}
	return env
}