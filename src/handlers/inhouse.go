package handlers

import (
	"ayaxos-inhouse/src/database"
	"ayaxos-inhouse/src/inhouse"
	"ayaxos-inhouse/src/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create an inhouse and return a JWT
func CreateInhouse(c *gin.Context) {
	var req struct {
		Players []string `json:"players"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tokenString, err := token.GenerateInhouseToken(req.Players, "waiting")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	inhouse.CreateLobby(tokenString, req.Players)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Player joins an inhouse
func JoinInhouse(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Player string `json:"player"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := inhouse.AddPlayer(id, req.Player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player added successfully"})
}

// Shuffle players into teams
func ShuffleTeams(c *gin.Context) {
	id := c.Param("id")
	teams, err := inhouse.ShuffleTeams(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teams": teams})
}

// Player confirms their team
func ConfirmTeam(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Player string `json:"player"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := inhouse.ConfirmTeam(id, req.Player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player confirmed team"})
}

// Get inhouse details
func GetInhouseDetails(c *gin.Context) {
	id := c.Param("id")
	lobby, err := inhouse.GetLobbyDetails(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lobby": lobby})
}

func FinishLobby(c *gin.Context) {
	inhouseID := c.Param("id")
	ctx := database.Ctx

	token.RevokeInhouseToken(inhouseID)

	revokedToken, tokenError := token.GetRevokedToken(inhouseID)
	if tokenError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get revoked token"})
	}

	// ✅ Step 1: Insert revoked token into PostgreSQL
	_, err := database.DB.Exec(ctx, "INSERT INTO revoked_tokens (token) VALUES ($1::TEXT)", revokedToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store revoked token"})
		return
	}

	// ✅ Step 2: Fetch all players from inhouse
	rows, err := database.DB.Query(ctx, "SELECT player_id FROM inhouse_players WHERE inhouse_id = $1", inhouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve players"})
		return
	}
	defer rows.Close()

	// ✅ Step 3: Insert inhouse data into game_results
	for rows.Next() {
		var playerID int
		if err := rows.Scan(&playerID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process players"})
			return
		}

		_, err := database.DB.Exec(ctx, "INSERT INTO game_results (inhouse_id, winning_team) VALUES ($1, 'TBD')", inhouseID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert game results"})
			return
		}
	}

	// ✅ Step 5: Update inhouse status to 'completed'
	_, err = database.DB.Exec(ctx, "UPDATE inhouses SET status = 'completed' WHERE id = $1", inhouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inhouse status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lobby finished successfully"})
}
