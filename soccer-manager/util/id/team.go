package id

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type TeamID uuid.UUID

func (id TeamID) Prefix() IDPrefix {
	return IDPrefixTeam
}

func (id TeamID) String() string {
	if id.IsZero() {
		return ""
	}
	return string(IDPrefixTeam) + id.UUIDString()
}

func (id TeamID) UUIDString() string {
	return uuid.UUID(id).String()
}

func (id TeamID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

// Returns empty string when zero for omitempty
func (id TeamID) JSONString() string {
	if id.IsZero() {
		return ""
	}

	return id.String()
}

func (id TeamID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *TeamID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	uid, err := ParseTeamID(s)
	if err != nil {
		return err
	}

	*id = uid
	return nil
}

// SQL value marshaller
func (id TeamID) Value() (driver.Value, error) {
	return id.UUIDString(), nil
}

// SQL scanner
func (id *TeamID) Scan(value interface{}) error {
	if value == nil {
		*id = TeamID{}
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

	*id = TeamID(uid)
	return nil
}

func NewTeamID() (TeamID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return TeamID{}, err
	}

	return TeamID(id), nil
}

func ParseTeamID(id string) (TeamID, error) {
	// Return nil id on empty string
	if id == "" {
		return TeamID{}, nil
	}

	// Check prefix
	if !strings.HasPrefix(id, string(IDPrefixTeam)) {
		return TeamID{}, errors.New("invalid team id prefix")
	}

	// Validate UUID
	uid, err := uuid.Parse(strings.TrimPrefix(id, string(IDPrefixTeam)))
	if err != nil {
		return TeamID{}, err
	}

	return TeamID(uid), nil
}
