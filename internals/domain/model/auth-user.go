package model

import "time"

type AuthUser struct {
	DisplayName    string                `json:"displayName,omitempty"`
	Email          string                `json:"email,omitempty"`
	PhoneNumber    string                `json:"phoneNumber,omitempty"`
	PhotoURL       string                `json:"photoUrl,omitempty"`
	ProviderID     string                `json:"providerId,omitempty"`
	UID            string                `json:"rawId,omitempty"`
	CreatedAt      time.Time             `json:"createdAt,omitempty"`
	LastLoggedInAt *time.Time            `json:"lastLoggedInAt,omitempty"`
	CustomClaims   *AuthUserCustomClaims `json:"customClaims,omitempty"`
}

type AuthUserCustomClaims struct {
	Admin bool `json:"admin"`
}
