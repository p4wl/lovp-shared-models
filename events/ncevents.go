package events

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

/*
	kafka events for NetCommander
*/

type EventType string

const (
	CREATE_USER              EventType = "user.create"
	CREATE_NETWORK           EventType = "network.create"
	ADD_USER_TO_NETWORK      EventType = "network.addusr"
	REMOVE_USER_FROM_NETWORK EventType = "network.rmusr"
	MODIFY_USER_ACCESS       EventType = "network.modusr"
)

type CmdEnvelope struct {
	EventId   uuid.UUID       `json:"eventId"`
	EventType EventType       `json:"eventType"`
	CreatedBy string          `json:"createdBy"`
	CreatedAt time.Time       `json:"createdAt"`
	Data      json.RawMessage `json:"data"`
}

type UserCreateCmd struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type NetworkCreateCmd struct {
	NetworkName string `json:"networkName"`
	Hosts       int    `json:"hosts"` // amount of hosts for the network
	Owner       string `json:"owner"` // username of network owner
}

type NetworkAddUserCmd struct {
	NetworkName string `json:"networkName"`
	Username    string `json:"username"`
}

type NetworkRmUserCmd struct {
	NetworkName string `json:"networkName"`
	Username    string `json:"username"`
}
