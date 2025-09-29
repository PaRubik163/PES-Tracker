package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type claims struct {
	ID string		`json:"uuid"`
	UserID int 	`json:"user_id"`
	jwt.RegisteredClaims
}

type Jwt struct {
	SecretKey []byte
}

type TokenResponse struct {
	Token string
	ID string
	UserID int
}

func NewJwt(secretKey []byte) *Jwt {
	return &Jwt{SecretKey: secretKey}
}

func (j *Jwt) GenerateToken(userID int) (*TokenResponse, error) {
	claims := claims{
		ID: uuid.New().String(),
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	strToken, err := token.SignedString(j.SecretKey)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token: strToken,
		ID: claims.ID,
	}, nil
}

func (j *Jwt) ValidateToken(token string) (*claims, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("wrong signing method")
		}

		return j.SecretKey, nil
	}

	claim := &claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claim, keyFunc)
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claim, nil
}