package data

type ApiRepositoryInterface interface {
	ApiRequest(url string, target interface{}) error
}