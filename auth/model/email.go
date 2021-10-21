package model

type WelcomeEmailContext struct {
	Name     string
	Email    string
	Password string
	Host     string
}

type ResetEmailContext struct {
	Name string
	Url  string
}
