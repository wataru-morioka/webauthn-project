package repository

import (
	"net/http"
	"encoding/json"
	"time"
)

type ApiRepository struct {}

func NewApiRepository() *ApiRepository {
	return &ApiRepository{}
}

func (a ApiRepository) ApiRequest(
    method string, 
    url string, 
    body interface{}, 
    header map[string]string,
) error {
    var apiClient = &http.Client{Timeout: 10 * time.Second}

    req, _ := http.NewRequest(method, url, nil)
    for key, value := range header {
        req.Header.Set(key, value)
    }

    res, err := apiClient.Do(req)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    // TODO response header httpステータス検証

    json.NewDecoder(res.Body).Decode(body)

	return nil
}

