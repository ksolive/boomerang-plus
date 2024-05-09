package bots

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getAppAccessToken(appId string, clientSecret string) (string, error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{
		  "appId": "%s",
		  "clientSecret": "%s"
		}`, appId, clientSecret))
	req, err := http.NewRequest("POST", "https://bots.qq.com/app/getAppAccessToken", data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   string `json:"expires_in"`
	}
	var response AccessTokenResponse
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return "", err
	}
	if response.AccessToken == "" {
		return "Failed to get access token", nil
	}
	accessToken := response.AccessToken
	return accessToken, nil
}
