package services

import (
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	GenerateAPIKey() string
	ValidateAPIKey(apiKey string) bool
}

type authService struct {
	apiKeys map[string]time.Time
}

func CreateAuthService() AuthService{
	return &authService{
		apiKeys: make(map[string]time.Time),
	}
}

func (a *authService) GenerateAPIKey() string {
	newKey := uuid.New().String()
	a.apiKeys[newKey] = time.Now()
	return newKey
}

func (a *authService) ValidateAPIKey(apiKey string) bool {		
	createTime, ok := a.apiKeys[apiKey]
	if !ok {
		return false
	}
	
	delete(a.apiKeys, apiKey)

	if time.Since(createTime) > time.Minute {		
		return false
	}

	return true
}