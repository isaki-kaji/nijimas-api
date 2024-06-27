package util

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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
		log.Fatalf("error getting Auth client: %v", err)
		return nil, err
	}
	return auth, nil
}

func ProviderFirestore(app *firebase.App) (*firestore.Client, error) {
	firestore, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
		return nil, err
	}
	return firestore, nil
}
