package token

import (
	"os"
	"time"

	"ayaxos-inhouse/src/database"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"github.com/redis/go-redis/v9"
)

var _ = godotenv.Load()
var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Change this!

type InhouseClaims struct {
	InhouseID string   `json:"inhouse_id"` // Now using ULID
	Players   []string `json:"players"`
	Status    string   `json:"status"`
	jwt.RegisteredClaims
}

func GenerateInhouseToken(players []string, status string) (string, error) {
	// Generate a ULID
	inhouseID := ulid.Make().String()

	claims := InhouseClaims{
		InhouseID: inhouseID,
		Players:   players,
		Status:    status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * time.Minute)), // 1h30m
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Store in Redis with TTL of 90 minutes
	err = database.RedisClient.Set(database.Ctx, inhouseID, tokenString, 90*time.Minute).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyInhouseToken(tokenString string) (*InhouseClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &InhouseClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*InhouseClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// Check if token exists in Redis
	_, err = database.RedisClient.Get(database.Ctx, claims.InhouseID).Result()
	if err == redis.Nil {
		return nil, jwt.ErrSignatureInvalid // Token expired
	} else if err != nil {
		return nil, err
	}

	return claims, nil
}
