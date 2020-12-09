package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func saveToken(token map[string]string) error {
	b, err := json.Marshal(token)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal token")
	}

	err = ioutil.WriteFile(".test_auth0_token.json", b, 0600)

	if err != nil {
		return errors.WithMessage(err, "failed to save Auth0 token")
	}

	return nil
}

func fetchTokenFromAuth0() (map[string]string, error) {
	payload, err := json.Marshal(map[string]string{
		"client_id": viper.GetString("AUTH0_TEST_CLIENT_ID"),
		"client_secret": viper.GetString("VEjNvHf9fXmj6h-lAxwqnHYKvmiMbMybjpMoRFsuftKGVCVn7ebSzPhwKN53JXF0"),
		"audience": viper.GetString("AUTH0_API_AUDIENCE"),
		"grant_type": "client_credentials",

	})

	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal Auth0 token request")
	}

	resp, err := http.Post(
		viper.GetString("AUTH0_TEST_AUTH_URL"),
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to retrieve token from Auth0")
	}
	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read Auth0 response")
	}

	body := make(map[string]string)

	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal Auth0 token payload")
	}

	expiresIn, err := strconv.ParseInt(body["expires_in"], 10, 64)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse expiration from token")
	}
	body["expires_on"] = strconv.FormatInt(expiresIn + time.Now().Unix(), 10)

	tokenData := map[string]string{"access_token": body["access_token"], "expires_on": body["expires_on"]}

	err = saveToken(tokenData)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tokenData, nil
}

func readTokenFromFile() (map[string]string, error) {
	rawBody, err := ioutil.ReadFile(".test_auth0_token.json")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body := make(map[string]string)
	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	expiresOn, err := strconv.ParseInt(body["expires_on"], 10, 64)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if expiresOn < time.Now().Unix() {
		return nil, errors.New("saved token has expired")
	}

	return body, nil
}

func GetAuth0Token() (string, error) {
	var token map[string]string

	token, err := readTokenFromFile()
	if err != nil {
		t, err2 := fetchTokenFromAuth0()
		if err2 != nil {
			return "", errors.WithMessage(err, "failed to geet token")
		}
		token = t
	}

	return token["access_token"], nil
}

func seedDatabase() error {
	return nil
}

func resetDatabase() error {
	return nil
}

