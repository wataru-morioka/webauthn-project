package data

type ApiRepositoryInterface interface {
	ApiRequest(method string, url string, target interface{}, header map[string]string) error
}

type DbRepositoryInterface interface {
	CreateAccount(uid *string) error
	// UpdateAccount() error
	// GetAccount() error
}