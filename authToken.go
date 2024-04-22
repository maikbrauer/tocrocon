package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func GetTokens(c AuthorizationConfig) (t Tokens, err error) {

	if debugmode == "true" {
		fmt.Println("GrantType: ", c.GrantType)
		fmt.Println("Code: ", c.Code)
		fmt.Println("RedirectURI: ", c.RedirectURI)
	}

	var formVals = url.Values{}
	formVals.Set("client_id", ClientID)

	if debugmode == "true" {
		fmt.Println("FormVals: ", formVals.Encode())
		fmt.Println("grant", c.GrantType)
	}

	if c.GrantType == "authorization_code" {
		formVals.Set("client_secret", getSecretData("eu-central-1"))
		formVals.Set("code", c.Code)
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("redirect_uri", c.RedirectURI)
		formVals.Set("scope", scope)
	} else if c.GrantType == "urn:ietf:params:oauth:grant-type:device_code" {
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("device_code", c.Code)
	} else if c.GrantType == "refresh_token" {
		formVals.Set("client_secret", getSecretData("eu-central-1"))
		formVals.Set("grant_type", c.GrantType)
		formVals.Set("refresh_token", c.Code)
	}

	response, err := http.PostForm(TokenURL, formVals)

	if debugmode == "true" {
		fmt.Println("Formvals 2nd: ", formVals.Encode())
		fmt.Println("TokenURL: ", TokenURL)
	}

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

	if debugmode == "true" {
		str1 := fmt.Sprintf("%s", body)
		fmt.Println("Tocrocon-Version:", version)
		fmt.Println("Complete Body Response =", str1)
		fmt.Println("AccessToken: ", t.AccessToken)
		fmt.Println("RefreshToken: ", t.RefreshToken)
		fmt.Println("Expiry: ", t.Expiry)
	}

	fmt.Println("GrantType: ", c.GrantType, " processing successfull!")

	return
}
