package token

import (
	"ayaxos-inhouse/config/database"
	"time"
)

// Store revoked tokens in Redis (for a short duration)
func RevokeInhouseToken(tokenString string) error {
	return database.RedisClient.Set(database.Ctx, "revoked:"+tokenString, "true", 10*time.Hour).Err()
}

func GetRevokedToken(tokenString string) (string, error) {
	return database.RedisClient.Get(database.Ctx, "revoked:"+tokenString).Result()

}

// Check if a token is revoked
func IsTokenRevoked(tokenString string) bool {
	val, err := GetRevokedToken(tokenString)
	return err == nil && val == "true"
}
