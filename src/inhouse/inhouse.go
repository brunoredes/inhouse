package inhouse

import (
	"errors"
	"math/rand"
	"sync"
)

type Lobby struct {
	InhouseID string
	Players   []string
	Teams     map[string][]string // "Team A" and "Team B"
	Confirmed map[string]bool     // Players who confirmed
}

var (
	lobbies = make(map[string]*Lobby)
	mutex   sync.Mutex
)

// Create a new inhouse lobby
func CreateLobby(id string, players []string) {
	mutex.Lock()
	defer mutex.Unlock()

	lobbies[id] = &Lobby{
		InhouseID: id,
		Players:   players,
		Teams:     make(map[string][]string),
		Confirmed: make(map[string]bool),
	}
}

// Add a player to an inhouse
func AddPlayer(id string, player string) error {
	mutex.Lock()
	defer mutex.Unlock()

	lobby, exists := lobbies[id]
	if !exists {
		return errors.New("inhouse not found")
	}

	for _, p := range lobby.Players {
		if p == player {
			return errors.New("player already in inhouse")
		}
	}

	lobby.Players = append(lobby.Players, player)
	return nil
}

// Shuffle players into random teams
func ShuffleTeams(id string) (map[string][]string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	lobby, exists := lobbies[id]
	if !exists {
		return nil, errors.New("inhouse not found")
	}

	rand.Shuffle(len(lobby.Players), func(i, j int) {
		lobby.Players[i], lobby.Players[j] = lobby.Players[j], lobby.Players[i]
	})

	mid := len(lobby.Players) / 2
	lobby.Teams["Team A"] = lobby.Players[:mid]
	lobby.Teams["Team B"] = lobby.Players[mid:]

	return lobby.Teams, nil
}

// Confirm team selection
func ConfirmTeam(id string, player string) error {
	mutex.Lock()
	defer mutex.Unlock()

	lobby, exists := lobbies[id]
	if !exists {
		return errors.New("inhouse not found")
	}

	lobby.Confirmed[player] = true
	return nil
}

// Get inhouse details
func GetLobbyDetails(id string) (*Lobby, error) {
	mutex.Lock()
	defer mutex.Unlock()

	lobby, exists := lobbies[id]
	if !exists {
		return nil, errors.New("inhouse not found")
	}

	return lobby, nil
}
