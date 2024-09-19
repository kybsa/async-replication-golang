package dto

type Message struct {
	IdempotencyKey []byte // uuid
	Content        []byte
	CreatedAt      int64
}
