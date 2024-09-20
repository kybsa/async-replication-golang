package dto

type Message struct {
	IdempotencyKey []byte // uuid
	Content        []byte
	ExternalID     string
}
