package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var apidata ApiData
	err := json.Unmarshal([]byte(request.Body), &apidata)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	var requestParams = AuthorizationConfig{
		Code:        apidata.Code,
		RedirectURI: apidata.RedirectURI,
		GrantType:   apidata.GrantType,
	}

	_token, err := GetTokens(requestParams)
	if err != nil {
		panic(err)
	}

	ResponseData := Tokens{
		AccessToken:  _token.AccessToken,
		RefreshToken: _token.RefreshToken,
		Expiry:       _token.Expiry,
	}

	aJson, err := json.Marshal(ResponseData)

	if err != nil {
		panic(err)
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(aJson),
		Headers:    map[string]string{"Tocrocon-Version": version},
	}
	return response, nil
}
