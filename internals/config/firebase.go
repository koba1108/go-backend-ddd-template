package config

import (
	"google.golang.org/api/option"
	"os"
)

func NewFirebaseConfig() option.ClientOption {
	return option.WithAPIKey(os.Getenv("FIREBASE_API_KEY"))
}
