package main

import "os"

const baseURL = "https://login.microsoftonline.com/"
const tokenPath = "/oauth2/v2.0/token"
const version = "1.0.0"

var TokenURL = baseURL + os.Getenv("tenant_id") + tokenPath

var SecretManagerName = os.Getenv("secretmanager_name")
var ClientID = os.Getenv("client_id")
var scope = ClientID + "/.default offline_access"
var debugmode = os.Getenv("debugmode")

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       int    `json:"expires_in"`
}

type AuthorizationConfig struct {
	Code        string
	RedirectURI string
	GrantType   string
}

type ApiData struct {
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
	GrantType   string `json:"grant_type"`
}
