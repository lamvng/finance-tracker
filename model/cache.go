package model

type UserCache struct {
	UserID    string `redis:"UserID"`
	AuthToken string `redis:"AuthToken"`
}
