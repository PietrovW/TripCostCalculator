package dtos

import uuid "github.com/satori/go.uuid"

type CreateCalculateCommandResponse struct {
	CalculateID uuid.UUID `json:"calculateId"`
}
