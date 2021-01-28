package repository

import (
	"net/http"
	"encoding/json"
	"time"
    "log"
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

    if res.StatusCode != 200 {
        log.Printf("apiリクエストエラー: %s %+v", res.StatusCode, res)
        return err
    }

    json.NewDecoder(res.Body).Decode(body)

    return nil
}

