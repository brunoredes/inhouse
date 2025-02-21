package models

type Player struct {
	ID        int    `json:"id"`
	DiscordID string `json:"discord_id"`
	Username  string `json:"username"`
}
