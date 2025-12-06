package notification

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

func InitFCM() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("internal/config/firebase_service_account.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting messaging client: %v\n", err)
	}

	fcmClient = client
}

func SendToToken(token string, data map[string]string) error {
	ctx := context.Background()

	msg := &messaging.Message{
		Token: token,
		Data:  data,
	}

	_, err := fcmClient.Send(ctx, msg)
	return err
}
