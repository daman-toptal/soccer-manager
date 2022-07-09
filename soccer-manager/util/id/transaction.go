package id

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

/*
 * Transaction prefix + uuid will be used internally when passed between services and
 * removed when passing to/from the database.
 */
type TransactionID uuid.UUID

func (id TransactionID) Prefix() IDPrefix {
	return IDPrefixTransaction
}

func (id TransactionID) String() string {
	if id.IsZero() {
		return ""
	}
	return string(IDPrefixTransaction) + id.UUIDString()
}

func (id TransactionID) UUIDString() string {
	return uuid.UUID(id).String()
}

func (id TransactionID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

// Returns empty string when zero for omitempty
func (id TransactionID) JSONString() string {
	if id.IsZero() {
		return ""
	}

	return id.String()
}

func (id TransactionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *TransactionID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	uid, err := ParseTransactionID(s)
	if err != nil {
		return err
	}

	*id = uid
	return nil
}

// SQL value marshaller
func (id TransactionID) Value() (driver.Value, error) {
	return id.UUIDString(), nil
}

// SQL scanner
func (id *TransactionID) Scan(value interface{}) error {
	if value == nil {
		*id = TransactionID{}
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

	*id = TransactionID(uid)
	return nil
}

func NewTransactionID() (TransactionID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return TransactionID{}, err
	}

	return TransactionID(id), nil
}

func ParseTransactionID(id string) (TransactionID, error) {
	// Return nil id on empty string
	if id == "" {
		return TransactionID{}, nil
	}

	// Check prefix
	if !strings.HasPrefix(id, string(IDPrefixTransaction)) {
		return TransactionID{}, errors.New("invalid Transaction id prefix")
	}

	// Validate UUID
	uid, err := uuid.Parse(strings.TrimPrefix(id, string(IDPrefixTransaction)))
	if err != nil {
		return TransactionID{}, err
	}

	return TransactionID(uid), nil
}
