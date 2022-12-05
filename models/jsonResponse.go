package models

type JsonResponse struct {
	Verify  bool     `json:"verify"`
	NoMatch []string `json:"noMatch"`
}
