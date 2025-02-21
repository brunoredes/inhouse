package routes

import (
	"ayaxos-inhouse/src/inhouse"
	"ayaxos-inhouse/src/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/inhouse", CreateInhouse)
	r.POST("/inhouse/:id/join", JoinInhouse)
	r.POST("/inhouse/:id/shuffle", ShuffleTeams)
	r.POST("/inhouse/:id/confirm", ConfirmTeam)
	r.GET("/inhouse/:id", GetInhouseDetails)
}

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
