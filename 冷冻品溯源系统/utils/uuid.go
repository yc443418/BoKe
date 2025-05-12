package utils

import (
	"github.com/google/uuid"
)

// GenerateUUID 生成一个UUID
func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}
