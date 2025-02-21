package routes

import (
	"ayaxos-inhouse/src/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/inhouse", handlers.CreateInhouse)
	r.POST("/inhouse/:id/join", handlers.JoinInhouse)
	r.POST("/inhouse/:id/shuffle", handlers.ShuffleTeams)
	r.POST("/inhouse/:id/confirm", handlers.ConfirmTeam)
	r.GET("/inhouse/:id", handlers.GetInhouseDetails)
	r.POST("/inhouse/:id/finish", handlers.FinishLobby) // New route to finish lobby
}
