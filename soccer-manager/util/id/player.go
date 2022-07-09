package id

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type PlayerID uuid.UUID

func (id PlayerID) Prefix() IDPrefix {
	return IDPrefixPlayer
}

func (id PlayerID) String() string {
	if id.IsZero() {
		return ""
	}
	return string(IDPrefixPlayer) + id.UUIDString()
}

func (id PlayerID) UUIDString() string {
	return uuid.UUID(id).String()
}

func (id PlayerID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

// Returns empty string when zero for omitempty
func (id PlayerID) JSONString() string {
	if id.IsZero() {
		return ""
	}

	return id.String()
}

func (id PlayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *PlayerID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	uid, err := ParsePlayerID(s)
	if err != nil {
		return err
	}

	*id = uid
	return nil
}

// SQL value marshaller
func (id PlayerID) Value() (driver.Value, error) {
	return id.UUIDString(), nil
}

// SQL scanner
func (id *PlayerID) Scan(value interface{}) error {
	if value == nil {
		*id = PlayerID{}
		return nil
	}

	val, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Unable to scan value %v", value)
	}

	uid, err := uuid.Parse(string(val))
	if err != nil {
		return err
	}

	*id = PlayerID(uid)
	return nil
}

func NewPlayerID() (PlayerID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return PlayerID{}, err
	}

	return PlayerID(id), nil
}

func ParsePlayerID(id string) (PlayerID, error) {
	// Return nil id on empty string
	if id == "" {
		return PlayerID{}, nil
	}

	// Check prefix
	if !strings.HasPrefix(id, string(IDPrefixPlayer)) {
		return PlayerID{}, errors.New("invalid player id prefix")
	}

	// Validate UUID
	uid, err := uuid.Parse(strings.TrimPrefix(id, string(IDPrefixPlayer)))
	if err != nil {
		return PlayerID{}, err
	}

	return PlayerID(uid), nil
}
