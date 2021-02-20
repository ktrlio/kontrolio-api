package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

// CustomClaims contains the structure of the custom
// JWT token claimer
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Help function to generate an IAM policy
func generatePolicy(principalID, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalID}

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

	// Optional output with custom properties of the String, Number or Boolean type.
	authResponse.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}
	return authResponse
}

// Authorizer middleware responsible for managing third party APIs access to any controller.
func Authorizer(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := event.AuthorizationToken
	switch strings.ToLower(token) {
	case "allow":
		return generatePolicy("user", "Allow", event.MethodArn), nil
	case "deny":
		return generatePolicy("user", "Deny", event.MethodArn), nil
	case "unauthorized":
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized") // Return a 401 Unauthorized response
	default:
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Error: Invalid token")
	}
}

// Issues JWT token for user authentication.
func signToken(email string) (*string, error) {
	jwtSecret := []byte(config.JWT_SECRET)

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := CustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func validateToken(tokenString string) (*string, error) {
	parsedToken, err := strconv.Unquote(tokenString)

	if err != nil && err != strconv.ErrSyntax {
		fmt.Println("Could not parse token: " + err.Error())
		return nil, errors.New("Something went wrong while parsing your token")
	}

	claims := &CustomClaims{}

	_, err = jwt.ParseWithClaims(parsedToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}

		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Unauthorized")
		}

		// This is far from ideal, but unfortunately the go-jwt package
		// doesn't expose this error type.
		// Might be interesting to submit a PR for this @ https://github.com/dgrijalva/jwt-go/
		if strings.Contains(err.Error(), "token is expired by") {
			return nil, errors.New("Your session expired. Please log in again to refresh it")
		}

		fmt.Println("Token parsing error: ", err.Error())
		return nil, errors.New("Sorry, something went wrong on our end")
	}

	return &claims.Email, nil
}

func isLoggedIn(jwtToken string) (*string, error) {
	data, err := parseSecret(jwtToken)

	if err != nil {
		fmt.Println("[isLoggedIn [0]] failed with: ", err.Error())
		return nil, err
	}

	email, err := validateToken(*data)

	if err != nil {
		fmt.Println("[isLoggedIn [1]] failed with: ", err.Error())
		return nil, err
	}

	return email, nil
}

// Login is used to sign in users
func Login(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data, err := parseUser(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	user := database.GetUserByEmail(data.Email)

	if user == nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("User not found or incorrect password.")})
	}

	match := utils.CheckPasswordHash(data.Password, user.Password)

	if !match {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("User not found or incorrect password.")})
	}

	token, err := signToken(user.Email)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String("Sorry, something went wrong on our end.")})
	}

	secret := Secret{
		SecretString: *token,
	}

	return apiResponse(http.StatusOK, secretResponseBody{secret})
}

// GetAPIKey is responsible for providing the API key for clients to authenticate requests
// without needing to sign in.
func GetAPIKey(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	email, err := isLoggedIn(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String(err.Error())})
	}

	user := database.GetUserByEmail(*email)

	if user == nil {
		return apiResponse(http.StatusBadRequest, errorBody{aws.String("User not found.")})
	}

	secret := Secret{
		SecretString: user.ApiKey,
	}

	return apiResponse(http.StatusOK, secretResponseBody{secret})
}
