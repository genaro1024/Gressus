package blockchain

type Transaction struct {
	Id                 int64
	Type               TypeTransaction
	SourceAddress      string
	DestinationAddress string
	Amount             float32
	Coin               string
}
