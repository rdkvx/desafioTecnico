package model

type ValidatePassword struct {
	Password string `json:"password"`
	Rules    []struct {
		Rule  string `json:"rule"`
		Value int    `json:"value"`
	} `json:"rules"`
}
