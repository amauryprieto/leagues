package tournaments

import (
	"time"

	"github.com/google/uuid"
)

type ID uuid.UUID

// Tournament struct
type Tournament struct {
	ID          ID
	Name        string
	Start       time.Time
	End         time.Time
	Description string
}

// Player struct
type Player struct {
	userID ID
}

// Team struct
type Team struct {
	players []Player
}
