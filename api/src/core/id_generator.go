package core

import (
	"github.com/google/uuid"
)

func GetRandomId() string {
	return uuid.New().String()
}
