package models

import (
	"time"

	"github.com/google/uuid"
	//"go.starlark.net/lib/time"
)

type Payload struct {
    ID        uuid.UUID `json:d"`
    Username  string    `json:"usern`
    IssuedAt  time.Time `json:"Issue`
    ExpiredAt time.Time `json:"Expir"`
}


func NewPayload(username string, duration time.Duration) (*Payload, error) {
    tokenID, err := uuid.NewRandom()
    if err != nil {
        return nil, err
    }

    payload := &Payload{
        ID:        tokenID,
        Username:  username,
        IssuedAt:  time.Now(),
        ExpiredAt: time.Now().Add(duration),
    }
    return payload, nil
}