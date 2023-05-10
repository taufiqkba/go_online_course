package dto

import "github.com/golang-jwt/jwt/v5"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`       //using bearer_token
	ExpiredAt    string `json:"expired_at"` //expired for access_token
	Scope        string `json:"scope"`      // token available what that can do?
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

type MapClaimResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	IsAdmin       bool   `json:"is_admin"`
	jwt.MapClaims `json:"omitempty"`
}
