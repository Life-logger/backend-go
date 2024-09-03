package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CertificationsGetter defines the interface for obtaining OAuth tokens.
type TokenGetter interface {
	GetAccessToken(authorizeCode string) (string, string)
	GetNewAccessToken(refreshToken string) string
	GetUserInfo(accessToken string) (string, string, error)
}

// kakaoImpl is a struct that implements the CertificationsGetter interface for Kakao.
type kakaoImpl struct{}

// NewCertificationsGetter is a constructor function that returns an implementation of CertificationsGetter.
// It decides the implementation based on environment variables.
var NewTokenGetter = func() TokenGetter {
	// Currently only the Kakao implementation is available
	return new(kakaoImpl)
}

// GetAccessToken exchanges the authorization code for an access token from the Kakao API.
func (m kakaoImpl) GetAccessToken(authorizeCode string) (string, string) {
	kakaotalk_url := "https://kauth.kakao.com/oauth/token"

	// Construct the data for the POST request
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", "047f9409d2369cfb8f5e8afa41b2d9eb")
	data.Set("redirect_uri", "http://localhost:22250/login/callback")
	data.Set("code", authorizeCode)

	// Create a POST request
	req, err := http.NewRequest(http.MethodPost, kakaotalk_url, strings.NewReader(data.Encode()))
	if err != nil {
		panic(fmt.Sprintf("client: could not create request: %v", err))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("client: error making http request: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("client: received non-OK HTTP status: %s", res.Status))
	}

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Sprintf("client: error reading http response body: %v", err))
	}

	// Parse the JSON response to extract the access token
	var result map[string]interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		panic(fmt.Sprintf("client: error unmarshalling response body: %v", err))
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		panic("client: access_token not found in response")
	}

	refreshToken, ok := result["refresh_token"].(string)
	if !ok {
		panic("client: refresh_token not found in response")
	}

	return accessToken, refreshToken
}

func (m kakaoImpl) GetNewAccessToken(refreshToken string) string {
	kakaotalk_url := "https://kauth.kakao.com/oauth/token"

	// Construct the data for the POST request
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", "047f9409d2369cfb8f5e8afa41b2d9eb")
	data.Set("refresh_token", refreshToken)

	// Create a POST request
	req, err := http.NewRequest(http.MethodPost, kakaotalk_url, strings.NewReader(data.Encode()))
	if err != nil {
		panic(fmt.Sprintf("client: could not create request: %v", err))
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("client: error making http request: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("client: received non-OK HTTP status: %s", res.Status))
	}

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Sprintf("client: error reading http response body: %v", err))
	}

	// Parse the JSON response to extract the access token
	var result map[string]interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		panic(fmt.Sprintf("client: error unmarshalling response body: %v", err))
	}

	newAccessToken, ok := result["access_token"].(string)
	if !ok {
		panic("client: access_token not found in response")
	}

	return newAccessToken
}

func (m kakaoImpl) GetUserInfo(accessToken string) (string, string, error) {
	kakaotalkUserInfoURL := "https://kapi.kakao.com/v2/user/me"

	req, err := http.NewRequest(http.MethodGet, kakaotalkUserInfoURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("client: could not create request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("client: error making http request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("client: received non-OK HTTP status: %s", res.Status)
	}

	// Read the response body
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("client: error reading http response body: %v", err)
	}

	// Parse the JSON response to extract the nickname
	var result map[string]interface{}
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return "", "", fmt.Errorf("client: error unmarshalling response body: %v", err)
	}

	// Access the nickname from the response
	nickname, ok := result["kakao_account"].(map[string]interface{})["profile"].(map[string]interface{})["nickname"].(string)
	if !ok {
		return "", "", fmt.Errorf("client: nickname not found in response")
	}

	email, ok := result["kakao_account"].(map[string]interface{})["email"].(string)
	if !ok {
		return "", "", fmt.Errorf("client: email not found in response")
	}

	return nickname, email, nil
}
