package models

type User struct {
	ID                   uint
	Username             string
	Coins                int
	Purchases            []Purchase
	SentTransactions     []Transaction
	ReceivedTransactions []Transaction
}
