package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User authenticated User
type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name"`
	DisplayName   string             `json:"displayName"`
	Photo         *Image             `json:"photo"`
	Email         string             `json:"email"`
	EmailVerified bool               `json:"emailVerified"`
	ProviderID    string             `json:"providerId"`
	UID           string             `json:"uid"`
	Password      *string            `json:"password"`
}

//CheckedUser in jwt token
type CheckedUser struct {
	ID    string   `json:"id"`
	Roles []string `json:"role"`
}

//UserToken ...
type UserToken struct {
	User   User   `json:"user"`
	Tokens Tokens `json:"tokens"`
}

//Tokens ...
type Tokens struct {
	Token                 string    `json:"token"`
	TokenExpiresAt        time.Time `json:"tokenExpiresAt"`
	RefreshToken          string    `json:"refreshToken"`
	RefreshTokenExpiresAt time.Time `json:"refreshTokenExpiresAt"`
}
