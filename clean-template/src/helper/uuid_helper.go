package helper

import "github.com/google/uuid"

func Generate16DigitUUID() string {
	id, _ := uuid.NewUUID()
	return id.String()
}
