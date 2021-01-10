package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"
	"github.com/marcelovicentegc/kontrolio-api/config"
	"github.com/marcelovicentegc/kontrolio-api/database"
	"github.com/marcelovicentegc/kontrolio-api/utils"
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

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

func validateToken(tokenString string) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.JWT_SECRET, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"])
	} else {
		fmt.Println(err)
	}
}

func Login(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	data, err := parseUser(req.Body)

	if err != nil {
		return apiResponse(http.StatusBadGateway, errorBody{aws.String("Sorry, something went wrong while parsing the request")})
	}

	user := database.GetUser(data.Email)

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

	return apiResponse(http.StatusOK, responseBody{token})
}
