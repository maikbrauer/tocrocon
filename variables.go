package main

import "os"

const baseURL = "https://login.microsoftonline.com/"
const tokenPath = "/oauth2/v2.0/token"
const scope = "openid profile offline_access user.read"

var TokenURL = baseURL + os.Getenv("tenant_id") + tokenPath
var ClientSecret = os.Getenv("client_secret")
var ClientID = os.Getenv("client_id")

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       int    `json:"expires_in"`
}

type AuthorizationConfig struct {
	AuthCode    string
	RedirectURI string
}

type ApiData struct {
	AuthCode    string `json:"authcode"`
	RedirectURI string `json:"redirect_uri"`
}
