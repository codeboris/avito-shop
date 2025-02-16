package models

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type CoinTransaction struct {
	User   string `json:"user"`
	Amount int    `json:"amount"`
}

type CoinHistory struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type InfoResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}
