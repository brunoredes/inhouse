package token

import (
	"ayaxos-inhouse/src/database"
	"time"
)

// Store revoked tokens in Redis (for a short duration)
func RevokeInhouseToken(tokenString string) error {
	return database.RedisClient.Set(database.Ctx, "revoked:"+tokenString, "true", 2*time.Hour).Err()
}

// Check if a token is revoked
func IsTokenRevoked(tokenString string) bool {
	val, err := database.RedisClient.Get(database.Ctx, "revoked:"+tokenString).Result()
	return err == nil && val == "true"
}
