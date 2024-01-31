package composite

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/jinzhu/gorm"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/model"
	"github.com/koba1108/go-backend-ddd-template/internals/domain/service"
	"github.com/koba1108/go-backend-ddd-template/internals/helper"
)

type sampleService struct {
	db         *gorm.DB
	authClient *auth.Client
}

func NewSampleService(db *gorm.DB, authClient *auth.Client) service.SampleService {
	return &sampleService{
		db:         db,
		authClient: authClient,
	}
}

func (s sampleService) GetSampleWithAuthUser(ctx context.Context, id int) (*model.Sample, *model.AuthUser, error) {
	var sample model.Sample
	if err := s.db.Find(&sample, id).Error; err != nil {
		return nil, nil, err
	}
	au, err := s.authClient.GetUser(ctx, sample.UID)
	if err != nil {
		return nil, nil, err
	}
	authUser := &model.AuthUser{
		DisplayName:    au.DisplayName,
		Email:          au.Email,
		PhoneNumber:    au.PhoneNumber,
		PhotoURL:       au.PhotoURL,
		ProviderID:     au.ProviderID,
		UID:            au.UID,
		CreatedAt:      helper.Int64ToTime(au.UserMetadata.CreationTimestamp),
		LastLoggedInAt: helper.Int64ToTimePrt(au.UserMetadata.LastLogInTimestamp),
		CustomClaims: &model.AuthUserCustomClaims{
			Admin: au.CustomClaims["admin"].(bool),
		},
	}
	return &sample, authUser, nil
}
