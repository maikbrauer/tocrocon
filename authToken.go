package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetTokens(c AuthorizationConfig) (t Tokens, err error) {

	fmt.Println(c.GrantType)
	fmt.Println(c.Code)
	fmt.Println(c.RedirectURI)

	var formVals = url.Values{}
	formVals.Set("client_secret", getSecretData("eu-central-1"))
	formVals.Set("client_id", ClientID)
	fmt.Println(formVals.Encode())
	fmt.Println("grant", c.GrantType)

	if c.GrantType == "authorization_code" {
		formVals.Set("code", c.Code)
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("redirect_uri", c.RedirectURI)
		formVals.Set("scope", scope)
	} else if c.GrantType == "urn:ietf:params:oauth:grant-type:device_code" {
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("device_code", c.Code)
	} else if c.GrantType == "refresh_token" {
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("refresh_token", c.Code)
	}

	response, err := http.PostForm(TokenURL, formVals)

	fmt.Println(formVals.Encode())
	fmt.Println(TokenURL)

	if err != nil {
		return t, errors.Wrap(err, "error while trying to get tokens")
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return t, errors.Wrap(err, "error while trying to read token json body")
	}

	err = json.Unmarshal(body, &t)
	if err != nil {
		return t, errors.Wrap(err, "error while trying to parse token json body")
	}

	return
}
