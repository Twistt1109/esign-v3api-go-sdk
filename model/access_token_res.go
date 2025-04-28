package model

type AccessTokenResData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

type AccessTokenRes struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    AccessTokenResData `json:"data"`
}
