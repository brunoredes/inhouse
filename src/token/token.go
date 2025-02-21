package token

import (
	"time"

	"math/rand"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid"
)

var jwtSecret = []byte("89202d3e6b0be05e67ae39ea6d5d3de5") // Change to a secure secret

type InhouseClaims struct {
	InhouseID string   `json:"inhouse_id"` // ULID instead of int
	Players   []string `json:"players"`
	Status    string   `json:"status"`
	jwt.RegisteredClaims
}

func GenerateInhouseToken(players []string, status string) (string, error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	inhouseID := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	claims := InhouseClaims{
		InhouseID: inhouseID,
		Players:   players,
		Status:    status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token valid for 1 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyInhouseToken(tokenString string) (*InhouseClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &InhouseClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*InhouseClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
