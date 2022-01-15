package entities

import "github.com/google/uuid"

type Transaction struct {
	ID        uuid.UUID
	Type      string
	Processor string
	Client    string
	Location  string
	Parent    *Transaction
}


func(t *Transaction) IsCardholderNotPresent() bool {
	return true
}
