package middleware

import (
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

func generatePolicy(principalID, effect, resource string, context map[string]interface{}) utils.AuthResponse {
	authResponse := utils.AuthResponse{PrincipalID: principalID}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	authResponse.Context = context
	return authResponse
}

func Authenticate(request utils.AuthRequest) (utils.AuthResponse, error) {
	token := request.AuthorizationToken
	tokenSlice := strings.Split(token, " ")
	var bearerToken string
	if len(tokenSlice) > 1 {
		bearerToken = tokenSlice[len(tokenSlice)-1]
	}
	if bearerToken == "" {
		return utils.AuthResponse{}, errors.New("Unauthorized")
	}

	return generatePolicy("user", "Allow", request.MethodArn, map[string]interface{}{"name": bearerToken}), nil
}
