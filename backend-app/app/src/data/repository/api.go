package repository

import (
	"net/http"
	"encoding/json"
	"time"
)

type ApiRepository struct {}

func (a ApiRepository) ApiRequest(url string, target interface{}) error {
    var apiClient = &http.Client{Timeout: 10 * time.Second}
    res, err := apiClient.Get(url)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    json.NewDecoder(res.Body).Decode(target)

	return nil
}