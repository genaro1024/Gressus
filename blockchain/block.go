package blockchain

type Block struct {
	Index        int64
	Timestamp    int64
	Transactions []Transaction
	Hash         string
	PrevHash     string
	Nonce        int64
	Log          string
}
