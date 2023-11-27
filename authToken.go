package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetTokens(c AuthorizationConfig) (t Tokens, err error) {

	formVals := url.Values{}
	formVals.Set("code", c.AuthCode)
	formVals.Set("grant_type", "authorization_code")
	formVals.Set("redirect_uri", c.RedirectURI)
	formVals.Set("scope", scope)
	formVals.Set("client_secret", ClientSecret)
	formVals.Set("client_id", ClientID)

	response, err := http.PostForm(TokenURL, formVals)

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
