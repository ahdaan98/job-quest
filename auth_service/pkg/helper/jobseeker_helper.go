package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"Auth/pkg/utils/models"
)

type authCustomClaimsJobSeeker struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokenJobSeeker(jobSeeker models.JobSeekerDetailsResponse) (string, error) {
	claims := &authCustomClaimsJobSeeker{
		Id:    jobSeeker.ID,
		Email: jobSeeker.Email,
		Role:  "jobseeker",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	return tokenString, nil
}

func ValidateTokenJobSeeker(tokenString string) (*authCustomClaimsJobSeeker, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsJobSeeker{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*authCustomClaimsJobSeeker); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims or token is not valid")
}
