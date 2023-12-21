package config

import "os"

type Email struct {
	EmailAuthPass string `json:"emailAuthPass"`
	EmailAddress  string `json:"emailAddress"`
	EmailFrom     string `json:"emailFrom"`
	EmailHost     string `json:"emailHost"`
}

func newEmail() *Email {
	return &Email{
		EmailAuthPass: os.Getenv("EMAIL_AUTH_PASS"),
		EmailAddress:  os.Getenv("EMAIL_ADDRESS"),
		EmailFrom:     os.Getenv("EMAIL_FROM"),
		EmailHost:     os.Getenv("EMAIL_HOST"),
	}
}
