package client

import (
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func NewFirebase(opt option.ClientOption) (*firebase.App, error) {
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}
