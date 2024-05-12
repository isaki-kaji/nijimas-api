package util

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// LoadFirebaseApp initializes and returns a Firebase App instance
func LoadFirebaseApp(config *Config) (*firebase.App, error) {
	opt := option.WithCredentialsFile(config.FirebaseCredentialPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
		return nil, err
	}
	return app, nil
}

// Provide Firebase services
func ProvideFirebaseAuth(app *firebase.App) (*auth.Client, error) {
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return auth, nil
}
