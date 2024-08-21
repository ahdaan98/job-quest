package helper

import (
	"Auth/pkg/utils/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsEmployer struct {
	Id           uint   `json:"id"`
	Company_name string `json:"company_name"`
	Email        string `json:"email"`
	Role         string `json:"role"` // Add role field
	jwt.StandardClaims
}


func GenerateTokenEmployer(employer models.EmployerDetailsResponse) (string, error) {
	claims := &authCustomClaimsEmployer{
		Id:           employer.ID,
		Company_name: employer.CompanyName,
		Email:        employer.ContactEmail,
		Role:         "employer", // Set role to 'employer'
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}


func ValidateTokenEmployer(tokenString string) (*authCustomClaimsEmployer, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsEmployer{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*authCustomClaimsEmployer); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
