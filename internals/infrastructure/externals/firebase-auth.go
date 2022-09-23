package externals

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/repository"
	"github.com/koba1108/go-backend-ddd-template/internals/helper"
)

type firebaseAuthRepository struct {
	client *auth.Client
}

func NewFirebaseAuthRepository(fbApp *firebase.App) repository.AuthRepository {
	client, _ := fbApp.Auth(context.Background())
	return &firebaseAuthRepository{
		client: client,
	}
}

func (far *firebaseAuthRepository) GetUserByID(ctx context.Context, id string) (*model.AuthUser, error) {
	u, err := far.client.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.AuthUser{
		DisplayName:    u.DisplayName,
		Email:          u.Email,
		PhoneNumber:    u.PhoneNumber,
		PhotoURL:       u.PhotoURL,
		ProviderID:     u.ProviderID,
		UID:            u.UID,
		CreatedAt:      helper.Int64ToTime(u.UserMetadata.CreationTimestamp),
		LastLoggedInAt: helper.Int64ToTimePrt(u.UserMetadata.LastLogInTimestamp),
		CustomClaims: &model.AuthUserCustomClaims{
			Admin: u.CustomClaims["admin"].(bool),
		},
	}, nil
}
