package models

type Produto struct {
	Id    string  `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}
